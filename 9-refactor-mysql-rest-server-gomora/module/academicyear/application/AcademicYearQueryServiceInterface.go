package application

import (
	"context"

	"api-term/module/academicyear/domain/entity"
	"api-term/module/academicyear/infrastructure/service/types"
)

// AcademicYearQueryServiceInterface holds the implementable method for the academicyear query service
type AcademicYearQueryServiceInterface interface {
	GetAcademicYears(ctx context.Context, data types.GetAcademicYear) ([]entity.AcademicYear, error)
	GetLatestActiveAcademicYear(ctx context.Context, tenantID string) (entity.AcademicYear, error)
}
