package main

import "fmt"

type contactinfo struct {
	city    string
	pincode int
}

type person struct {
	firstname string
	lastname  string
	contact   contactinfo
}

func main() {
	suji := person{
		firstname: "Sujith",
		lastname:  "Varma",
		contact: contactinfo{
			city:    "Salur",
			pincode: 535591,
		},
	}
	//sujiPointer := &suji
	suji.updateFirstName("Suji")
	fmt.Printf("%+v", suji)
}

//works with both types and pointera
func (pointerToPerson *person) updateFirstName(name string) {
	(*pointerToPerson).firstname = name
}
