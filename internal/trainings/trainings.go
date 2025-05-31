package trainings

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

type Training struct {
	// TODO: добавить поля
	personaldata.Personal
	Steps        int
	TrainingType string
	Duration     time.Duration
}

func (t *Training) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	parts := strings.Split(datastring, ",")
	if len(parts) != 3 {
		err := errors.New("invalid input format: expected 'steps, training type, duration'")
		log.Printf("datastring error: %v", err)
		return err
	}
	// Преобразование шагов
	stepStr := strings.TrimSpace(parts[0])
	steps, err := strconv.Atoi(stepStr)

	if err != nil {
		err := errors.New("invalid steps format: " + err.Error())
		log.Printf("datastring error: %v", err)
		return err
	}

	if steps <= 0 {
		err := errors.New("steps must be positive")
		log.Printf("steps must be positive: %v", err)
		return err
	}

	t.Steps = steps
	// Тип тренировки
	trainingType := strings.TrimSpace(parts[1])
	if trainingType == "" {
		log.Printf("training type is empty: %v", err)
		return errors.New("training type can't be empty")
	}
	t.TrainingType = trainingType

	//Время тренировки
	durationStr := strings.TrimSpace(parts[2])
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		log.Printf("ParseDuration error: %v", err)
		return errors.New("invalid duration format: " + err.Error())
	}
	if duration <= 0 {
		log.Printf("Durarition not positive: %v", err)
		return errors.New("duration must be positive")
	}

	t.Duration = duration
	return nil
}

func (t Training) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	distance := spentenergy.Distance(t.Steps, t.Personal.Height)
	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Personal.Height, t.Duration)
	energyRun, err := spentenergy.RunningSpentCalories(t.Steps, t.Personal.Height, t.Personal.Weight, t.Duration)
	if err != nil {
		log.Printf("energyRun error: %v", err)
		return "", fmt.Errorf("Error in RunningSpentCalories: %w", err)
	}
	energyWalk, err := spentenergy.WalkingSpentCalories(t.Steps, t.Personal.Height, t.Personal.Weight, t.Duration)
	if err != nil {
		log.Printf("energyRun error: %v", err)
		return "", fmt.Errorf("Error in WalkingSpentCalories: %w", err)
	}

	if t.TrainingType == "Бег" {
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, t.Duration.Hours(), distance, meanSpeed, energyRun), nil
	}
	if t.TrainingType == "Ходьба" {
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, t.Duration.Hours(), distance, meanSpeed, energyWalk), nil
	}
	log.Printf("training type error: %v", err)
	return "", errors.New("неизвестный тип тренировки")

}
