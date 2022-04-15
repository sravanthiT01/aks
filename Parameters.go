package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type User struct {
	Location  string
	Password  string
	Groupname string
	Prefix    string
	ClientId  string
	CreatedAt time.Time
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	user := User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		panic(err)
	}
	user.CreatedAt = time.Now().Local()
	userJson, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = ioutil.WriteFile("parameter_details.json", userJson, 0644)
	check(err)
	//w.Write(userJson)
	w.Write([]byte(`{"status":"OK"}`))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/echo", echoHandler)

	http.ListenAndServe(":8080", mux)
}
