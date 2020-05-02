package main

import (
	"io/ioutil"
	"fmt"
	"net/http"
	"log"
	"encoding/json"
	"strconv"
	"github.com/gorilla/mux"
)

type task struct {
	ID      int    `json:ID`
	Name    string `json:Name`
	Content string `json:Content`
}

type allTasks []task

var tasks = allTasks{
	{
		ID:      1,
		Name:    "Task one",
		Content: "Some Content",
	},
}

func indexRoute(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome")
}

func getTasks(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func createNewTask(w http.ResponseWriter, r *http.Request){
	var newTask task
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a valid task")
	}

	json.Unmarshal(reqBody, &newTask)
	newTask.ID = len(tasks) + 1

	tasks = append(tasks, newTask)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}

func getTaskByID(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
		return
	}

	for _, task := range tasks{
		if task.ID == taskID {

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(task)
	}
	}
}

func deleteTaskByID(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
		return
	}

	for i, task := range tasks{
		if task.ID == taskID {
			tasks= append(tasks[:i], tasks[i + 1:]...)
			fmt.Fprintf(w,"The task %v have been deleted", taskID)
		}
		
	}
}

func updateTaskByID(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])
	var updateTaskByID task 

	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
	}

	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Please enter valid Data")
	}

	json.Unmarshal(reqBody, &updateTaskByID)

	for i,task := range tasks {
		if task.ID == taskID {
			tasks = append(tasks[:i], tasks[i+1:]...)
			updateTaskByID.ID = taskID
			tasks = append(tasks, updateTaskByID)

			fmt.Fprintf(w, "Task ID %v was updted", taskID)
		}
	}
}


func main() {
	router := mux.NewRouter().StrictSlash(true)

	//routes
	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/tasks", getTasks).Methods("GET")
	router.HandleFunc("/tasks", createNewTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", getTaskByID).Methods("GET")
	router.HandleFunc("/tasks/{id}", deleteTaskByID).Methods("DELETE")
	router.HandleFunc("/tasks/{id}", updateTaskByID).Methods("PUT")

	//Init server
	log.Fatal(http.ListenAndServe(":3000", router))

}
