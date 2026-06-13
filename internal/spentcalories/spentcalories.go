package spentcalories

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	if data == "" {
		return 0, "", 0, errors.New("данные отсутствуют")
	}

	parts := strings.Split(data, ",")
	if len(parts) != 3 {
		return 0, "", 0, errors.New("невалидные данные")
	}

	steps, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil || steps <= 0 {
		return 0, "", 0, errors.New("количество шагов должно быть больше 0")
	}

	duration, err := time.ParseDuration(strings.TrimSpace(parts[2]))
	if err != nil || duration <= 0 {
		return 0, "", 0, errors.New("продолжительность должна быть больше 0")
	}

	activityType := strings.TrimSpace(parts[1])

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
		return "", errors.New("неизвестный тип тренировки")
	}

	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", activityType, durationHours, distanceKm, meanSpeed, calories), nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errors.New("количество шагов должно быть больше 0")
	}

	if weight <= 0 {
		return 0, errors.New("вес должен быть больше 0")
	}

	if height <= 0 {
		return 0, errors.New("рост должен быть больше 0")
	}

	if duration <= 0 {
		return 0, errors.New("продолжительность должна быть больше 0")
	}

	meanSpeed := meanSpeed(steps, height, duration)
	durationMinutes := duration.Minutes()

	return weight * meanSpeed * durationMinutes / minInH, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errors.New("количество шагов должно быть больше 0")
	}

	if weight <= 0 {
		return 0, errors.New("вес должен быть больше 0")
	}

	if height <= 0 {
		return 0, errors.New("рост должен быть больше 0")
	}

	if duration <= 0 {
		return 0, errors.New("продолжительность должна быть больше 0")
	}

	meanSpeed := meanSpeed(steps, height, duration)
	durationMinutes := duration.Minutes()

	return weight * meanSpeed * durationMinutes / minInH * walkingCaloriesCoefficient, nil
}
