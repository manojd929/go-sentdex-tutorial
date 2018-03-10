package htmlTemplate

import (
	"fmt"
	"net/http"
	"html/template"
)

type newsAggPage struct {
	Title string
	News string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Whoa Whoa Whoa Go is Awesome</h1>")
}

func newsAggHandler(w http.ResponseWriter, r *http.Request) {
	p := newsAggPage {Title: "Amazing news aggregator", News: "Some news"}
	t, _ := template.ParseFiles("basictemplating.html")
	t.Execute(w, p)
}

func HTMLTemplate() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/agg/", newsAggHandler)
	http.ListenAndServe(":8000", nil)
}