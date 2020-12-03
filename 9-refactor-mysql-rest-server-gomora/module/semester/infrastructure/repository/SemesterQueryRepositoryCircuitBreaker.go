package repository

import (
	"github.com/afex/hystrix-go/hystrix"

	"api-term/module/semester/domain/entity"
	"api-term/module/semester/domain/repository"
	repositoryTypes "api-term/module/semester/infrastructure/repository/types"
)

// SemesterQueryRepositoryCircuitBreaker is the circuit breaker for the semester query repository
type SemesterQueryRepositoryCircuitBreaker struct {
	repository.SemesterQueryRepositoryInterface
}

// SelectSemesters is a decorator for the select semesters repository
func (repository *SemesterQueryRepositoryCircuitBreaker) SelectSemesters(data repositoryTypes.GetSemester) ([]entity.Semester, error) {
	output := make(chan []entity.Semester, 1)
	hystrix.ConfigureCommand("select_semesters", config.Settings())
	errors := hystrix.Go("select_semesters", func() error {
		semesters, err := repository.SemesterQueryRepositoryInterface.SelectSemesters(data)
		if err != nil {
			return err
		}

		output <- semesters
		return nil
	}, nil)

	select {
	case out := <-output:
		return out, nil
	case err := <-errors:
		return []entity.Semester{}, err
	}
}

// SelectLatestActiveSemester is a decorator for the select semesters repository
func (repository *SemesterQueryRepositoryCircuitBreaker) SelectLatestActiveSemester(tenantID string) (entity.Semester, error) {
	output := make(chan entity.Semester, 1)
	hystrix.ConfigureCommand("select_latest_active_semester", config.Settings())
	errors := hystrix.Go("select_latest_active_semester", func() error {
		semester, err := repository.SemesterQueryRepositoryInterface.SelectLatestActiveSemester(tenantID)
		if err != nil {
			return err
		}

		output <- semester
		return nil
	}, nil)

	select {
	case out := <-output:
		return out, nil
	case err := <-errors:
		return entity.Semester{}, err
	}
}
