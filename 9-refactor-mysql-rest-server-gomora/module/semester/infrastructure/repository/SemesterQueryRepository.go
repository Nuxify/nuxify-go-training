package repository

import (
	"errors"
	"fmt"

	"api-term/infrastructures/database/mysql/types"
	apiError "api-term/internal/errors"
	"api-term/module/semester/domain/entity"
	repositoryTypes "api-term/module/semester/infrastructure/repository/types"
)

// SemesterQueryRepository handles databas access logic
type SemesterQueryRepository struct {
	types.MySQLDBHandlerInterface
}

// SelectSemesters select a semester by semester id
func (repository *SemesterQueryRepository) SelectSemesters(data repositoryTypes.GetSemester) ([]entity.Semester, error) {
	var semester entity.Semester
	var semesters []entity.Semester

	condModel := map[string]interface{}{
		"tenant_id": data.TenantID,
	}

	stmt := fmt.Sprintf("SELECT * FROM %s WHERE tenant_id=:tenant_id", semester.GetModelName())

	err := repository.Query(stmt, condModel, &semesters)
	if err != nil {
		return semesters, errors.New(apiError.DatabaseError)
	} else if len(semesters) == 0 {
		return semesters, errors.New(apiError.MissingRecord)
	}

	return semesters, nil
}

// SelectLatestActiveSemester select latest active semester
func (repository *SemesterQueryRepository) SelectLatestActiveSemester(tenantID string) (entity.Semester, error) {
	var semester entity.Semester
	var semesters []entity.Semester

	condModel := map[string]interface{}{
		"tenant_id": tenantID,
	}

	stmt := fmt.Sprintf("SELECT * FROM %s WHERE tenant_id=:tenant_id AND is_active=1", semester.GetModelName())

	err := repository.Query(stmt, condModel, &semesters)
	if err != nil {
		return semester, errors.New(apiError.DatabaseError)
	} else if len(semesters) == 0 {
		return semester, errors.New(apiError.MissingRecord)
	}

	return semesters[0], nil
}
