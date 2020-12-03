package repository

import (
	"errors"
	"fmt"

	"api-term/infrastructures/database/mysql/types"
	apiError "api-term/internal/errors"
	"api-term/module/academicyear/domain/entity"
	repositoryTypes "api-term/module/academicyear/infrastructure/repository/types"
)

// AcademicYearQueryRepository handles databas access logic
type AcademicYearQueryRepository struct {
	types.MySQLDBHandlerInterface
}

// SelectAcademicYears select a academicYear by academicYear id
func (repository *AcademicYearQueryRepository) SelectAcademicYears(data repositoryTypes.GetAcademicYear) ([]entity.AcademicYear, error) {
	var academicYear entity.AcademicYear
	var academicYears []entity.AcademicYear

	condModel := map[string]interface{}{
		"tenant_id": data.TenantID,
	}

	stmt := fmt.Sprintf("SELECT * FROM %s WHERE tenant_id=:tenant_id", academicYear.GetModelName())

	err := repository.Query(stmt, condModel, &academicYears)
	if err != nil {
		return academicYears, errors.New(apiError.DatabaseError)
	} else if len(academicYears) == 0 {
		return academicYears, errors.New(apiError.MissingRecord)
	}

	return academicYears, nil
}

// SelectLatestActiveAcademicYear select latest active academicYear
func (repository *AcademicYearQueryRepository) SelectLatestActiveAcademicYear(tenantID string) (entity.AcademicYear, error) {
	var academicYear entity.AcademicYear
	var academicYears []entity.AcademicYear

	condModel := map[string]interface{}{
		"tenant_id": tenantID,
	}

	stmt := fmt.Sprintf("SELECT * FROM %s WHERE tenant_id=:tenant_id AND is_active=1", academicYear.GetModelName())

	err := repository.Query(stmt, condModel, &academicYears)
	if err != nil {
		return academicYear, errors.New(apiError.DatabaseError)
	} else if len(academicYears) == 0 {
		return academicYear, errors.New(apiError.MissingRecord)
	}

	return academicYears[0], nil
}
