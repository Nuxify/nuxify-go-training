package service

import (
	"context"

	"api-term/module/semester/domain/entity"
	"api-term/module/semester/domain/repository"
	repositoryTypes "api-term/module/semester/infrastructure/repository/types"
	"api-term/module/semester/infrastructure/service/types"
)

// SemesterQueryService handles business logic in the service layer
type SemesterQueryService struct {
	repository.SemesterQueryRepositoryInterface
}

// GetLatestActiveSemester returns the latest active semester
func (service *SemesterQueryService) GetLatestActiveSemester(ctx context.Context, tenantID string) (entity.Semester, error) {
	res, err := service.SemesterQueryRepositoryInterface.SelectLatestActiveSemester(tenantID)
	if err != nil {
		return res, err
	}

	return res, nil
}

// GetSemesters returns the semesters
func (service *SemesterQueryService) GetSemesters(ctx context.Context, data types.GetSemester) ([]entity.Semester, error) {
	var semester repositoryTypes.GetSemester

	semester.TenantID = data.TenantID

	res, err := service.SemesterQueryRepositoryInterface.SelectSemesters(semester)
	if err != nil {
		return res, err
	}

	return res, nil
}
