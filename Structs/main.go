package main

import (
	"fmt"
)

type contact struct {
	email string
	zip   string
}

type person struct {
	firstname string
	lastname  string
	contact
}

func main() {

	p := person{
		firstname: "xx",
		lastname:  "yy",
		contact: contact{
			email: "z.gmail.com",
			zip:   "123456",
		},
	}
	//pp := &p
	//fmt.Println("Original pointer address of person instance ", &pp)
	p.updateName("aa")
	//fmt.Printf("%+v", p)

	//Map section
	//personMap := make(map[string]person)
	//personMap[p.firstname] = p
	personMap := map[string]string{
		"first name": p.firstname,
		"last name":  p.lastname,
		"email":      p.contact.email,
		"zip":        p.contact.zip,
	}
	print(personMap)
}

func (pp *person) updateName(name string) {
	//fmt.Println("Copied pointer address ", &pp)
	(*pp).firstname = name //* operator of pointer gives the actual struct
}

func print(pm map[string]string) {
	for key, value := range pm {
		fmt.Println(key, "is", value)
	}
}
