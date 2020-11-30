package repository

import (
	"github.com/afex/hystrix-go/hystrix"

	hystrix_config "gomora/configs/hystrix"
	"gomora/module/waitlist/domain/entity"
	"gomora/module/waitlist/domain/repository"
	repositoryTypes "gomora/module/waitlist/infrastructure/repository/types"
)

// WaitlistCommandRepositoryCircuitBreaker circuit breaker for Waitlist command repository
type WaitlistCommandRepositoryCircuitBreaker struct {
	repository.WaitlistCommandRepositoryInterface
}

var config = hystrix_config.Config{}

// InsertWaitlist decorator pattern to insert Waitlist
func (repository *WaitlistCommandRepositoryCircuitBreaker) InsertWaitlist(data repositoryTypes.CreateWaitlist) (entity.Waitlist, error) {
	output := make(chan entity.Waitlist, 1)
	hystrix.ConfigureCommand("insert_waitlist", config.Settings())
	errors := hystrix.Go("insert_waitlist", func() error {
		waitlist, err := repository.WaitlistCommandRepositoryInterface.InsertWaitlist(data)
		if err != nil {
			return err
		}

		output <- waitlist
		return nil
	}, nil)

	select {
	case out := <-output:
		return out, nil
	case err := <-errors:
		return entity.Waitlist{}, err
	}
}
