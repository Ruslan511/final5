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

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	splitStrings := strings.Split(datastring, ",")
	if len(splitStrings) != 2 {
		return errors.New("количество элементов в datastring более 2х")
	}
	steps, err := strconv.Atoi(splitStrings[0])
	if err != nil {
		return err
	}
	if steps <= 0 {
		return errors.New("количество шагов отрицательное значение или равно нулю")
	}
	duration, err := time.ParseDuration(splitStrings[1])
	if err != nil {
		return err
	}
	if duration <= 0 {
		return errors.New("продолжительность отрицательное значение или равно нулю")
	}
	ds.Steps = steps
	ds.Duration = duration
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {

	distanceKm := spentenergy.Distance(ds.Steps, ds.Height)
	spentCalories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, distanceKm, spentCalories), nil
}
