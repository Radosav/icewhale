package main

var users Users

type User struct {
	Id         int       "json:\"id,omitempty\""
	Name       string    "json:\"name,omitempty\""
    Desc       string    "json:\"description,omitempty\""
	Status     bool      "json:\"status,omitempty\""
    Password   string    "json:\"password,omitempty\""
    Cost       string    "json:\"cost,omitempty\""
    Mail       string    "json:\"mail,omitempty\""
    Role       string    "json:\"role,omitempty\""
    Position   string    "json:\"position,omitempty\""
    Image      string    "json:\"image,omitempty\"" // URL of Image
}

type Users []User

func UserFind(n string) User{
	for _, u := range users {
		if u.Name == n {
			return u
		}
	}
	// return empty Task if not found
	return User{}
}

//this is bad, I don't think it passes race condtions
func UserCreate(u User) User{
	currentId += 1
	u.Id = currentId
	users = append(users, u)
	return u
}
