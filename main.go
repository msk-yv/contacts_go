package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var database *sql.DB

type Contact struct {
	Id    int
	Name  string
	Phone string
	Email string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// Parse templates
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	//Get query answer grom db
	rows, err := database.Query("select * from contacts.contacts")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	//set array of Contact structures to give it in template
	contacts := []Contact{}
	for rows.Next() {
		p := Contact{}
		err := rows.Scan(&p.Id, &p.Name, &p.Phone, &p.Email)
		if err != nil {
			fmt.Println(err)
			continue
		}
		contacts = append(contacts, p)
	}
	//render template
	t.ExecuteTemplate(w, "index", contacts)
}

//simple handler to render page for insert new contact
func writeHandler(w http.ResponseWriter, r *http.Request) {
	// Parse templates
	t, err := template.ParseFiles("templates/write.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	//render template
	t.ExecuteTemplate(w, "write", nil)
}

// handler to render page for update  contact
func editHandler(w http.ResponseWriter, r *http.Request) {
	// Parse templates
	t, err := template.ParseFiles("templates/edit.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	//set array of Contact structures to give it in template
	id := r.FormValue("id")
	rows, err := database.Query("select * from contacts.contacts where id = " + id)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	contacts := []Contact{}

	for rows.Next() {
		p := Contact{}
		err := rows.Scan(&p.Id, &p.Name, &p.Phone, &p.Email)
		if err != nil {
			fmt.Println(err)
			continue
		}
		contacts = append(contacts, p)
	}
	//render template
	t.ExecuteTemplate(w, "edit", contacts)

}

// handler to save new or update contact
func saveContactHandler(w http.ResponseWriter, r *http.Request) {
	//get values from form
	id := r.FormValue("id")
	name := r.FormValue("name")
	phone := r.FormValue("phone")
	email := r.FormValue("email")

	//in case when have id - update row, another create a new contact
	if id != "" {
		rows, err := database.Query("UPDATE `contacts`.`contacts` SET `name`='" + name + "', `phone`='" + phone + "', `email`='" + email + "' WHERE `id`= " + id)
		if err != nil {
			log.Println(err)
		}
		defer rows.Close()
	} else {
		rows, err := database.Query("INSERT INTO `contacts`.`contacts` (`name`, `phone`, `email`) VALUES ('" + name + "', '" + phone + "', '" + email + "');")
		if err != nil {
			log.Println(err)
		}
		defer rows.Close()
	}
	//redirect to index page
	http.Redirect(w, r, "/", 302)
}

//delete handler
func deleteHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	if id == "" {
		http.NotFound(w, r)
		return
	}
	rows, err := database.Query("DELETE FROM `contacts`.`contacts` WHERE `id`= " + id)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	//redirect to index page
	http.Redirect(w, r, "/", 302)
}

//handler for update page
func searchHandler(w http.ResponseWriter, r *http.Request) {
	// Parse templates
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	//get search query from form
	toSearch := r.FormValue("srchstr")
	//get query answer from db
	rows, err := database.Query("select * from contacts.contacts where name like '%" + toSearch + "%'")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	//set array of Contact structures to give it in template
	contacts := []Contact{}
	for rows.Next() {
		p := Contact{}
		err := rows.Scan(&p.Id, &p.Name, &p.Phone, &p.Email)
		if err != nil {
			fmt.Println(err)
			continue
		}
		contacts = append(contacts, p)
	}
	//render tamplate
	t.ExecuteTemplate(w, "index", contacts)
}

func main() {
	//starting message
	fmt.Println("Listening on port :3000")
	//connection to db IT NEED TO
	db, err := sql.Open("mysql", "root:1@/contacts")

	if err != nil {
		log.Println(err)
	}
	database = db
	defer db.Close()
	//handler for get file for logo
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("./img/"))))
	//handlers to all actions
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/write", writeHandler)
	http.HandleFunc("/edit", editHandler)
	http.HandleFunc("/delete", deleteHandler)
	http.HandleFunc("/SaveContact", saveContactHandler)
	http.HandleFunc("/Search", searchHandler)
	//server start
	http.ListenAndServe(":3000", nil)
}
