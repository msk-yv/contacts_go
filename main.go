package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "index", nil)
}

func main() {
	fmt.Printf("Listening on port :3000\n")

	http.HandleFunc("/", indexHandler)

	http.ListenAndServe(":3000", nil)
}

//p1 8:40
