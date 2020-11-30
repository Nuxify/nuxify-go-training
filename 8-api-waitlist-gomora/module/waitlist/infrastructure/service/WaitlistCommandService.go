package service

import (
	"context"

	"gomora/module/waitlist/domain/entity"
	"gomora/module/waitlist/domain/repository"
	repositoryTypes "gomora/module/waitlist/infrastructure/repository/types"
	"gomora/module/waitlist/infrastructure/service/types"
)

// WaitlistCommandService handles the waitlist command service logic
type WaitlistCommandService struct {
	repository.WaitlistCommandRepositoryInterface
}

// CreateWaitlist create a waitlist
func (service *WaitlistCommandService) CreateWaitlist(ctx context.Context, data types.CreateWaitlist) (entity.Waitlist, error) {
	waitlist := repositoryTypes.CreateWaitlist{
		Email: data.Email,
	}

	res, err := service.WaitlistCommandRepositoryInterface.InsertWaitlist(waitlist)
	if err != nil {
		return entity.Waitlist{}, err
	}

	return res, nil
}
