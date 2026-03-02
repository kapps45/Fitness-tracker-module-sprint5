package trainings

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

// Training хранит данные о тренировке, включая данные персоны.
type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

// Parse разбирает строку формата "3456,Ходьба,3h00m".
//
// Принимает:
//   - string — строку с данными о тренировке.
//
// Возвращает:
//   - error — ошибку, если данные некорректны.
func (t *Training) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 3 {
		return errors.New("неверный формат набора данных для Training")
	}
	stepsStr := strings.TrimSpace(parts[0])
	steps, err := strconv.Atoi(stepsStr)
	if err != nil {
		return err
	}
	t.Steps = steps

	t.TrainingType = strings.TrimSpace(parts[1])

	durStr := strings.TrimSpace(parts[2])
	dur, perr := time.ParseDuration(durStr)
	if perr != nil {
		dur, perr = time.ParseDuration("0h" + durStr)
		if perr != nil {
			return perr
		}
	}
	t.Duration = dur

	if t.Steps <= 0 {
		return errors.New("количество шагов должно быть положительным")
	}
	if dur <= 0 {
		return errors.New("продолжительность должна быть положительной")
	}

	return nil
}

// ActionInfo формирует строку с информацией о тренировке.
//
// Возвращает:
//   - string — строку с результатами;
//   - error — ошибку, если тип тренировки неизвестен
//     или произошла ошибка при расчётах.
func (t Training) ActionInfo() (string, error) {
	distanceKm := spentenergy.Distance(t.Steps, t.Height)
	meanSp := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	var calories float64
	var err error
	switch t.TrainingType {
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "", errors.New("неизвестный тип тренировки")
	}
	if err != nil {
		return "", err
	}

	// Форматированная строка с завершающим переносом строки
	info := "Тип тренировки: " + t.TrainingType + "\n" +
		"Длительность: " + formatDurationHours(t.Duration) + "\n" +
		"Дистанция: " + formatDistance(distanceKm) + "\n" +
		"Скорость: " + formatSpeed(meanSp) + "\n" +
		"Сожгли калорий: " + formatCalories(calories) + "\n"

	return info, nil
}

// Вспомогательные функции форматирования
func formatDurationHours(d time.Duration) string {
	hours := d.Hours()
	return strings.TrimSpace(formatFloat(hours) + " ч.")
}

func formatDistance(km float64) string {
	return formatFloat(km) + " км."
}

func formatSpeed(kmh float64) string {
	return formatFloat(kmh) + " км/ч"
}

func formatCalories(c float64) string {
	return formatFloat(c)
}

func formatFloat(v float64) string {
	return strconv.FormatFloat(v, 'f', 2, 64)
}
