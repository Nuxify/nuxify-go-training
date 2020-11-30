package repository

import (
	"gomora/module/waitlist/domain/entity"
	"gomora/module/waitlist/infrastructure/repository/types"
)

// WaitlistCommandRepositoryInterface holds the implementable methods for waitlist command repository
type WaitlistCommandRepositoryInterface interface {
	InsertWaitlist(data types.CreateWaitlist) (entity.Waitlist, error)
}
