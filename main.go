package main

import (
	"net/http"

	"./scripts"
	"github.com/gorilla/mux"
)

func main() {
	port := ":8080"
	r := mux.NewRouter()
	r.HandleFunc("/courses", scripts.GetCourses)
	http.ListenAndServe(port, r)
}
