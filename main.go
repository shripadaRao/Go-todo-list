package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
	"github.com/shripadaRao/crud-app/initializers"
	"github.com/shripadaRao/crud-app/models"
	"github.com/shripadaRao/crud-app/routes"
	)

func Initialize() {
	initializers.ConnectToDB()
	initializers.DB.AutoMigrate(&models.Task{})
}


func MuxCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "API is up and running")
}


func main() {
	Initialize()
	router := mux.NewRouter()

	//route handlers - api endpoints
	router.HandleFunc("/mux-check", MuxCheck).Methods("GET")
	router.HandleFunc("/tasks", routes.GetAllTasks).Methods("GET")
	router.HandleFunc("/task/{id}", routes.GetTaskByID).Methods("GET")
	router.HandleFunc("/addTask", routes.AddTask).Methods("POST")
	router.HandleFunc("/updateTask/{id}", routes.UpdateTask).Methods("PUT")
	router.HandleFunc("/deleteTask/{id}", routes.DeleteTask).Methods("DELETE")
	http.ListenAndServe(":8000", router)
	
}