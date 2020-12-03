package application

import (
	"context"

	"api-term/module/semester/domain/entity"
	"api-term/module/semester/infrastructure/service/types"
)

// SemesterCommandServiceInterface holds the implementable method for the semester command service
type SemesterCommandServiceInterface interface {
	CreateSemester(ctx context.Context, data types.CreateSemester) (entity.Semester, error)
	DeleteSemesterByID(semesterID int64) error
	UpdateSemesterByID(ctx context.Context, data types.UpdateSemester) (entity.Semester, error)
}
