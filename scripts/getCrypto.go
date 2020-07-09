package scripts

import (
	"html/template"
	"net/http"
)

func GetCrypto(w http.ResponseWriter, r *http.Request) {

	tmp, err := template.ParseFiles("static/crypto.html")
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	if err := tmp.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

}
