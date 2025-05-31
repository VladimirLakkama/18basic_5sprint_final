package daysteps

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	// TODO: добавить поля
	personaldata.Personal
	Steps    int
	Duration time.Duration
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		err := errors.New("invalid input format: expected 'steps, duration'")
		log.Printf("datastring split error: %v", err)
		return err
	}

	// Преобразование шагов
	//stepStr := strings.TrimSpace(parts[0])
	steps, err := strconv.Atoi(parts[0])

	if err != nil {
		err := errors.New("invalid steps format: " + err.Error())
		log.Printf("datastring steps error: %v", err)
		return err
	}

	if steps <= 0 {
		err := errors.New("steps must be positive")
		log.Printf("steps must be positive: %v", err)
		return err
	}

	ds.Steps = steps

	//Время тренировки
	//durationStr := strings.TrimSpace(parts[1])
	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		log.Printf("datastring duration error: %v", err)
		return errors.New("invalid duration format: " + err.Error())
	}
	if duration <= 0 {
		log.Printf("duration must be positive: %v", err)
		return errors.New("duration must be positive")
	}

	ds.Duration = duration
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	distance := spentenergy.Distance(ds.Steps, ds.Personal.Height)
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Personal.Height, ds.Personal.Weight, ds.Duration)
	if err != nil {
		log.Printf("WalkingSpentCalories error: %v", err)
		return "", fmt.Errorf("Error in WalkingSpentCalories: %w", err)
	}

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, distance, calories), nil

}
