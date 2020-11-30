package application

import (
	"context"

	"gomora/module/waitlist/domain/entity"
	"gomora/module/waitlist/infrastructure/service/types"
)

// WaitlistCommandServiceInterface holds the implementable methods for the waitlist command service
type WaitlistCommandServiceInterface interface {
	CreateWaitlist(ctx context.Context, data types.CreateWaitlist) (entity.Waitlist, error)
}
