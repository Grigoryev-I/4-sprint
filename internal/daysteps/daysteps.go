package daysteps

import (
	"fmt"
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
	str := strings.Split(data, ",")

	if len(str) != 2 {
		return 0, 0, fmt.Errorf("невалидные данные")
	}

	steps, err := strconv.Atoi(str[0])
	if err != nil {
		return 0, 0, err
	}

	if steps == 0 {
		return 0, 0, fmt.Errorf("количество шагов должно быть больше 0")
	}

	duration, err := time.ParseDuration(str[1])
	if err != nil {
		return 0, 0, err
	}

	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	steps, duration, err := parsePackage(data)

	if err != nil {
		fmt.Println(err)
		return ""
	}

	if steps == 0 {
		return ""
	}

	distanceM := float64(steps) * stepLength
	distanceKm := distanceM / mInKm
	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила: %.2f км.\nВы сожгли %.2f ккал.", steps, distanceKm, calories)
}
