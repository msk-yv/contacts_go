package models

type Contact struct {
	Id    string
	Name  string
	Phone string
	Email string
}

func NewContact(id, name, phone, email string) *Contact {
	return &Contact{id, name, phone, email}
}
