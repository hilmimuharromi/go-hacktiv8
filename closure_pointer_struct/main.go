package main

import "fmt"

type Person struct{ name string }

var PrintPerson = func(persons []*Person) {

	for _, friend := range persons {
		fmt.Println("friend:", friend)
	}
}

func main() {

	listPerson := []*Person{
		{name: "Nugraha Muthus"},
		{name: "Bayu Fajar Sidik"},
		{name: "I Putu Eka"},
		{name: "Satrio Utomo"},
		{name: "Agus Supriyatna"},
		{name: "Barru Kurniawan"},
		{name: "Hasanudin"},
		{name: "Yudha Nugraha"},
		{name: "Eka Muhendra"},
		{name: "Fatur Ewing Fadh"},
	}

	PrintPerson(listPerson)

}
