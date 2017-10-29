package main

import (
    "time"
    "fmt"
)

var tasks Tasks

type Task struct {
	Id         int       "json:\"id,omitempty\""
	Name       string    "json:\"name,omitempty\""
    Desc       string    "json:\"description,omitempty\""
    Start      time.Time "json:\"start,omitempty\""
	Status     bool      "json:\"status,omitempty\""
	Due        time.Time "json:\"due,omitempty\""
    Assignee   User      "json:\"assignee,omitempty\""
    Comments   string    "json:\"comments,omitempty\""
    Estimation string    "json:\"estimation,omitempty\""
    Spent      time.Time "json:\"time spent,omitempty\""
    Price      string    "json:\"price,omitempty\""
}

type Tasks []Task

func TaskEdit(nt Task) Task{
    for i,t := range tasks {
		if t.Id == nt.Id { // Get the wanted task
            tasks[i] = nt
		}
	}
    // Return empty task if we can't find it
    return Task{}
}

//this is bad, I don't think it passes race condtions
func TaskCreate(t Task) Task {
    fmt.Println(t)
	currentId += 1
	t.Id = currentId
	tasks = append(tasks, t)
	return t
}

func TaskDestroy(id int) error {
	for i, t := range tasks {
		if t.Id == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Task with id of %d to delete", id)
}

func TaskFind(id int) Task {
	for _, t := range tasks {
		if t.Id == id {
			return t
		}
	}
	// return empty Task if not found
	return Task{}
}
