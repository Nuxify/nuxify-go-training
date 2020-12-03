package application

import (
	"context"

	"api-term/module/academicyear/domain/entity"
	"api-term/module/academicyear/infrastructure/service/types"
)

// AcademicYearCommandServiceInterface holds the implementable method for the academicyear command service
type AcademicYearCommandServiceInterface interface {
	CreateAcademicYear(ctx context.Context, data types.CreateAcademicYear) (entity.AcademicYear, error)
	DeleteAcademicYearByID(academicyearID int64) error
	UpdateAcademicYearByID(ctx context.Context, data types.UpdateAcademicYear) (entity.AcademicYear, error)
}
