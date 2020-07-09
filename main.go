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
	r.HandleFunc("/crypto", scripts.GetCrypto)
	http.ListenAndServe(port, r)
}
