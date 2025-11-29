package user

import (
	"errors"
	response "ms-practice/pkg/http/mux"
	apperror "ms-practice/user-service/pkg/utils/app_error"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func (h *userHandler) GetMe(w http.ResponseWriter, r *http.Request) {
	idValue := r.Header.Get("X-User-ID")
	if idValue == "" {
		idValue = r.URL.Query().Get("user_id")
	}
	if idValue == "" {
		response.ResponseWithError(w, apperror.ErrBadRequest.Wrap(errors.New("user id is required")))
		return
	}
	id, err := strconv.Atoi(idValue)
	if err != nil {
		response.ResponseWithError(w, apperror.ErrBadRequest.Wrap(err))
		return
	}

	user, err := h.userUC.GetUser(r.Context(), int32(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.ResponseWithError(w, apperror.ErrNotFound.Wrap(err))
			return
		}
		response.ResponseWithError(w, apperror.ErrInternalServer.Wrap(err))
		return
	}
	response.ResponseWithSuccess(w, user)
}

func (h *userHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idParam := vars["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		response.ResponseWithError(w, apperror.ErrBadRequest.Wrap(err))
		return
	}

	user, err := h.userUC.GetUser(r.Context(), int32(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.ResponseWithError(w, apperror.ErrNotFound.Wrap(err))
			return
		}
		response.ResponseWithError(w, apperror.ErrInternalServer.Wrap(err))
		return
	}
	response.ResponseWithSuccess(w, user)
}
