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

	const CLR_R = "\x1b[31;1m"
	const CLR_G = "\x1b[32;1m"
	const CLR_N = "\x1b[0m"

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
		fmt.Printf("\n%12s%12s\t", user.Username, user.Port)
		resp, err := http.Get(url + user.Port)
		if err != nil {
			fmt.Printf("%s%s%s", CLR_R, "Fail", CLR_N)
		} else {
			fmt.Printf("%s%s%s", CLR_G, resp.Status, CLR_N)
		}
	}
	fmt.Println()
}
