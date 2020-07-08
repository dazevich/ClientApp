package scripts

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

//Type Courses ...
type Courses struct {
	Course []struct {
		From      string  `xml:"from" json:"from"`
		To        string  `xml:"to" json:"to"`
		In        float64 `xml:"in" json:"in"`
		Out       float64 `xml:"out" json:"out"`
		Amount    int     `xml:"amount" json:"amount"`
		Minamount string  `xml:"minamount" json:"minamount"`
		Maxamount string  `xml:"maxamount" json:"maxamount"`
		Param     string  `xml:"param" json:"param"`
		City      string  `xml:"city" json:"city"`
	} `xml:"item" json:"courses"`
}

//Function for getting courses
func GetCourses(w http.ResponseWriter, r *http.Request) {

	resp, err := http.Get("http://localhost:9097/getCourses")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	coursesJson := &Courses{}

	dataErr := json.Unmarshal(body, &coursesJson)
	if dataErr != nil {
		log.Fatalln(dataErr)
	}

	tmp, err := template.ParseFiles("static/index.html")
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	if err := tmp.Execute(w, &coursesJson); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

}
