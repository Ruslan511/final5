package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errors.New("количество шагов тренировки отрицательное значение или равна нулю")
	} else if weight <= 0 {
		return 0, errors.New("вес тренировки отрицательное значение или равна нулю")
	} else if height <= 0 {
		return 0, errors.New("рост отрицательное значение или равна нулю")
	} else if duration <= time.Duration(0) {
		return 0, errors.New("продожительность тренировки отрицательное значение или равна нулю")
	}
	return ((weight * MeanSpeed(steps, height, duration) * duration.Minutes()) / minInH) * walkingCaloriesCoefficient, nil

}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errors.New("количество шагов тренировки отрицательное значение или равна нулю")
	} else if weight <= 0 {
		return 0, errors.New("вес тренировки отрицательное значение или равна нулю")
	} else if height <= 0 {
		return 0, errors.New("рост отрицательное значение или равна нулю")
	} else if duration <= time.Duration(0) {
		return 0, errors.New("продожительность тренировки отрицательное значение или равна нулю")
	}
	return (weight * MeanSpeed(steps, height, duration) * duration.Minutes()) / minInH, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps < 0 || duration <= time.Duration(0) || height <= 0 {
		return 0
	}
	return Distance(steps, height) / duration.Hours()
}

func Distance(steps int, height float64) float64 {
	distance := (height * stepLengthCoefficient * float64(steps)) / mInKm
	return distance
}
