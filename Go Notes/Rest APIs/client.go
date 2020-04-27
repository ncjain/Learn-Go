package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

/*
url := "http://localhost:8080/users/1"
method := "PUT"

payload := strings.NewReader("")
client := &http.Client{}
req, err := http.NewRequest(method, url, payload)
res, err := client.Do(req)
defer res.Body.Close()
body, err := ioutil.ReadAll(res.Body)
fmt.Println(string(body))

*/

type User1 struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func get() {
	url := "http://localhost:8080/users/1"
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	} else {
		data, _ := ioutil.ReadAll(res.Body)
		fmt.Println(string(data))
	}
}

func post() {
	jsonData := User{
		ID:       4,
		Username: "Kjain4",
		Email:    "Kjain4@gmail.com",
	}
	jsonValue, _ := json.Marshal(jsonData)
	fmt.Println(string(jsonValue))
	url := "http://localhost:8080/users"
	res, err := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
	} else {
		fmt.Println("Sucess")
		data, _ := ioutil.ReadAll(res.Body)
		fmt.Println(string(data))
	}

}

func get_users() {

	url := "http://localhost:8080/users"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
	}
	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
}

func get_user() {

	url := "http://localhost:8080/users/1"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
	}
	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
}

func create_user() {

	url := "http://localhost:8080/users"
	method := "POST"

	//payload := strings.NewReader("{\n    \"id\": 3,\n    \"username\": \"Kjain3\",\n    \"email\": \"kjain3@gmail.com\"\n}")

	payload := strings.NewReader(`
		{
			"id": 3,
			"username": "Kjain3",
			"email": "kjain3@gmail.com"
		}
	`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
}

func delete_user() {

	url := "http://localhost:8080/users/3"
	method := "DELETE"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
	}
	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
}

func update_user() {

	url := "http://localhost:8080/users/1"
	method := "PUT"

	payload := strings.NewReader("{\n    \"id\": 1,\n    \"username\": \"Kjain1\",\n    \"email\": \"kjain1@gmail.com\"\n}")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
}
