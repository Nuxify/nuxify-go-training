package repository

import (
	"api-term/module/gradingperiod/domain/entity"
)

// GradingPeriodQueryRepositoryInterface holds the methods for the grading period query repository
type GradingPeriodQueryRepositoryInterface interface {
	SelectGradingPeriods() ([]entity.GradingPeriod, error)
	SelectGradingPeriodByID(gradingPeriodID int) (entity.GradingPeriod, error)
}
