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
	"rest-server/module/discussion/application"
	serviceTypes "rest-server/module/discussion/infrastructure/service/types"
	types "rest-server/module/discussion/interfaces/http"
)

// PostCommandController handles the rest api post command requests
type PostCommandController struct {
	application.PostCommandServiceInterface
}

// CommentCommandController handles the rest api comment command requests
type CommentCommandController struct {
	application.CommentCommandServiceInterface
}

// =====================================POST=====================================

// CreatePost invokes the create post service
func (controller *PostCommandController) CreatePost(w http.ResponseWriter, r *http.Request) {
	var post serviceTypes.CreatePost

	var request types.CreatePostRequest
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

	post.AuthorID = request.AuthorID
	post.Content = request.Content

	res, err := controller.PostCommandServiceInterface.CreatePost(context.TODO(), post)
	if err != nil {
		var httpCode int
		var errorMsg string

		switch err.Error() {
		case errors.DatabaseError:
			httpCode = http.StatusInternalServerError
			errorMsg = "Error occurred while saving post."
		case errors.DuplicateRecord:
			httpCode = http.StatusConflict
			errorMsg = "Post code already exist."
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
		Message: "Post successfully created.",
		Data: &types.CreatePostResponse{
			Content:   res.Content,
			CreatedAt: time.Now().Unix(),
			UpdatedAt: time.Now().Unix(),
		},
	}

	response.JSON(w)
}

// DeletePostByID delete post by post id
func (controller *PostCommandController) DeletePostByID(w http.ResponseWriter, r *http.Request) {
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

	err = controller.PostCommandServiceInterface.DeletePostByID(int64(id))
	if err != nil {
		response := viewmodels.HTTPResponseVM{
			Status:    http.StatusInternalServerError,
			Success:   false,
			Message:   "An error occurred while deleting post.",
			ErrorCode: err.Error(),
		}

		response.JSON(w)
		return
	}

	response := viewmodels.HTTPResponseVM{
		Status:  http.StatusOK,
		Success: true,
		Message: "Post successfully deleted.",
	}

	response.JSON(w)
}

// UpdatePostByID invokes the create post service
func (controller *PostCommandController) UpdatePostByID(w http.ResponseWriter, r *http.Request) {
	var post serviceTypes.UpdatePost

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

	var request types.UpdatePostRequest
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

	post.AuthorID = request.AuthorID
	post.Content = request.Content

	res, err := controller.PostCommandServiceInterface.UpdatePostByID(context.TODO(), post)
	if err != nil {
		var httpCode int
		var errorMsg string

		switch err.Error() {
		case errors.DatabaseError:
			httpCode = http.StatusInternalServerError
			errorMsg = "Error occurred while updating post."
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
		Message: "Post successfully updated.",
		Data: &types.UpdatePostResponse{
			ID:      res.ID,
			Content: res.Content,
		},
	}

	response.JSON(w)
}

// =====================================COMMENT=====================================

// CreateComment invokes the create comment service
func (controller *CommentCommandController) CreateComment(w http.ResponseWriter, r *http.Request) {
	var comment serviceTypes.CreateComment

	var request types.CreateCommentRequest
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

	comment.PostID = request.PostID
	comment.AuthorID = request.AuthorID
	comment.Content = request.Content

	res, err := controller.CommentCommandServiceInterface.CreateComment(context.TODO(), comment)
	if err != nil {
		var httpCode int
		var errorMsg string

		switch err.Error() {
		case errors.DatabaseError:
			httpCode = http.StatusInternalServerError
			errorMsg = "Error occurred while saving comment."
		case errors.DuplicateRecord:
			httpCode = http.StatusConflict
			errorMsg = "Comment code already exist."
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
		Message: "Comment successfully created.",
		Data: &types.CreateCommentResponse{
			PostID:    res.PostID,
			Content:   res.Content,
			CreatedAt: time.Now().Unix(),
			UpdatedAt: time.Now().Unix(),
		},
	}

	response.JSON(w)
}

// DeleteCommentByID delete comment by comment id
func (controller *CommentCommandController) DeleteCommentByID(w http.ResponseWriter, r *http.Request) {
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

	err = controller.CommentCommandServiceInterface.DeleteCommentByID(int64(id))
	if err != nil {
		response := viewmodels.HTTPResponseVM{
			Status:    http.StatusInternalServerError,
			Success:   false,
			Message:   "An error occurred while deleting comment.",
			ErrorCode: err.Error(),
		}

		response.JSON(w)
		return
	}

	response := viewmodels.HTTPResponseVM{
		Status:  http.StatusOK,
		Success: true,
		Message: "Comment successfully deleted.",
	}

	response.JSON(w)
}

// UpdateCommentByID invokes the create comment service
func (controller *CommentCommandController) UpdateCommentByID(w http.ResponseWriter, r *http.Request) {
	var comment serviceTypes.UpdateComment

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

	var request types.UpdateCommentRequest
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

	comment.Content = request.Content

	res, err := controller.CommentCommandServiceInterface.UpdateCommentByID(context.TODO(), comment)
	if err != nil {
		var httpCode int
		var errorMsg string

		switch err.Error() {
		case errors.DatabaseError:
			httpCode = http.StatusInternalServerError
			errorMsg = "Error occurred while updating comment."
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
		Message: "Comment successfully updated.",
		Data: &types.UpdateCommentResponse{
			ID:      res.ID,
			Content: res.Content,
		},
	}

	response.JSON(w)
}