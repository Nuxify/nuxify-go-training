package repository

import (
	"github.com/afex/hystrix-go/hystrix"

	"api-term/module/academicyear/domain/entity"
	"api-term/module/academicyear/domain/repository"
	repositoryTypes "api-term/module/academicyear/infrastructure/repository/types"
)

// AcademicYearQueryRepositoryCircuitBreaker is the circuit breaker for the academicYear query repository
type AcademicYearQueryRepositoryCircuitBreaker struct {
	repository.AcademicYearQueryRepositoryInterface
}

// SelectAcademicYears is a decorator for the select academicYears repository
func (repository *AcademicYearQueryRepositoryCircuitBreaker) SelectAcademicYears(data repositoryTypes.GetAcademicYear) ([]entity.AcademicYear, error) {
	output := make(chan []entity.AcademicYear, 1)
	hystrix.ConfigureCommand("select_academic_years", config.Settings())
	errors := hystrix.Go("select_academic_years", func() error {
		academicYears, err := repository.AcademicYearQueryRepositoryInterface.SelectAcademicYears(data)
		if err != nil {
			return err
		}

		output <- academicYears
		return nil
	}, nil)

	select {
	case out := <-output:
		return out, nil
	case err := <-errors:
		return []entity.AcademicYear{}, err
	}
}

// SelectLatestActiveAcademicYear is a decorator for the select academicYears repository
func (repository *AcademicYearQueryRepositoryCircuitBreaker) SelectLatestActiveAcademicYear(tenantID string) (entity.AcademicYear, error) {
	output := make(chan entity.AcademicYear, 1)
	hystrix.ConfigureCommand("select_latest_active_academicYear", config.Settings())
	errors := hystrix.Go("select_latest_active_academicYear", func() error {
		academicYear, err := repository.AcademicYearQueryRepositoryInterface.SelectLatestActiveAcademicYear(tenantID)
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
