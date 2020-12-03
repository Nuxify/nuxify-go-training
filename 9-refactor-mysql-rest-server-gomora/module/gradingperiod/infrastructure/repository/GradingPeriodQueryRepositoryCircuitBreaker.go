package repository

import (
	"github.com/afex/hystrix-go/hystrix"

	hystrix_config "api-term/configs/hystrix"
	"api-term/module/gradingperiod/domain/entity"
	"api-term/module/gradingperiod/domain/repository"
)

// GradingPeriodQueryRepositoryCircuitBreaker is the circuit breaker for the gradingPeriod query repository
type GradingPeriodQueryRepositoryCircuitBreaker struct {
	repository.GradingPeriodQueryRepositoryInterface
}

var config = hystrix_config.Config{}

// SelectGradingPeriods is a decorator for the select grading period repository
func (repository *GradingPeriodQueryRepositoryCircuitBreaker) SelectGradingPeriods() ([]entity.GradingPeriod, error) {
	output := make(chan []entity.GradingPeriod, 1)
	hystrix.ConfigureCommand("select_grading_periods", config.Settings())
	errors := hystrix.Go("select_grading_periods", func() error {
		gradingPeriods, err := repository.GradingPeriodQueryRepositoryInterface.SelectGradingPeriods()
		if err != nil {
			return err
		}

		output <- gradingPeriods
		return nil
	}, nil)

	select {
	case out := <-output:
		return out, nil
	case err := <-errors:
		return []entity.GradingPeriod{}, err
	}
}

// SelectGradingPeriodByID is a decorator for the select grading period by id repository
func (repository *GradingPeriodQueryRepositoryCircuitBreaker) SelectGradingPeriodByID(gradingPeriodID int) (entity.GradingPeriod, error) {
	output := make(chan entity.GradingPeriod, 1)
	hystrix.ConfigureCommand("select_grading_period_by_id", config.Settings())
	errors := hystrix.Go("select_grading_period_by_id", func() error {
		gradingPeriod, err := repository.GradingPeriodQueryRepositoryInterface.SelectGradingPeriodByID(gradingPeriodID)
		if err != nil {
			return err
		}

		output <- gradingPeriod
		return nil
	}, nil)

	select {
	case out := <-output:
		return out, nil
	case err := <-errors:
		return entity.GradingPeriod{}, err
	}
}
