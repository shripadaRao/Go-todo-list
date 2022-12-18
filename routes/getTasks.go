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
}

func GetAllTasksForProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	var tasks []models.Task
	params := mux.Vars(r)
	project := params["project"]
	result := initializers.DB.Where("project = ?",project).Find(&tasks)
	json.NewEncoder(w).Encode(tasks)
	if result.Error != nil {
		return 
	}
}

func GetTaskByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	var task models.Task
	params := mux.Vars(r)
	result := initializers.DB.First(&task, params["id"])
	json.NewEncoder(w).Encode(task)//return either success or error
	if result.Error != nil {
		return 
	}
}