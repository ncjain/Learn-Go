package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type User struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

var users []User

func main() {
	user1 := User{
		ID:       1,
		Username: "Kjain1",
		Email:    "kjain@gmail.com",
	}
	user2 := User{
		ID:       2,
		Username: "Kjain2",
		Email:    "kjain2@gmail.com",
	}
	users = append(users, user1, user2)

	Router := mux.NewRouter()
	Router.HandleFunc("/users", GetUsers).Methods("GET")
	Router.HandleFunc("/users", CreateUser).Methods("POST")
	Router.HandleFunc("/users/{id}", GetUser).Methods("GET")
	Router.HandleFunc("/users/{id}", UpdateUser).Methods("PUT")
	Router.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")
	fmt.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe(":8080", Router))
}

func GetUsers(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&users)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	user := User{}
	_ = json.NewDecoder(r.Body).Decode(&user)

	users = append(users, user)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&user)
}

func GetUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, _ := strconv.ParseUint(vars["id"], 10, 32)
	for _, user := range users {
		if user.ID == id {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(&user)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&User{})

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, _ := strconv.ParseUint(vars["id"], 10, 32)

	uuser := User{}
	_ = json.NewDecoder(r.Body).Decode(&uuser)

	for index, user := range users {
		if user.ID == id {
			users = append(users[:index], users[index+1:]...)
			users = append(users, uuser)
			break
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&users)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, _ := strconv.ParseUint(vars["id"], 10, 32)

	for index, user := range users {
		if user.ID == id {
			users = append(users[:index], users[index+1:]...)
			break
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&users)
}
