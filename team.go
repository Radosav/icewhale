package main

type Team struct {
	Id         int       "json:\"id,omitempty\""
	Name       string    "json:\"name,omitempty\""
    Desc       string    "json:\"description,omitempty\""
	Status     bool      "json:\"status,omitempty\""
    TeamLead   User      "json:\"assignee,omitempty\""
}

type Teams []Team
