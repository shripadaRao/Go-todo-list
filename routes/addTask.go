package routes

import (
	"github.com/shripadaRao/crud-app/models"
	"github.com/shripadaRao/crud-app/initializers"
	"net/http"
	"fmt"
	"encoding/json"
)
func AddTask(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Content-Type","application/json")
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)
	fmt.Println("data to be added: ", &task)
	initializers.DB.Create(&task)
}
