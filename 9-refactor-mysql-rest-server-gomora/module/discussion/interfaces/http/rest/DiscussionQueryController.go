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

// PostQueryController handles the rest requests for post queries
type PostQueryController struct {
	application.PostQueryServiceInterface
}

// CommentQueryController handles the rest requests for comment queries
type CommentQueryController struct {
	application.CommentQueryServiceInterface
}

// ======================================POST======================================

// GetPosts get post
func (controller *PostQueryController) GetPosts(w http.ResponseWriter, r *http.Request) {
	res, err := controller.PostQueryServiceInterface.GetPosts(context.TODO())
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

	var posts []types.PostResponse

	for _, post := range res {
		posts = append(posts, types.PostResponse{
			Content:   post.Content,
			CreatedAt: post.CreatedAt.Unix(),
			UpdatedAt: post.UpdatedAt.Unix(),
		})
	}
	response := viewmodels.HTTPResponseVM{
		Status:  http.StatusOK,
		Success: true,
		Message: "Successfully fetched post data.",
		Data:    posts,
	}

	response.JSON(w)
}

// GetPostByID get post
func (controller *PostQueryController) GetPostByID(w http.ResponseWriter, r *http.Request) {
	var post serviceTypes.GetPost

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

	post.ID = int64(id)

	res, err := controller.PostQueryServiceInterface.GetPostByID(context.TODO(), post)
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

	var posts []types.PostResponse

	for _, post := range res {
		posts = append(posts, types.PostResponse{
			ID:        post.ID,
			Content:   post.Content,
			CreatedAt: post.CreatedAt.Unix(),
			UpdatedAt: post.UpdatedAt.Unix(),
		})
	}
	response := viewmodels.HTTPResponseVM{
		Status:  http.StatusOK,
		Success: true,
		Message: "Successfully fetched post data.",
		Data:    posts,
	}

	response.JSON(w)
}

// ======================================COMMENT======================================

// GetComments get comment
func (controller *CommentQueryController) GetComments(w http.ResponseWriter, r *http.Request) {
	res, err := controller.CommentQueryServiceInterface.GetComments(context.TODO())
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
