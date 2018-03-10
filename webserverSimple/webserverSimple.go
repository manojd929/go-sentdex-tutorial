package webserverSimple

import (
	"fmt"
	"net/http"
)

func index_handler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "<h1>Index Handler /</h1>")
	// fmt.Fprintf(w, "<p>Paragraph 1 </p>")
	// fmt.Fprintf(w, "<p>Paragraph 2 </p>")

	fmt.Fprintf(w, `
		<h1>Index Handler /</h1>
		<p>Paragraph 1 </p>
		<p>Paragraph 2 </p>
	`)
}

func WebserverSimple() {
	http.HandleFunc("/", index_handler)
	http.ListenAndServe(":8080", nil);
}