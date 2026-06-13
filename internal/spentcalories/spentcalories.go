package spentcalories

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	str := strings.Split(data, ",")

	if len(str) != 3 {
		return 0, "", 0, fmt.Errorf("невалидные данные")
	}

	steps, err := strconv.Atoi(str[0])
	if err != nil {
		return 0, "", 0, err
	}

	duration, err := time.ParseDuration(str[2])
	if err != nil {
		return 0, "", 0, err
	}

	activityType := str[1]

	return steps, activityType, duration, nil
}

func distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient

	return float64(steps) * stepLength / mInKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}

	distance := distance(steps, height)

	return distance / duration.Hours()
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	steps, activityType, duration, err := parseTraining(data)
	if err != nil {
		log.Println(err)
		return "", err
	}

	durationHours := duration.Hours()
	distanceKm := distance(steps, height)
	meanSpeed := meanSpeed(steps, height, duration)
	var calories float64

	switch activityType {
	case "Бег":
		calories, err = RunningSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", err
		}
	case "Ходьба":
		calories, err = WalkingSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", err
		}
	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}

	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f", activityType, durationHours, distanceKm, meanSpeed, calories), nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps == 0 {
		return 0, fmt.Errorf("количество шагов должно быть больше 0")
	}

	if weight == 0 {
		return 0, fmt.Errorf("вес должен быть больше 0")
	}

	if height == 0 {
		return 0, fmt.Errorf("рост должен быть больше 0")
	}

	if duration == 0 {
		return 0, fmt.Errorf("продолжительность должна быть больше 0")
	}

	meanSpeed := meanSpeed(steps, height, duration)
	durationMinutes := duration.Minutes()

	return weight * meanSpeed * durationMinutes / minInH, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps == 0 {
		return 0, fmt.Errorf("количество шагов должно быть больше 0")
	}

	if weight == 0 {
		return 0, fmt.Errorf("вес должен быть больше 0")
	}

	if height == 0 {
		return 0, fmt.Errorf("рост должен быть больше 0")
	}

	if duration == 0 {
		return 0, fmt.Errorf("продолжительность должна быть больше 0")
	}

	meanSpeed := meanSpeed(steps, height, duration)
	durationMinutes := duration.Minutes()

	return weight * meanSpeed * durationMinutes / minInH * walkingCaloriesCoefficient, nil
}
