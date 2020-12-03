package rest

import (
	"context"
	"net/http"

	"github.com/go-chi/jwtauth"

	"api-term/interfaces/http/rest/viewmodels"
	"api-term/internal/errors"
	"api-term/module/academicyear/application"
	serviceTypes "api-term/module/academicyear/infrastructure/service/types"
	types "api-term/module/academicyear/interfaces/http"
)

// AcademicYearQueryController handles the rest requests for academicYear queries
type AcademicYearQueryController struct {
	application.AcademicYearQueryServiceInterface
}

// GetAcademicYears get academicYears
func (controller *AcademicYearQueryController) GetAcademicYears(w http.ResponseWriter, r *http.Request) {
	var academicYear serviceTypes.GetAcademicYear

	_, claims, _ := jwtauth.FromContext(r.Context())
	academicYear.TenantID = claims["tenant_id"].(string)

	res, err := controller.AcademicYearQueryServiceInterface.GetAcademicYears(context.TODO(), academicYear)
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

	var academicYears []types.AcademicYearResponse

	for _, academicYear := range res {
		academicYears = append(academicYears, types.AcademicYearResponse{
			ID:          academicYear.ID,
			Name:        academicYear.Name,
			DisplayName: academicYear.DisplayName,
			IsActive:    academicYear.IsActive,
			CreatedAt:   academicYear.CreatedAt.Unix(),
			UpdatedAt:   academicYear.UpdatedAt.Unix(),
		})
	}
	response := viewmodels.HTTPResponseVM{
		Status:  http.StatusOK,
		Success: true,
		Message: "Successfully fetched academic year data.",
		Data:    academicYears,
	}

	response.JSON(w)
}
