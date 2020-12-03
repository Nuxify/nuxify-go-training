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
	"api-term/module/academicyear/application"
	serviceTypes "api-term/module/academicyear/infrastructure/service/types"
	types "api-term/module/academicyear/interfaces/http"
)

// AcademicYearCommandController handles the rest api academicYear command requests
type AcademicYearCommandController struct {
	application.AcademicYearCommandServiceInterface
}

// CreateAcademicYear invokes the create academicYear service
func (controller *AcademicYearCommandController) CreateAcademicYear(w http.ResponseWriter, r *http.Request) {
	var academicYear serviceTypes.CreateAcademicYear

	_, claims, _ := jwtauth.FromContext(r.Context())
	academicYear.TenantID = claims["tenant_id"].(string)
	academicYear.CreatedBy = claims["user_id"].(string)
	academicYear.UpdatedBy = claims["user_id"].(string)

	var request types.CreateAcademicYearRequest
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

	academicYear.Name = request.Name
	academicYear.DisplayName = request.DisplayName
	academicYear.IsActive = request.IsActive

	res, err := controller.AcademicYearCommandServiceInterface.CreateAcademicYear(context.TODO(), academicYear)
	if err != nil {
		var httpCode int
		var errorMsg string

		switch err.Error() {
		case errors.DatabaseError:
			httpCode = http.StatusInternalServerError
			errorMsg = "Error occurred while saving user."
		case errors.DuplicateRecord:
			httpCode = http.StatusConflict
			errorMsg = "AcademicYear code already exist."
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
		Message: "Academic Year successfully created.",
		Data: &types.CreateAcademicYearResponse{
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

// DeleteAcademicYearByID delete academicYear by academicYear id
func (controller *AcademicYearCommandController) DeleteAcademicYearByID(w http.ResponseWriter, r *http.Request) {
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

	err = controller.AcademicYearCommandServiceInterface.DeleteAcademicYearByID(int64(id))
	if err != nil {
		response := viewmodels.HTTPResponseVM{
			Status:    http.StatusInternalServerError,
			Success:   false,
			Message:   "An error occurred while deleting academicYear.",
			ErrorCode: err.Error(),
		}

		response.JSON(w)
		return
	}

	response := viewmodels.HTTPResponseVM{
		Status:  http.StatusOK,
		Success: true,
		Message: "Academic year successfully deleted.",
	}

	response.JSON(w)
}

// UpdateAcademicYearByID invokes the create academicYear service
func (controller *AcademicYearCommandController) UpdateAcademicYearByID(w http.ResponseWriter, r *http.Request) {
	var academicYear serviceTypes.UpdateAcademicYear

	_, claims, _ := jwtauth.FromContext(r.Context())
	academicYear.TenantID = claims["tenant_id"].(string)
	academicYear.UpdatedBy = claims["user_id"].(string)

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

	academicYear.ID = int64(id)

	var request types.UpdateAcademicYearRequest
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

	academicYear.Name = request.Name
	academicYear.DisplayName = request.DisplayName
	academicYear.IsActive = request.IsActive

	res, err := controller.AcademicYearCommandServiceInterface.UpdateAcademicYearByID(context.TODO(), academicYear)
	if err != nil {
		var httpCode int
		var errorMsg string

		switch err.Error() {
		case errors.DatabaseError:
			httpCode = http.StatusInternalServerError
			errorMsg = "Error occurred while updating academicYear."
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
		Message: "Academic year successfully updated.",
		Data: &types.CreateAcademicYearResponse{
			ID: res.ID,
		},
	}

	response.JSON(w)
}
