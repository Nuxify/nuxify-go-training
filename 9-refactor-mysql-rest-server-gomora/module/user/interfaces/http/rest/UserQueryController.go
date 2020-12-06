package rest

import (
	"context"
	"net/http"
	"strconv"

	"rest-server/interfaces/http/rest/viewmodels"
	"rest-server/internal/errors"
	"rest-server/module/user/application"
	serviceTypes "rest-server/module/user/infrastructure/service/types"
	types "rest-server/module/user/interfaces/http"

	"github.com/go-chi/chi"
)

// UserQueryController handles the rest requests for user queries
type UserQueryController struct {
	application.UserQueryServiceInterface
}

// GetUsers get user
func (controller *UserQueryController) GetUsers(w http.ResponseWriter, r *http.Request) {
	var user serviceTypes.GetUser

	res, err := controller.UserQueryServiceInterface.GetUsers(context.TODO(), user)
	if err != nil {
		var httpCode int
		var errorMsg string

		switch err.Error() {
		case errors.MissingRecord:
			httpCode = http.StatusNotFound
			errorMsg = "No records found."
		default:
			httpCode = http.StatusInternalServerError
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

	var users []types.UserResponse

	for _, user := range res {
		users = append(users, types.UserResponse{
			ID:            user.ID,
			Email:         user.Email,
			FirstName:     user.FirstName,
			LastName:      user.LastName,
			ContactNumber: user.ContactNumber,
			CreatedAt:     user.CreatedAt.Unix(),
			UpdatedAt:     user.UpdatedAt.Unix(),
		})
	}
	response := viewmodels.HTTPResponseVM{
		Status:  http.StatusOK,
		Success: true,
		Message: "Successfully fetched user data.",
		Data:    users,
	}

	response.JSON(w)
}

// GetUserByID get user
func (controller *UserQueryController) GetUserByID(w http.ResponseWriter, r *http.Request) {
	var user serviceTypes.GetUser

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

	res, err := controller.UserQueryServiceInterface.GetUserByID(context.TODO(), user)
	if err != nil {
		var httpCode int
		var errorMsg string

		switch err.Error() {
		case errors.MissingRecord:
			httpCode = http.StatusNotFound
			errorMsg = "No records found."
		default:
			httpCode = http.StatusInternalServerError
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

	var users []types.UserResponse

	for _, user := range res {
		users = append(users, types.UserResponse{
			ID:            user.ID,
			Email:         user.Email,
			FirstName:     user.FirstName,
			LastName:      user.LastName,
			ContactNumber: user.ContactNumber,
			CreatedAt:     user.CreatedAt.Unix(),
			UpdatedAt:     user.UpdatedAt.Unix(),
		})
	}
	response := viewmodels.HTTPResponseVM{
		Status:  http.StatusOK,
		Success: true,
		Message: "Successfully fetched user data.",
		Data:    users,
	}

	response.JSON(w)
}
