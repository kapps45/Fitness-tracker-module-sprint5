package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

// DaySteps содержит данные о дневной активности.
type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

// Parse разбирает строку формата "678,0h50m".
//
// Принимает:
//   - string — строку с данными о прогулке.
//
// Возвращает:
//   - error — ошибку, если данные некорректны.
func (ds *DaySteps) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return errors.New("неверный формат набора данных для DaySteps")
	}
	stepsStr := parts[0]
	if stepsStr == "" {
		return errors.New("неверный формат количества шагов")
	}
	steps, err := strconv.Atoi(stepsStr)
	if err != nil {
		return fmt.Errorf("неверный формат количества шагов: %w", err)
	}
	if steps <= 0 {
		return errors.New("количество шагов должно быть положительным")
	}
	ds.Steps = steps

	durStr := strings.TrimSpace(parts[1])
	dur, perr := time.ParseDuration(durStr)
	if perr != nil {
		dur, perr = time.ParseDuration("0h" + durStr)
		if perr != nil {
			return perr
		}
	}
	if dur <= 0 {
		return errors.New("продолжительность должна быть положительной")
	}
	ds.Duration = dur

	return nil
}

// ActionInfo формирует строку с информацией о прогулке.
//
// Возвращает:
//   - string — строку с результатами;
//   - error — ошибку, если произошла ошибка расчётов.
func (ds DaySteps) ActionInfo() (string, error) {
	distanceKm := spentenergy.Distance(ds.Steps, ds.Height)
	cal, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}
	info := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, distanceKm, cal)
	return info, nil
}
