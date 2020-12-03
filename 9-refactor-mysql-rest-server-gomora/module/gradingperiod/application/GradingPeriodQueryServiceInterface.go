package application

import (
	"context"

	"api-term/module/gradingperiod/domain/entity"
)

// GradingPeriodQueryServiceInterface holds the implementable method for the grading period query service
type GradingPeriodQueryServiceInterface interface {
	GetGradingPeriods(ctx context.Context) ([]entity.GradingPeriod, error)
	GetGradingPeriodByID(ctx context.Context, gradingPeriodID int) (entity.GradingPeriod, error)
}
