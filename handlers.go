package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func HandlerIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func HandlerTasksIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(tasks); err != nil {
		panic(err)
	}
}

func HandlerTaskShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var taskId int
    var err error
	if taskId, err = strconv.Atoi(vars["taskId"]); err != nil {
		panic(err)
	}
	task := TaskFind(taskId)
	if task.Id > 0 {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(task); err != nil {
			panic(err)
		}
		return
	}

	// If we didn't find it, 404
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}

}


/*
Test with this curl command:

curl -X "DELETE" http://localhost:8080/tasks/{ID}

*/
func HandlerTaskDestroy(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var taskId int
	var err error
	if taskId, err = strconv.Atoi(vars["taskId"]); err != nil {
		panic(err)
	}
	task := TaskDestroy(taskId)
	if task == nil{
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
        if err := json.NewEncoder(w).Encode("Delete:true"); err != nil {
			panic(err)
		}
		return
	}

	// If we didn't find it, 404
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}
}

/*
Test with this curl command:

curl -H "Content-Type: application/json" -d '{"name":"New Task"}' http://localhost:8080/tasks

curl -X "PUT" -H "Content-Type: application/json" -d '{"name":"Renamed Task"}' http://localhost:8080/tasks/{ID}

*/
func HandlerTaskCreate(w http.ResponseWriter, r *http.Request) {
	var task Task
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &task); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

    t := Task{}
    if(TaskFind(task.Id) != t){
        t = TaskEdit(task)
    }else{
	    t = TaskCreate(task)
    }
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}
