package repository

import (
	"errors"
	"fmt"
	"strings"

	"gomora/infrastructures/database/mysql/types"
	apiError "gomora/internal/errors"
	"gomora/module/waitlist/domain/entity"
	repositoryTypes "gomora/module/waitlist/infrastructure/repository/types"
)

// WaitlistCommandRepository handles the waitlist command repository logic
type WaitlistCommandRepository struct {
	types.MySQLDBHandlerInterface
}

// InsertWaitlist creates a new waitlist
func (repository *WaitlistCommandRepository) InsertWaitlist(data repositoryTypes.CreateWaitlist) (entity.Waitlist, error) {
	waitlist := entity.Waitlist{
		Email: data.Email,
	}

	stmt := fmt.Sprintf("INSERT INTO %s (email) VALUES (:email)", waitlist.GetModelName())
	_, err := repository.MySQLDBHandlerInterface.Execute(stmt, waitlist)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return entity.Waitlist{}, errors.New(apiError.DuplicateRecord)
		}
		return entity.Waitlist{}, errors.New(apiError.DatabaseError)
	}

	return waitlist, nil
}
