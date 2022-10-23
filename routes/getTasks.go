package routes

import (
	"net/http"
	// "fmt"
	"github.com/shripadaRao/crud-app/initializers"
	"github.com/shripadaRao/crud-app/models"
	"encoding/json"
	"github.com/gorilla/mux"
)

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	var tasks []models.Task
	result := initializers.DB.Find(&tasks)
	json.NewEncoder(w).Encode(tasks)
	if result.Error != nil {
		return 
	}
	// fmt.Println("no of rows: ",result.RowsAffected)
	// fmt.Println("content: ",tasks)
}

func GetTaskByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	var task models.Task
	params := mux.Vars(r)
	result := initializers.DB.First(&task, params["id"])
	json.NewEncoder(w).Encode(task)
	if result.Error != nil {
		return 
	}
}