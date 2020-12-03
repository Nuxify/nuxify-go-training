package service

import (
	"context"

	"api-term/module/academicyear/domain/entity"
	"api-term/module/academicyear/domain/repository"
	repositoryTypes "api-term/module/academicyear/infrastructure/repository/types"
	"api-term/module/academicyear/infrastructure/service/types"
)

// AcademicYearCommandService handles the academicYear command service logic
type AcademicYearCommandService struct {
	repository.AcademicYearCommandRepositoryInterface
}

// CreateAcademicYear creates a resource and persist it in repository
func (service *AcademicYearCommandService) CreateAcademicYear(ctx context.Context, data types.CreateAcademicYear) (entity.AcademicYear, error) {
	var academicYear repositoryTypes.CreateAcademicYear

	academicYear.TenantID = data.TenantID
	academicYear.Name = data.Name
	academicYear.DisplayName = data.DisplayName
	academicYear.IsActive = data.IsActive
	academicYear.CreatedBy = data.CreatedBy
	academicYear.UpdatedBy = data.UpdatedBy

	res, err := service.AcademicYearCommandRepositoryInterface.InsertAcademicYear(academicYear)
	if err != nil {
		return res, err
	}

	return res, nil
}

// DeleteAcademicYearByID delete academicYear by academicYear id
func (service *AcademicYearCommandService) DeleteAcademicYearByID(academicYearID int64) error {
	err := service.AcademicYearCommandRepositoryInterface.DeleteAcademicYearByID(academicYearID)
	if err != nil {
		return err
	}

	return nil
}

// UpdateAcademicYearByID updates the resource and persist it in repository
func (service *AcademicYearCommandService) UpdateAcademicYearByID(ctx context.Context, data types.UpdateAcademicYear) (entity.AcademicYear, error) {
	var academicYear repositoryTypes.UpdateAcademicYear

	academicYear.ID = data.ID
	academicYear.TenantID = data.TenantID
	academicYear.Name = data.Name
	academicYear.DisplayName = data.DisplayName
	academicYear.IsActive = data.IsActive
	academicYear.UpdatedBy = data.UpdatedBy

	res, err := service.AcademicYearCommandRepositoryInterface.UpdateAcademicYearByID(academicYear)
	if err != nil {
		return res, err
	}

	return res, nil
}
