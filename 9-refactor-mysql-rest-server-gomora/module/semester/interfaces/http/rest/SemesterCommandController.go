package rest

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"

	"api-term/interfaces/http/rest/viewmodels"
	"api-term/internal/errors"
	"api-term/module/semester/application"
	serviceTypes "api-term/module/semester/infrastructure/service/types"
	types "api-term/module/semester/interfaces/http"
)

// SemesterCommandController handles the rest api semester command requests
type SemesterCommandController struct {
	application.SemesterCommandServiceInterface
}

// CreateSemester invokes the create semester service
func (controller *SemesterCommandController) CreateSemester(w http.ResponseWriter, r *http.Request) {
	var semester serviceTypes.CreateSemester

	_, claims, _ := jwtauth.FromContext(r.Context())
	semester.TenantID = claims["tenant_id"].(string)
	semester.CreatedBy = claims["user_id"].(string)
	semester.UpdatedBy = claims["user_id"].(string)

	var request types.CreateSemesterRequest
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

	semester.Name = request.Name
	semester.DisplayName = request.DisplayName
	semester.IsActive = request.IsActive

	res, err := controller.SemesterCommandServiceInterface.CreateSemester(context.TODO(), semester)
	if err != nil {
		var httpCode int
		var errorMsg string

		switch err.Error() {
		case errors.DatabaseError:
			httpCode = http.StatusInternalServerError
			errorMsg = "Error occurred while saving user."
		case errors.DuplicateRecord:
			httpCode = http.StatusConflict
			errorMsg = "Semester code already exist."
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
		Data: &types.CreateSemesterResponse{
			ID:          res.ID,
			Name:        res.Name,
			DisplayName: res.DisplayName,
			IsActive:    res.IsActive,
			CreatedAt:   time.Now().Unix(),
			UpdatedAt:   time.Now().Unix(),
		},
	}

	response.JSON(w)
}

// DeleteSemesterByID delete semester by semester id
func (controller *SemesterCommandController) DeleteSemesterByID(w http.ResponseWriter, r *http.Request) {
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

	err = controller.SemesterCommandServiceInterface.DeleteSemesterByID(int64(id))
	if err != nil {
		response := viewmodels.HTTPResponseVM{
			Status:    http.StatusInternalServerError,
			Success:   false,
			Message:   "An error occurred while deleting semester.",
			ErrorCode: err.Error(),
		}

		response.JSON(w)
		return
	}

	response := viewmodels.HTTPResponseVM{
		Status:  http.StatusOK,
		Success: true,
		Message: "Semester successfully deleted.",
	}

	response.JSON(w)
}

// UpdateSemesterByID invokes the create semester service
func (controller *SemesterCommandController) UpdateSemesterByID(w http.ResponseWriter, r *http.Request) {
	var semester serviceTypes.UpdateSemester

	_, claims, _ := jwtauth.FromContext(r.Context())
	semester.TenantID = claims["tenant_id"].(string)
	semester.UpdatedBy = claims["user_id"].(string)

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

	semester.ID = int64(id)

	var request types.UpdateSemesterRequest
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

	semester.Name = request.Name
	semester.DisplayName = request.DisplayName
	semester.IsActive = request.IsActive

	res, err := controller.SemesterCommandServiceInterface.UpdateSemesterByID(context.TODO(), semester)
	if err != nil {
		var httpCode int
		var errorMsg string

		switch err.Error() {
		case errors.DatabaseError:
			httpCode = http.StatusInternalServerError
			errorMsg = "Error occurred while updating semester."
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
		Message: "Semester successfully updated.",
		Data: &types.CreateSemesterResponse{
			ID: res.ID,
		},
	}

	response.JSON(w)
}
