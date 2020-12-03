package rest

import (
	"context"
	"net/http"

	"github.com/go-chi/jwtauth"

	"api-term/interfaces/http/rest/viewmodels"
	"api-term/internal/errors"
	"api-term/module/semester/application"
	serviceTypes "api-term/module/semester/infrastructure/service/types"
	types "api-term/module/semester/interfaces/http"
)

// SemesterQueryController handles the rest requests for semester queries
type SemesterQueryController struct {
	application.SemesterQueryServiceInterface
}

// GetSemesters get semesters
func (controller *SemesterQueryController) GetSemesters(w http.ResponseWriter, r *http.Request) {
	var semester serviceTypes.GetSemester

	_, claims, _ := jwtauth.FromContext(r.Context())
	semester.TenantID = claims["tenant_id"].(string)

	res, err := controller.SemesterQueryServiceInterface.GetSemesters(context.TODO(), semester)
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

	var semesters []types.SemesterResponse

	for _, semester := range res {
		semesters = append(semesters, types.SemesterResponse{
			ID:          semester.ID,
			Name:        semester.Name,
			DisplayName: semester.DisplayName,
			IsActive:    semester.IsActive,
			CreatedAt:   semester.CreatedAt.Unix(),
			UpdatedAt:   semester.UpdatedAt.Unix(),
		})
	}
	response := viewmodels.HTTPResponseVM{
		Status:  http.StatusOK,
		Success: true,
		Message: "Successfully fetched semester data.",
		Data:    semesters,
	}

	response.JSON(w)
}
