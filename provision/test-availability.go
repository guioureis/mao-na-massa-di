package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type User struct {
	Username string
	Pubkey   string
	Port     string
}

type UserList struct {
	Users []User
}

func main() {

	filename, _ := filepath.Abs("../users.yml")
	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	var userList UserList

	err = yaml.Unmarshal(yamlFile, &userList)
	if err != nil {
		panic(err)
	}

  url := "http://ec2-18-221-178-11.us-east-2.compute.amazonaws.com:"
	for _, user := range userList.Users {
		fmt.Print("\n" + user.Port + ": ")
		resp, err := http.Get(url + user.Port)
		if err != nil {
			fmt.Print("Fail")
		} else {
			fmt.Print(resp.Status)
		}
	}
	fmt.Println()
}
