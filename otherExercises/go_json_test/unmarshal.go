package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	Id int			`json:"id"`
	Name string		`json:"name"`
}

func main () {
	file, err := os.Open("users.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var users []User

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&users)
	if err != nil {
		panic(err)
	}

	for _, u := range users {
		fmt.Println(u.Name)
	}

}
