package spentenergy

import (
	"fmt"
	"time"
)

const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

// Distance вычисляет дистанцию в километрах.
//
// Принимает:
//   - int — количество шагов;
//   - float64 — рост пользователя.
//
// Возвращает:
//   - float64 — дистанцию в километрах.
func Distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient
	totalMeters := float64(steps) * stepLength
	return totalMeters / mInKm
}

// MeanSpeed вычисляет среднюю скорость в км/ч.
//
// Принимает:
//   - int — количество шагов;
//   - float64 — рост пользователя;
//   - time.Duration — продолжительность активности.
//
// Возвращает:
//   - float64 — среднюю скорость в км/ч.
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps < 0 || duration <= 0 {
		return 0
	}
	d := Distance(steps, height)
	return d / duration.Hours()
}

// WalkingSpentCalories рассчитывает количество калорий,
// потраченных при ходьбе.
//
// Принимает:
//   - int — количество шагов;
//   - float64 — вес пользователя;
//   - float64 — рост пользователя;
//   - time.Duration — продолжительность ходьбы.
//
// Возвращает:
//   - float64 — количество сожжённых калорий;
//   - error — ошибку, если входные параметры некорректны.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, fmt.Errorf("количество шагов должно быть положительным")
	}
	if weight <= 0 {
		return 0, fmt.Errorf("вес должен быть положительным")
	}
	if height <= 0 {
		return 0, fmt.Errorf("рост должен быть положительным")
	}
	if duration <= 0 {
		return 0, fmt.Errorf("продолжительность должна быть положительной")
	}

	meanSp := MeanSpeed(steps, height, duration)
	cal := (weight * meanSp * duration.Minutes()) / float64(minInH)
	return cal * walkingCaloriesCoefficient, nil
}

// RunningSpentCalories рассчитывает количество калорий,
// потраченных при беге.
//
// Принимает:
//   - int — количество шагов;
//   - float64 — вес пользователя;
//   - float64 — рост пользователя;
//   - time.Duration — продолжительность бега.
//
// Возвращает:
//   - float64 — количество сожжённых калорий;
//   - error — ошибку, если входные параметры некорректны.
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, fmt.Errorf("количество шагов должно быть положительным")
	}
	if weight <= 0 {
		return 0, fmt.Errorf("вес должен быть положительным")
	}
	if height <= 0 {
		return 0, fmt.Errorf("рост должен быть положительным")
	}
	if duration <= 0 {
		return 0, fmt.Errorf("продолжительность должна быть положительной")
	}

	meanSp := MeanSpeed(steps, height, duration)
	cal := (weight * meanSp * duration.Minutes()) / float64(minInH)
	return cal, nil
}
