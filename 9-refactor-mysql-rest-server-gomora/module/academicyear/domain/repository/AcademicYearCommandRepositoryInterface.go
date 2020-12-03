package repository

import (
	"api-term/module/academicyear/domain/entity"
	repositoryTypes "api-term/module/academicyear/infrastructure/repository/types"
)

// AcademicYearCommandRepositoryInterface holds the implementable methods for the academic year command repository
type AcademicYearCommandRepositoryInterface interface {
	DeleteAcademicYearByID(academicyearID int64) error
	InsertAcademicYear(data repositoryTypes.CreateAcademicYear) (entity.AcademicYear, error)
	UpdateAcademicYearByID(data repositoryTypes.UpdateAcademicYear) (entity.AcademicYear, error)
}
