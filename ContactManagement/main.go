package main

import "fmt"

type Contact struct {
	ID     int
	Name   string
	Email  string
	Mobile string
}

var contactList []Contact
var contactIndexByName map[string]int
var nextID int = 1

func init() {
	contactList = make([]Contact, 0)
	contactIndexByName = make(map[string]int)
}

func addContact(name, email, mobile string) {
	if _, exist := contactIndexByName[name]; exist {
		fmt.Printf("Contact already exists: %v\n", name)
		return
	}

	newContact := Contact{
		ID:     nextID,
		Name:   name,
		Email:  email,
		Mobile: mobile,
	}
	nextID++

	contactList = append(contactList, newContact)
	contactIndexByName[name] = len(contactList) - 1

	fmt.Printf("Contact added: %v\n", newContact)
}

func findContact(name string) *Contact {
	index, exist := contactIndexByName[name]
	if exist {
		return &contactList[index]
	}
	return nil
}

func listContacts() {
	fmt.Printf("-------- Contacts List ----------\n")

	if len(contactList) == 0 {
		fmt.Printf("No contacts found\n")
		return
	}

	for i, contact := range contactList {
		fmt.Printf("%d. ID: %d, Name: %s, Email: %s, Mobile: %s\n", i+1, contact.ID, contact.Name, contact.Email, contact.Mobile)
	}
	fmt.Printf("")
}

func main() {
	addContact("Esther", "esther@gmail.com", "06034274936")
	addContact("Joshua", "Joshua@gmail.com", "06045274936")
	addContact("Blessing", "Blessing@gmail.com", "09034274936")
	addContact("James", "James@gmail.com", "06034288936")
	addContact("Blessing", "Micheal@gmail.com", "08034274936")
	listContacts()

	bob := findContact("Bob the Developer")
	if bob == nil {
		fmt.Println("Bob contact not found.")
	} else {
		fmt.Println("Bob contact found.")
	}
}
