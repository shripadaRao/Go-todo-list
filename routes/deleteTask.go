package routes
import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/shripadaRao/crud-app/initializers"
	"github.com/shripadaRao/crud-app/models"
	"fmt"
)

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	initializers.DB.Delete(&models.Task{}, params)
	fmt.Println("deleted record")
}