package rest

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"gomora/interfaces/http/rest/viewmodels"
	"gomora/internal/errors"
	apiError "gomora/internal/errors"
	"gomora/module/waitlist/application"
	serviceTypes "gomora/module/waitlist/infrastructure/service/types"
	types "gomora/module/waitlist/interfaces/http"
)

// WaitlistCommandController request controller for waitlist command
type WaitlistCommandController struct {
	application.WaitlistCommandServiceInterface
}

// CreateWaitlist request handler to create Waitlist
func (controller *WaitlistCommandController) CreateWaitlist(w http.ResponseWriter, r *http.Request) {
	var request types.CreateWaitlistRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		response := viewmodels.HTTPResponseVM{
			Status:    http.StatusUnprocessableEntity,
			Success:   false,
			Message:   "Invalid payload sent.",
			ErrorCode: apiError.InvalidPayload,
		}

		response.JSON(w)
		return
	}

	// verify content must not empty
	if len(request.Email) == 0 {
		response := viewmodels.HTTPResponseVM{
			Status:  http.StatusUnprocessableEntity,
			Success: false,
			Message: "Data input cannot be empty.",
		}

		response.JSON(w)
		return
	}

	waitlist := serviceTypes.CreateWaitlist{
		Email: request.Email,
	}

	res, err := controller.WaitlistCommandServiceInterface.CreateWaitlist(context.TODO(), waitlist)
	if err != nil {
		var httpCode int
		var errorMsg string

		switch err.Error() {
		case errors.DatabaseError:
			httpCode = http.StatusInternalServerError
			errorMsg = "Error occurred while saving waitlist."
		case errors.DuplicateRecord:
			httpCode = http.StatusConflict
			errorMsg = "Record Email already exist."
		default:
			httpCode = http.StatusUnprocessableEntity
			errorMsg = "Please contact technical support."
		}

		response := viewmodels.HTTPResponseVM{
			Status:    httpCode,
			Success:   false,
			Message:   errorMsg,
			ErrorCode: err.Error(),
		}

		response.JSON(w)
		return
	}

	response := viewmodels.HTTPResponseVM{
		Status:  http.StatusOK,
		Success: true,
		Message: "Successfully created record.",
		Data: &types.WaitlistResponse{
			Email:     res.Email,
			CreatedAt: time.Now().Unix(),
		},
	}

	response.JSON(w)
}
