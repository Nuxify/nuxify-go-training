package repository

import (
	"errors"
	"fmt"
	"strings"

	"api-term/infrastructures/database/mysql/types"
	apiError "api-term/internal/errors"
	"api-term/module/semester/domain/entity"
	repositoryTypes "api-term/module/semester/infrastructure/repository/types"
)

// SemesterCommandRepository handles the database operations
type SemesterCommandRepository struct {
	types.MySQLDBHandlerInterface
}

// DeleteSemesterByID removes semester by id
func (repository *SemesterCommandRepository) DeleteSemesterByID(semesterID int64) error {
	semester := &entity.Semester{
		ID: semesterID,
	}

	stmt := fmt.Sprintf("DELETE FROM %s WHERE id=:id", semester.GetModelName())
	_, err := repository.MySQLDBHandlerInterface.Execute(stmt, semester)
	if err != nil {
		return err
	}

	return nil
}

// InsertSemester creates a new semester resource
func (repository *SemesterCommandRepository) InsertSemester(data repositoryTypes.CreateSemester) (entity.Semester, error) {
	semester := &entity.Semester{
		TenantID:    data.TenantID,
		Name:        data.Name,
		DisplayName: data.DisplayName,
		IsActive:    data.IsActive,
		CreatedBy:   data.CreatedBy,
		UpdatedBy:   data.UpdatedBy,
	}

	// insert semester
	stmt := fmt.Sprintf("INSERT INTO %s (id,tenant_id,name,display_name,is_active,created_by,updated_by) "+
		"VALUES (:id,:tenant_id,:name,:display_name,:is_active,:created_by,:updated_by)", semester.GetModelName())
	res, err := repository.MySQLDBHandlerInterface.Execute(stmt, semester)
	if err != nil {
		var errStr string

		if strings.Contains(err.Error(), "Duplicate entry") {
			errStr = apiError.DuplicateEmail
		} else {
			errStr = apiError.DatabaseError
		}

		return *semester, errors.New(errStr)
	}

	// check active status
	if semester.IsActive {
		id, err := res.LastInsertId()
		if err != nil {
			return *semester, errors.New(apiError.DatabaseError)
		}

		err = repository.updateActiveStatus(semester.TenantID, id)
		if err != nil {
			return *semester, errors.New(apiError.DatabaseError)
		}
	}

	return *semester, nil
}

// UpdateSemesterByID update resource
func (repository *SemesterCommandRepository) UpdateSemesterByID(data repositoryTypes.UpdateSemester) (entity.Semester, error) {
	semester := &entity.Semester{
		ID:          data.ID,
		Name:        data.Name,
		DisplayName: data.DisplayName,
		IsActive:    data.IsActive,
		UpdatedBy:   data.UpdatedBy,
	}

	stmt := fmt.Sprintf("UPDATE %s SET name=:name,display_name=:display_name,updated_by=:updated_by "+
		"WHERE id=:id", semester.GetModelName())
	_, err := repository.MySQLDBHandlerInterface.Execute(stmt, semester)
	if err != nil {
		return *semester, errors.New(apiError.DatabaseError)
	}

	// check active status
	if semester.IsActive {
		err := repository.updateActiveStatus(data.TenantID, semester.ID)
		if err != nil {
			return *semester, errors.New(apiError.DatabaseError)
		}
	}

	return *semester, nil
}

func (repository *SemesterCommandRepository) updateActiveStatus(tenantID string, semesterID int64) error {
	semester := &entity.Semester{
		ID:       semesterID,
		TenantID: tenantID,
	}

	// set academic year to true
	stmt := fmt.Sprintf("UPDATE %s SET is_active=1 "+
		"WHERE id=:id AND tenant_id=:tenant_id", semester.GetModelName())
	_, err := repository.MySQLDBHandlerInterface.Execute(stmt, semester)
	if err != nil {
		return errors.New(apiError.DatabaseError)
	}

	// set others to false
	stmt = fmt.Sprintf("UPDATE %s SET is_active=0 "+
		"WHERE id<>:id AND tenant_id=:tenant_id", semester.GetModelName())
	_, err = repository.MySQLDBHandlerInterface.Execute(stmt, semester)
	if err != nil {
		return errors.New(apiError.DatabaseError)
	}

	return nil
}
