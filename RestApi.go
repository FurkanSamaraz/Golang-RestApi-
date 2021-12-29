package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("index"))
}
func users(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var a string
	if r.Method == "GET" {
		a = "GET içi"
	} else if r.Method == "POST" {
		a = "POST içi"
	} else {
		a = "Değil!"
	}

	w.Write([]byte("users id :=" + id + " Method =" + a))
	//w.Write([]byte("users"))
}

func post(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	category := vars["category"]
	id := vars["id"]

	w.Write([]byte("index" + " " + category + " " + id))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", index)
	r.HandleFunc("/users", users)
	r.HandleFunc("/users/{id}", users)
	r.HandleFunc("/post/{category}/{id}", post)
	http.ListenAndServe(":8080", r)
}
