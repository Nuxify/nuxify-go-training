package repository

import (
	"api-term/module/academicyear/domain/entity"
	repositoryTypes "api-term/module/academicyear/infrastructure/repository/types"
)

// AcademicYearQueryRepositoryInterface holds the methods for the academic year query repository
type AcademicYearQueryRepositoryInterface interface {
	SelectAcademicYears(data repositoryTypes.GetAcademicYear) ([]entity.AcademicYear, error)
	SelectLatestActiveAcademicYear(tenantID string) (entity.AcademicYear, error)
}
