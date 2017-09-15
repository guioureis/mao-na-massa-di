package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type User struct {
	Username string
	Pubkey   string
	Port     int
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
	fmt.Printf("Value: %#v\n", userList.Users[0].Port)

}
