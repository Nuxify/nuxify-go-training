package repository

import (
	"github.com/afex/hystrix-go/hystrix"

	hystrix_config "api-term/configs/hystrix"
	"api-term/module/academicyear/domain/entity"
	"api-term/module/academicyear/domain/repository"
	repositoryTypes "api-term/module/academicyear/infrastructure/repository/types"
)

// AcademicYearCommandRepositoryCircuitBreaker is the circuit breaker for persistence for the academicYear command repository
type AcademicYearCommandRepositoryCircuitBreaker struct {
	repository.AcademicYearCommandRepositoryInterface
}

var config = hystrix_config.Config{}

// DeleteAcademicYearByID is the decorator for the the academicYear repository delete by id method
func (repository *AcademicYearCommandRepositoryCircuitBreaker) DeleteAcademicYearByID(academicYearID int64) error {
	hystrix.ConfigureCommand("delete_academic_year_by_id", config.Settings())
	errors := hystrix.Go("delete_academic_year_by_id", func() error {
		err := repository.AcademicYearCommandRepositoryInterface.DeleteAcademicYearByID(academicYearID)
		if err != nil {
			return err
		}

		return nil
	}, nil)

	select {
	case err := <-errors:
		return err
	default:
		return nil
	}
}

// InsertAcademicYear is the decorator for the user repository insert academicYear method
func (repository *AcademicYearCommandRepositoryCircuitBreaker) InsertAcademicYear(data repositoryTypes.CreateAcademicYear) (entity.AcademicYear, error) {
	output := make(chan entity.AcademicYear, 1)
	hystrix.ConfigureCommand("insert_academic_year", config.Settings())
	errors := hystrix.Go("insert_academic_year", func() error {
		academicYear, err := repository.AcademicYearCommandRepositoryInterface.InsertAcademicYear(data)
		if err != nil {
			return err
		}

		output <- academicYear
		return nil
	}, nil)

	select {
	case out := <-output:
		return out, nil
	case err := <-errors:
		return entity.AcademicYear{}, err
	}
}

// UpdateAcademicYearByID is the decorator for the academicYear repository update academicYear method
func (repository *AcademicYearCommandRepositoryCircuitBreaker) UpdateAcademicYearByID(data repositoryTypes.UpdateAcademicYear) (entity.AcademicYear, error) {
	output := make(chan entity.AcademicYear, 1)
	hystrix.ConfigureCommand("update_academic_year", config.Settings())
	errors := hystrix.Go("update_academic_year", func() error {
		academicYear, err := repository.AcademicYearCommandRepositoryInterface.UpdateAcademicYearByID(data)
		if err != nil {
			return err
		}

		output <- academicYear
		return nil
	}, nil)

	select {
	case out := <-output:
		return out, nil
	case err := <-errors:
		return entity.AcademicYear{}, err
	}
}
