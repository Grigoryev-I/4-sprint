package daysteps

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	if data == "" {
		return 0, 0, errors.New("данные отсутствуют")
	}

	if data != strings.TrimSpace(data) {
		return 0, 0, fmt.Errorf("данные содержат начальные или конечные пробелы")
	}

	parts := strings.Split(data, ",")
	if len(parts) != 2 {
		return 0, 0, errors.New("невалидные данные")
	}

	steps, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil || steps <= 0 {
		return 0, 0, errors.New("количество шагов должно быть больше 0")
	}

	duration, err := time.ParseDuration(strings.TrimSpace(parts[1]))
	if err != nil || duration <= 0 {
		return 0, 0, errors.New("продолжительность должна быть больше 0")
	}

	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	steps, duration, err := parsePackage(data)

	if err != nil || steps <= 0 {
		log.Println(err.Error())
		return ""
	}

	distanceM := float64(steps) * stepLength
	distanceKm := distanceM / mInKm
	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, distanceKm, calories)
}
