package routes
import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/shripadaRao/crud-app/initializers"
	"github.com/shripadaRao/crud-app/models"
	// "fmt"
	"encoding/json"
)

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	result := initializers.DB.Delete(&models.Task{}, params)//how to get each parameter independently
	if result.Error != nil {
		json.NewEncoder(w).Encode(result.Error)
	}
	json.NewEncoder(w).Encode(http.StatusOK)	//msg deleted
}