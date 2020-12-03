package repository

import (
	"errors"
	"fmt"
	"log"

	"api-term/infrastructures/database/mysql/types"
	apiError "api-term/internal/errors"
	"api-term/module/gradingperiod/domain/entity"
)

// GradingPeriodQueryRepository handles databas access logic
type GradingPeriodQueryRepository struct {
	types.MySQLDBHandlerInterface
}

// SelectGradingPeriods select a grading period by grading period id
func (repository *GradingPeriodQueryRepository) SelectGradingPeriods() ([]entity.GradingPeriod, error) {
	var gradingPeriod entity.GradingPeriod
	var gradingPeriods []entity.GradingPeriod

	stmt := fmt.Sprintf("SELECT * FROM %s", gradingPeriod.GetModelName())

	err := repository.Query(stmt, map[string]interface{}{}, &gradingPeriods)
	if err != nil {
		return gradingPeriods, errors.New(apiError.DatabaseError)
	} else if len(gradingPeriods) == 0 {
		return gradingPeriods, errors.New(apiError.MissingRecord)
	}

	return gradingPeriods, nil
}

// SelectGradingPeriodByID select a gradingPeriod by gradingPeriod id
func (repository *GradingPeriodQueryRepository) SelectGradingPeriodByID(gradingPeriodID int) (entity.GradingPeriod, error) {
	var gradingPeriod entity.GradingPeriod
	var gradingPeriods []entity.GradingPeriod

	condModel := map[string]interface{}{
		"id": gradingPeriodID,
	}

	stmt := fmt.Sprintf("SELECT * FROM %s WHERE id=:id", gradingPeriod.GetModelName())

	err := repository.Query(stmt, condModel, &gradingPeriods)
	if err != nil {
		log.Println(err)
		return gradingPeriod, errors.New(apiError.DatabaseError)
	} else if len(gradingPeriods) == 0 {
		return gradingPeriod, errors.New(apiError.MissingRecord)
	}

	return gradingPeriods[0], nil
}
