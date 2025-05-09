package user

import (
	apperror "ms-practice/user-service/pkg/utils/app_error"
	"ms-practice/user-service/pkg/utils/response"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h *userHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	user := []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}{
		{
			ID:   "123",
			Name: "Test",
		},
		{
			ID:   "123232",
			Name: "Test3",
		},
	}
	response.ResponseWithSuccess(w, user)

}

func (h *userHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// Retrieve the `id` from the route
	id := vars["id"]
	user := struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}{
		ID:   id,
		Name: "Test",
	}
	userId, err := strconv.Atoi(id)
	if err != nil {
		response.ResponseWithError(w, apperror.ErrInternalServer.Origin(err))
	}
	h.userUC.GetUser(r.Context(), int32(userId))
	response.ResponseWithSuccess(w, user)
}
