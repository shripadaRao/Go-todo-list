package routes

import (
	"net/http"
	"github.com/shripadaRao/crud-app/initializers"
	"github.com/shripadaRao/crud-app/models"
	"github.com/gorilla/mux"
	"encoding/json"

)
func UpdateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	var task models.Task
	params := mux.Vars(r)
	result := initializers.DB.First(&task, params["id"])
	if result.Error != nil {
		return 
	}

	json.NewDecoder(r.Body).Decode(&task)
	initializers.DB.Save(&task)

	json.NewEncoder(w).Encode(task)
}