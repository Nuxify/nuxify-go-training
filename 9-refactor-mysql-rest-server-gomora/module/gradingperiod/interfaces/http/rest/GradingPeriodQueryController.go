package rest

import (
	"context"
	"net/http"

	"api-term/interfaces/http/rest/viewmodels"
	"api-term/internal/errors"
	"api-term/module/gradingperiod/application"
	types "api-term/module/gradingperiod/interfaces/http"
)

// GradingPeriodQueryController handles the rest requests for gradingPeriod queries
type GradingPeriodQueryController struct {
	application.GradingPeriodQueryServiceInterface
}

// GetGradingPeriods get gradingPeriods
func (controller *GradingPeriodQueryController) GetGradingPeriods(w http.ResponseWriter, r *http.Request) {
	res, err := controller.GradingPeriodQueryServiceInterface.GetGradingPeriods(context.TODO())
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

	var gradingPeriods []types.GradingPeriodResponse

	for _, gradingPeriod := range res {
		gradingPeriods = append(gradingPeriods, types.GradingPeriodResponse{
			ID:          gradingPeriod.ID,
			Name:        gradingPeriod.Name,
			DisplayName: gradingPeriod.DisplayName,
			CreatedAt:   gradingPeriod.CreatedAt.Unix(),
			UpdatedAt:   gradingPeriod.UpdatedAt.Unix(),
		})
	}
	response := viewmodels.HTTPResponseVM{
		Status:  http.StatusOK,
		Success: true,
		Message: "Successfully fetched grading periods.",
		Data:    gradingPeriods,
	}

	response.JSON(w)
}
