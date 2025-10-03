package spentcalories

import (
	"time"
	"strconv"
	"strings"
	"errors"
	"log"
	"fmt"
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
	// TODO: реализовать функцию
    parts := strings.Split(data, ",")// Алгоритм 1.
    if len(parts) != 3 {
        return 0, "", 0, errors.New("Некорректный формат данных")
    }
    steps, err := strconv.Atoi(parts[0])// Алгоритм 2.
    if err != nil {
        return 0, "", 0, fmt.Errorf("Ошибка парсинга шагов: %w", err)
    }
    if steps <= 0 {
        return 0, "", 0, errors.New("Кол-во шагов должно быть больше 0")
    }
    duration, err := time.ParseDuration(parts[2])// Алгоритм 3.
    if err != nil {
        return 0, "", 0, fmt.Errorf("Ошибка парсинга длительности: %w", err)
    }
    return steps, parts[1], duration, nil
}


func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
    distM := float64(steps) * lenStep
    return distM / mInKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	distKm := distance(steps, height)
    hours := duration.Hours()
    if hours == 0 {
        return 0
    }
    return distKm / hours
}


func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
    steps, activityType, duration, err := parseTraining(data)// Алгоритм 1.
    if err != nil {
        log.Println(err)
	return "", err
    }

    var calories float64// Алгоритм 2.
    switch strings.ToLower(activityType) {
    case "ходьба":
        calories, err = WalkingSpentCalories(steps, weight, height, duration)
        if err != nil {
            log.Println(err)
            return "", err
        }
    case "бег":
        calories, err = RunningSpentCalories(steps, weight, height, duration)
        if err != nil {
            log.Println(err)
            return "", err
        }
    default:
        return "", fmt.Errorf("Неизвестный тип тренировки: %s", activityType)
    }

    dist := distance(steps, height)
    speed := meanSpeed(steps, height, duration)

    result := fmt.Sprintf(// Алгоритм 3.
        "Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f",
        activityType, duration.Hours(), dist, speed, calories,
    )

    return result, nil	
}


func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
        return 0, errors.New("Некорректные параметры")
    }
    speed := meanSpeed(steps, height, duration)
    durationMin := duration.Minutes()
    calories := weight * speed * durationMin / minInH
    return calories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
    if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 { // Алгоритм 1.
        return 0, errors.New("некорректные параметры")
    }

    speed := meanSpeed(steps, height, duration) // Алгоритм 2.

    durationMin := duration.Minutes() // Алгоритм 3.

    calories := weight * speed * durationMin / minInH // Алгоритм 4.

    calories *= walkingCaloriesCoefficient // Алгоритм 5.

    return calories, nil
}
