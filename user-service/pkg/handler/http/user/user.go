package user

import (
	"encoding/json"
	"net/http"

	"github.com/davecgh/go-spew/spew"
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
	spew.Dump("ttest12312312")
	jsonData, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)

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
	spew.Dump("ttest12312312")
	jsonData, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}
