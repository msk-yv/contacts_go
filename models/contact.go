package models

type Contact struct {
	Id    int
	Name  string
	Phone string
	Email string
}

func NewContact(id int, name, phone, email string) *Contact {
	return &Contact{id, name, phone, email}
}
