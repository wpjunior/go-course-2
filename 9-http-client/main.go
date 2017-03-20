package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const MyUsername = "wpjunior"

type User struct {
	Name    string `json:"name"`
	Company string `json:"company"`
}

func GetGithubUser(username string) (*User, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s", username)
	req, _ := http.NewRequest(http.MethodGet, url, nil)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	user := &User{}

	err = json.NewDecoder(res.Body).Decode(user)

	return user, nil
}

func main() {
	user, err := GetGithubUser(MyUsername)
	if err == nil {
		log.Printf("Temos o usuário: %#v\n", user)
	} else {
		log.Println("Falha ao buscar usuário: ", err)
	}
}
