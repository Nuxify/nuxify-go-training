package service

import (
	"context"

	"api-term/module/gradingperiod/domain/entity"
	"api-term/module/gradingperiod/domain/repository"
)

// GradingPeriodQueryService handles business logic in the service layer
type GradingPeriodQueryService struct {
	repository.GradingPeriodQueryRepositoryInterface
}

// GetGradingPeriods returns the grading periods
func (service *GradingPeriodQueryService) GetGradingPeriods(ctx context.Context) ([]entity.GradingPeriod, error) {
	res, err := service.GradingPeriodQueryRepositoryInterface.SelectGradingPeriods()
	if err != nil {
		return res, err
	}

	return res, nil
}

// GetGradingPeriodByID returns the grading period record by id
func (service *GradingPeriodQueryService) GetGradingPeriodByID(ctx context.Context, gradingPeriodID int) (entity.GradingPeriod, error) {
	res, err := service.GradingPeriodQueryRepositoryInterface.SelectGradingPeriodByID(gradingPeriodID)
	if err != nil {
		return res, err
	}

	return res, nil
}
