package application

import (
	"context"

	"api-term/module/semester/domain/entity"
	"api-term/module/semester/infrastructure/service/types"
)

// SemesterQueryServiceInterface holds the implementable method for the semester query service
type SemesterQueryServiceInterface interface {
	GetLatestActiveSemester(ctx context.Context, tenantID string) (entity.Semester, error)
	GetSemesters(ctx context.Context, data types.GetSemester) ([]entity.Semester, error)
}
