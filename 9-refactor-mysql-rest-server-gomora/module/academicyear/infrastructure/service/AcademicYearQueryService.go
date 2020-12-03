package service

import (
	"context"

	"api-term/module/academicyear/domain/entity"
	"api-term/module/academicyear/domain/repository"
	repositoryTypes "api-term/module/academicyear/infrastructure/repository/types"
	"api-term/module/academicyear/infrastructure/service/types"
)

// AcademicYearQueryService handles business logic in the service layer
type AcademicYearQueryService struct {
	repository.AcademicYearQueryRepositoryInterface
}

// GetAcademicYears returns the academicYears
func (service *AcademicYearQueryService) GetAcademicYears(ctx context.Context, data types.GetAcademicYear) ([]entity.AcademicYear, error) {
	var academicYear repositoryTypes.GetAcademicYear

	academicYear.TenantID = data.TenantID

	res, err := service.AcademicYearQueryRepositoryInterface.SelectAcademicYears(academicYear)
	if err != nil {
		return res, err
	}

	return res, nil
}

// GetLatestActiveAcademicYear returns the latest active academic year
func (service *AcademicYearQueryService) GetLatestActiveAcademicYear(ctx context.Context, tenantID string) (entity.AcademicYear, error) {
	res, err := service.AcademicYearQueryRepositoryInterface.SelectLatestActiveAcademicYear(tenantID)
	if err != nil {
		return res, err
	}

	return res, nil
}
