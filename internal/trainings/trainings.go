package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	splitStrings := strings.Split(datastring, ",")
	if len(splitStrings) != 3 {
		return errors.New("переданы некорректные данные")
	}
	steps, err := strconv.Atoi(splitStrings[0])
	if err != nil {
		return err
	}
	if steps <= 0 {
		return errors.New("количество шагов отрицательное значение или равно нулю")
	}
	duration, err := time.ParseDuration(splitStrings[2])
	if err != nil {
		return err
	}
	if duration <= time.Duration(0) {
		return errors.New("продолжительность тренировки отрицательное значение или равна нулю")
	}
	if splitStrings[1] == "" {
		return errors.New("не определен тип тренировки")
	}

	t.TrainingType = splitStrings[1]
	t.Duration = duration
	t.Steps = steps
	return nil
}

func (t Training) ActionInfo() (string, error) {
	if t.TrainingType != "Бег" && t.TrainingType != "Ходьба" {
		return "", errors.New("неизвестный тип тренировки")
	}
	distance := spentenergy.Distance(t.Steps, t.Height)
	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)
	var spentCalories float64
	var err error
	switch t.TrainingType {
	case "Бег":
		spentCalories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "Ходьба":
		spentCalories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	}
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, t.Duration.Hours(), distance, meanSpeed, spentCalories), nil
}
