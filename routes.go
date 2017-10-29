package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		HandlerIndex,
	},

	Route{ // List of all Tasks
		"TaskIndex",
		"GET",
		"/tasks",
		HandlerTasksIndex,
	},
	Route{ // Create new Task
		"TaskCreate",
		"PUT",
		"/tasks",
		HandlerTaskCreate,
	},
	Route{ // Show single task
		"TaskShow",
		"GET",
		"/tasks/{taskId}",
		HandlerTaskShow,
	},
	Route{ // Delete Task
		"TaskDelete",
		"DELETE",
		"/tasks/{taskId}",
		HandlerTaskDestroy,
	},
}
