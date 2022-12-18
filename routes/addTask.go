package routes

import (
	"github.com/shripadaRao/crud-app/models"
	"github.com/shripadaRao/crud-app/initializers"
	"net/http"
	// "fmt"
	"encoding/json"
)
func AddTask(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type","application/json")
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)
	result := initializers.DB.Create(&task)
	if result.Error != nil {
		json.NewEncoder(w).Encode(result.Error)
	}
	json.NewEncoder(w).Encode(http.StatusOK)

}
