package main

import "fmt"

type contact_info struct {
	email string
	phno  string
}

type person struct {
	firstName string
	lastName  string
	contact_info
}

func main() {
	hadi := person{
		firstName: "Hadi",
		lastName:  "Ahmed",
		contact_info: contact_info{
			email: "969hadiahmed.com", phno: "1234123443",
		}}
	// fmt.Printf("%+v", hadi)
	// hadiPointer := &hadi
	hadi.rename("Abdul Hadi") //interpreted as (&hadi).rename
	hadi.print()
}

func (p *person) rename(fn string) {
	p.firstName = fn // interpreted as (*p).firstName
}
func (p person) print() {
	fmt.Println("Name : " + p.firstName + " " + p.lastName)
	fmt.Println("Email : " + p.contact_info.email)
}
