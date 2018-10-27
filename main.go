package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	//"github.com/msk-yv/contacts_go/models"
)

var database *sql.DB

type Contact struct {
	Id    int
	Name  string
	Phone string
	Email string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	rows, err := database.Query("select * from contacts.contacts")
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
	t, err := template.ParseFiles("templates/edit.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	//contacts := []Contact{}
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
	/*
		contact, found := contacts[id]
		if !found {
			http.NotFound(w, r)
			return
		}
	*/
	t.ExecuteTemplate(w, "edit", contacts)

}

func saveContactHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	name := r.FormValue("name")
	phone := r.FormValue("phone")
	email := r.FormValue("email")

	//var contact *models.Contact
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

	http.Redirect(w, r, "/", 302)
}

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

	http.Redirect(w, r, "/", 302)
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	toSearch := r.FormValue("srchstr")
	fmt.Println("select * from contacts.contacts where `name` like '%" + toSearch + "%'")
	rows, err := database.Query("select * from contacts.contacts where name like '%" + toSearch + "%'")
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

	t.ExecuteTemplate(w, "index", contacts)
}

func main() {
	fmt.Println("Listening on port :3000")

	db, err := sql.Open("mysql", "root:1@/contacts")

	if err != nil {
		log.Println(err)
	}
	database = db
	defer db.Close()

	//contacts = make(map[string]*Contact, 0)

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/write", writeHandler)
	http.HandleFunc("/edit", editHandler)
	http.HandleFunc("/delete", deleteHandler)
	http.HandleFunc("/SaveContact", saveContactHandler)
	http.HandleFunc("/Search", searchHandler)

	http.ListenAndServe(":3000", nil)
}
