package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/msk-yv/contacts_go/models"
)

var contacts map[string]*models.Contact

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "index", contacts)
}

func writeHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/write.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	t.ExecuteTemplate(w, "write", nil)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/write.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	id := r.FormValue("id")
	contact, found := contacts[id]
	if !found {
		http.NotFound(w, r)
		return
	}

	t.ExecuteTemplate(w, "write", contact)

}

func saveContactHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	name := r.FormValue("name")
	phone := r.FormValue("phone")
	email := r.FormValue("email")

	var contact *models.Contact
	if id != "" {
		contact = contacts[id]
		contact.Name = name
		contact.Phone = phone
		contact.Email = email
	} else {
		id = GenerateId()
		contact := models.NewContact(id, name, phone, email)
		contacts[contact.Id] = contact
	}

	http.Redirect(w, r, "/", 302)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	if id == "" {
		http.NotFound(w, r)
		return
	}

	delete(contacts, id)

	http.Redirect(w, r, "/", 302)
}

func main() {
	fmt.Println("Listening on port :3000")

	contacts = make(map[string]*models.Contact, 0)

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/write", writeHandler)
	http.HandleFunc("/edit", editHandler)
	http.HandleFunc("/delete", deleteHandler)
	http.HandleFunc("/SaveContact", saveContactHandler)

	http.ListenAndServe(":3000", nil)
}
