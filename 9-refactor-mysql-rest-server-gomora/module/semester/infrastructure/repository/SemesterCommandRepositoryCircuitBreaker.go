package repository

import (
	"github.com/afex/hystrix-go/hystrix"

	hystrix_config "api-term/configs/hystrix"
	"api-term/module/semester/domain/entity"
	"api-term/module/semester/domain/repository"
	repositoryTypes "api-term/module/semester/infrastructure/repository/types"
)

// SemesterCommandRepositoryCircuitBreaker is the circuit breaker for persistence for the semester command repository
type SemesterCommandRepositoryCircuitBreaker struct {
	repository.SemesterCommandRepositoryInterface
}

var config = hystrix_config.Config{}

// DeleteSemesterByID is the decorator for the the semester repository delete by id method
func (repository *SemesterCommandRepositoryCircuitBreaker) DeleteSemesterByID(semesterID int64) error {
	hystrix.ConfigureCommand("delete_semester_by_id", config.Settings())
	errors := hystrix.Go("delete_semester_by_id", func() error {
		err := repository.SemesterCommandRepositoryInterface.DeleteSemesterByID(semesterID)
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

// InsertSemester is the decorator for the user repository insert semester method
func (repository *SemesterCommandRepositoryCircuitBreaker) InsertSemester(data repositoryTypes.CreateSemester) (entity.Semester, error) {
	output := make(chan entity.Semester, 1)
	hystrix.ConfigureCommand("insert_semester", config.Settings())
	errors := hystrix.Go("insert_semester", func() error {
		semester, err := repository.SemesterCommandRepositoryInterface.InsertSemester(data)
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

// UpdateSemesterByID is the decorator for the semester repository update semester method
func (repository *SemesterCommandRepositoryCircuitBreaker) UpdateSemesterByID(data repositoryTypes.UpdateSemester) (entity.Semester, error) {
	output := make(chan entity.Semester, 1)
	hystrix.ConfigureCommand("update_semester", config.Settings())
	errors := hystrix.Go("update_semester", func() error {
		semester, err := repository.SemesterCommandRepositoryInterface.UpdateSemesterByID(data)
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
