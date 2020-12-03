package service

import (
	"context"

	"api-term/module/semester/domain/entity"
	"api-term/module/semester/domain/repository"
	repositoryTypes "api-term/module/semester/infrastructure/repository/types"
	"api-term/module/semester/infrastructure/service/types"
)

// SemesterCommandService handles the semester command service logic
type SemesterCommandService struct {
	repository.SemesterCommandRepositoryInterface
}

// CreateSemester creates a resource and persist it in repository
func (service *SemesterCommandService) CreateSemester(ctx context.Context, data types.CreateSemester) (entity.Semester, error) {
	var semester repositoryTypes.CreateSemester

	semester.TenantID = data.TenantID
	semester.Name = data.Name
	semester.DisplayName = data.DisplayName
	semester.IsActive = data.IsActive
	semester.CreatedBy = data.CreatedBy
	semester.UpdatedBy = data.UpdatedBy

	res, err := service.SemesterCommandRepositoryInterface.InsertSemester(semester)
	if err != nil {
		return res, err
	}

	return res, nil
}

// DeleteSemesterByID delete semester by semester id
func (service *SemesterCommandService) DeleteSemesterByID(semesterID int64) error {
	err := service.SemesterCommandRepositoryInterface.DeleteSemesterByID(semesterID)
	if err != nil {
		return err
	}

	return nil
}

// UpdateSemesterByID updates the resource and persist it in repository
func (service *SemesterCommandService) UpdateSemesterByID(ctx context.Context, data types.UpdateSemester) (entity.Semester, error) {
	var semester repositoryTypes.UpdateSemester

	semester.ID = data.ID
	semester.TenantID = data.TenantID
	semester.Name = data.Name
	semester.DisplayName = data.DisplayName
	semester.IsActive = data.IsActive
	semester.UpdatedBy = data.UpdatedBy

	res, err := service.SemesterCommandRepositoryInterface.UpdateSemesterByID(semester)
	if err != nil {
		return res, err
	}

	return res, nil
}
