package repository

import (
	"errors"
	"fmt"
	"strings"

	"api-term/infrastructures/database/mysql/types"
	apiError "api-term/internal/errors"
	"api-term/module/academicyear/domain/entity"
	repositoryTypes "api-term/module/academicyear/infrastructure/repository/types"
)

// AcademicYearCommandRepository handles the database operations
type AcademicYearCommandRepository struct {
	types.MySQLDBHandlerInterface
}

// DeleteAcademicYearByID removes academicYear by id
func (repository *AcademicYearCommandRepository) DeleteAcademicYearByID(academicYearID int64) error {
	academicYear := &entity.AcademicYear{
		ID: academicYearID,
	}

	stmt := fmt.Sprintf("DELETE FROM %s WHERE id=:id", academicYear.GetModelName())
	_, err := repository.MySQLDBHandlerInterface.Execute(stmt, academicYear)
	if err != nil {
		return err
	}

	return nil
}

// InsertAcademicYear creates a new academicYear resource
func (repository *AcademicYearCommandRepository) InsertAcademicYear(data repositoryTypes.CreateAcademicYear) (entity.AcademicYear, error) {
	academicYear := &entity.AcademicYear{
		TenantID:    data.TenantID,
		Name:        data.Name,
		DisplayName: data.DisplayName,
		IsActive:    data.IsActive,
		CreatedBy:   data.CreatedBy,
		UpdatedBy:   data.UpdatedBy,
	}

	// insert academicYear
	stmt := fmt.Sprintf("INSERT INTO %s (tenant_id,name,display_name,is_active,created_by,updated_by) "+
		"VALUES (:tenant_id,:name,:display_name,:is_active,:created_by,:updated_by)", academicYear.GetModelName())
	res, err := repository.MySQLDBHandlerInterface.Execute(stmt, academicYear)
	if err != nil {
		var errStr string

		if strings.Contains(err.Error(), "Duplicate entry") {
			errStr = apiError.DuplicateEmail
		} else {
			errStr = apiError.DatabaseError
		}

		return *academicYear, errors.New(errStr)
	}

	// check active status
	if academicYear.IsActive {
		id, err := res.LastInsertId()
		if err != nil {
			return *academicYear, errors.New(apiError.DatabaseError)
		}

		err = repository.updateActiveStatus(academicYear.TenantID, id)
		if err != nil {
			return *academicYear, errors.New(apiError.DatabaseError)
		}
	}

	return *academicYear, nil
}

// UpdateAcademicYearByID update resource
func (repository *AcademicYearCommandRepository) UpdateAcademicYearByID(data repositoryTypes.UpdateAcademicYear) (entity.AcademicYear, error) {
	academicYear := &entity.AcademicYear{
		ID:          data.ID,
		Name:        data.Name,
		DisplayName: data.DisplayName,
		IsActive:    data.IsActive,
		UpdatedBy:   data.UpdatedBy,
	}

	stmt := fmt.Sprintf("UPDATE %s SET name=:name,display_name=:display_name,updated_by=:updated_by "+
		"WHERE id=:id", academicYear.GetModelName())
	_, err := repository.MySQLDBHandlerInterface.Execute(stmt, academicYear)
	if err != nil {
		return *academicYear, errors.New(apiError.DatabaseError)
	}

	// check active status
	if academicYear.IsActive {
		err := repository.updateActiveStatus(data.TenantID, academicYear.ID)
		if err != nil {
			return *academicYear, errors.New(apiError.DatabaseError)
		}
	}

	return *academicYear, nil
}

func (repository *AcademicYearCommandRepository) updateActiveStatus(tenantID string, academicYearID int64) error {
	academicYear := &entity.AcademicYear{
		ID:       academicYearID,
		TenantID: tenantID,
	}

	// set academic year to true
	stmt := fmt.Sprintf("UPDATE %s SET is_active=1 "+
		"WHERE id=:id AND tenant_id=:tenant_id", academicYear.GetModelName())
	_, err := repository.MySQLDBHandlerInterface.Execute(stmt, academicYear)
	if err != nil {
		return errors.New(apiError.DatabaseError)
	}

	// set others to false
	stmt = fmt.Sprintf("UPDATE %s SET is_active=0 "+
		"WHERE id<>:id AND tenant_id=:tenant_id", academicYear.GetModelName())
	_, err = repository.MySQLDBHandlerInterface.Execute(stmt, academicYear)
	if err != nil {
		return errors.New(apiError.DatabaseError)
	}

	return nil
}
