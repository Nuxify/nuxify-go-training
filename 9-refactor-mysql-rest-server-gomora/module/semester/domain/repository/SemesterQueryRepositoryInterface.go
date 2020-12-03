package repository

import (
	"api-term/module/semester/domain/entity"
	repositoryTypes "api-term/module/semester/infrastructure/repository/types"
)

// SemesterQueryRepositoryInterface holds the methods for the semester query repository
type SemesterQueryRepositoryInterface interface {
	SelectSemesters(data repositoryTypes.GetSemester) ([]entity.Semester, error)
	SelectLatestActiveSemester(tenantID string) (entity.Semester, error)
}
