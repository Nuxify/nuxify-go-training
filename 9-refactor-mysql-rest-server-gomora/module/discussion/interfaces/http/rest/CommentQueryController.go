package rest

import (
	"context"
	"net/http"
	"strconv"

	"rest-server/interfaces/http/rest/viewmodels"
	"rest-server/internal/errors"
	"rest-server/module/discussion/application"
	serviceTypes "rest-server/module/discussion/infrastructure/service/types"
	types "rest-server/module/discussion/interfaces/http"

	"github.com/go-chi/chi"
)

// CommentQueryController handles the rest requests for comment queries
type CommentQueryController struct {
	application.CommentQueryServiceInterface
}

// GetComments get comment
func (controller *CommentQueryController) GetComments(w http.ResponseWriter, r *http.Request) {
	var comment serviceTypes.GetComment

	res, err := controller.CommentQueryServiceInterface.GetComments(context.TODO(), comment)
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

	var comments []types.CommentResponse

	for _, comment := range res {
		comments = append(comments, types.CommentResponse{
			ID:        comment.ID,
			PostID:    comment.PostID,
			Content:   comment.Content,
			CreatedAt: comment.CreatedAt.Unix(),
			UpdatedAt: comment.UpdatedAt.Unix(),
		})
	}
	response := viewmodels.HTTPResponseVM{
		Status:  http.StatusOK,
		Success: true,
		Message: "Successfully fetched comment data.",
		Data:    comments,
	}

	response.JSON(w)
}

// GetCommentByID get comment
func (controller *CommentQueryController) GetCommentByID(w http.ResponseWriter, r *http.Request) {
	var comment serviceTypes.GetComment

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

	comment.ID = int64(id)

	res, err := controller.CommentQueryServiceInterface.GetCommentByID(context.TODO(), comment)
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

	var comments []types.CommentResponse

	for _, comment := range res {
		comments = append(comments, types.CommentResponse{
			ID:        comment.ID,
			Content:   comment.Content,
			CreatedAt: comment.CreatedAt.Unix(),
			UpdatedAt: comment.UpdatedAt.Unix(),
		})
	}
	response := viewmodels.HTTPResponseVM{
		Status:  http.StatusOK,
		Success: true,
		Message: "Successfully fetched comment data.",
		Data:    comments,
	}

	response.JSON(w)
}
