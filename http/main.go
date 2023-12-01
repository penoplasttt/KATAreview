package main

import (
	"encoding/json"
	"net/http"
)

type User struct {
	ID        int     `json:"id"`
	Email     string  `json:"email"`
	Username  string  `json:"username"`
	Password  string  `json:"password"`
	Amount    int     `json:"amount"`
	Profile   Profile `json:"profile"`
	CreatedAt string  `json:"createdAt"`
	CreatedBy string  `json:"createdBy"`
}

type Profile struct {
	Avatar     string `json:"avatar"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	StaticData string `json:"staticData"`
}

func main() {
	
}

func UsersHandler (w http.ResponseWriter, r *http.Request) {
	cli := http.Client{}

	resp, err := cli.Get("https://demo.apistubs.io/api/v1/users")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var users []User

	err = json.NewDecoder(resp.Body).Decode(&users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var modifiedUsers []User

	for _,user := range users {
		user.Profile.StaticData = ""
		user.Password = ""
		if user.Amount > 5000 {
			user.Email = ""
			user.Username = ""
			user.Profile.FirstName = ""
			user.Profile.LastName = ""
			user.Profile.Avatar = ""
		}

		modifiedUsers = append(modifiedUsers, user)
	}

	data, err := json.Marshal(modifiedUsers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)

}