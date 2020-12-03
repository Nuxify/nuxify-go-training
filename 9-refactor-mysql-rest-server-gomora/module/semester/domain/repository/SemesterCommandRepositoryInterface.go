package repository

import (
	"api-term/module/semester/domain/entity"
	repositoryTypes "api-term/module/semester/infrastructure/repository/types"
)

// SemesterCommandRepositoryInterface holds the implementable methods for the semester command repository
type SemesterCommandRepositoryInterface interface {
	DeleteSemesterByID(semesterID int64) error
	InsertSemester(data repositoryTypes.CreateSemester) (entity.Semester, error)
	UpdateSemesterByID(data repositoryTypes.UpdateSemester) (entity.Semester, error)
}
