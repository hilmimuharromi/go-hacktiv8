package main

import (
	"fmt"
	"os"
	"strconv"
)

type Person struct{ name string }

var PrintPerson = func(idx int) {
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

	if idx > len(listPerson) {
		fmt.Println("Person is not found")
		return
	}

	fmt.Println("person:", listPerson[idx].name)
}

func main() {
	idx := os.Args[1]
	idxString, _ := strconv.Atoi(idx)
	PrintPerson(idxString)
}
