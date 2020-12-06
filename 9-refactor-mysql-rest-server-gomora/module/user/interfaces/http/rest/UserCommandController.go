package rest

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"

	"rest-server/interfaces/http/rest/viewmodels"
	"rest-server/internal/errors"
	"rest-server/module/user/application"
	serviceTypes "rest-server/module/user/infrastructure/service/types"
	types "rest-server/module/user/interfaces/http"
)

// UserCommandController handles the rest api user command requests
type UserCommandController struct {
	application.UserCommandServiceInterface
}

// CreateUser invokes the create user service
func (controller *UserCommandController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user serviceTypes.CreateUser

	var request types.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		response := viewmodels.HTTPResponseVM{
			Status:    http.StatusUnprocessableEntity,
			Success:   false,
			Message:   "Invalid payload sent.",
			ErrorCode: errors.InvalidRequestPayload,
		}

		response.JSON(w)
		return
	}

	// verify content must not empty
	if len(request.FirstName) == 0 || len(request.LastName) == 0 {
		response := viewmodels.HTTPResponseVM{
			Status:  http.StatusUnprocessableEntity,
			Success: false,
			Message: "User input cannot be empty.",
		}

		response.JSON(w)
		return
	}

	user.Email = request.Email
	user.FirstName = request.FirstName
	user.LastName = request.LastName
	user.ContactNumber = request.ContactNumber

	res, err := controller.UserCommandServiceInterface.CreateUser(context.TODO(), user)
	if err != nil {
		var httpCode int
		var errorMsg string

		switch err.Error() {
		case errors.DatabaseError:
			httpCode = http.StatusInternalServerError
			errorMsg = "Error occurred while saving user."
		case errors.DuplicateRecord:
			httpCode = http.StatusConflict
			errorMsg = "User code already exist."
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
		Status:  http.StatusCreated,
		Success: true,
		Message: "User successfully created.",
		Data: &types.CreateUserResponse{
			ID:            res.ID,
			Email:         res.Email,
			FirstName:     res.FirstName,
			LastName:      res.LastName,
			ContactNumber: res.ContactNumber,
			CreatedAt:     time.Now().Unix(),
			UpdatedAt:     time.Now().Unix(),
		},
	}

	response.JSON(w)
}

// DeleteUserByID delete user by user id
func (controller *UserCommandController) DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response := viewmodels.HTTPResponseVM{
			Status:    http.StatusUnprocessableEntity,
			Success:   false,
			Message:   "Invalid request payload.",
			ErrorCode: err.Error(),
		}

		response.JSON(w)
		return
	}

	err = controller.UserCommandServiceInterface.DeleteUserByID(int64(id))
	if err != nil {
		response := viewmodels.HTTPResponseVM{
			Status:    http.StatusInternalServerError,
			Success:   false,
			Message:   "An error occurred while deleting user.",
			ErrorCode: err.Error(),
		}

		response.JSON(w)
		return
	}

	response := viewmodels.HTTPResponseVM{
		Status:  http.StatusOK,
		Success: true,
		Message: "User successfully deleted.",
	}

	response.JSON(w)
}

// UpdateUserByID invokes the create user service
func (controller *UserCommandController) UpdateUserByID(w http.ResponseWriter, r *http.Request) {
	var user serviceTypes.UpdateUser

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response := viewmodels.HTTPResponseVM{
			Status:    http.StatusUnprocessableEntity,
			Success:   false,
			Message:   "Invalid request payload.",
			ErrorCode: err.Error(),
		}

		response.JSON(w)
		return
	}

	user.ID = int64(id)

	var request types.UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		response := viewmodels.HTTPResponseVM{
			Status:    http.StatusUnprocessableEntity,
			Success:   false,
			Message:   "Invalid payload sent.",
			ErrorCode: errors.InvalidRequestPayload,
		}

		response.JSON(w)
		return
	}

	user.Email = request.Email
	user.FirstName = request.FirstName
	user.LastName = request.LastName
	user.ContactNumber = request.ContactNumber

	res, err := controller.UserCommandServiceInterface.UpdateUserByID(context.TODO(), user)
	if err != nil {
		var httpCode int
		var errorMsg string

		switch err.Error() {
		case errors.DatabaseError:
			httpCode = http.StatusInternalServerError
			errorMsg = "Error occurred while updating user."
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
		Status:  http.StatusCreated,
		Success: true,
		Message: "User year successfully updated.",
		Data: &types.CreateUserResponse{
			ID: res.ID,
		},
	}

	response.JSON(w)
}
