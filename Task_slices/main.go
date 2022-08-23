package main

import "fmt"

func main() {
	friends := []string{
		"Nugraha Muthus",
		"Bayu Fajar Sidik",
		"I Putu Eka",
		"Satrio Utomo",
		"Agus Supriyatna",
		"Barru Kurniawan",
		"Hasanudin",
		"Yudha Nugraha",
		"Eka Muhendra",
		"Fatur Ewing Fadh",
	}

	for index, friend := range friends {
		fmt.Println("friend:", index+1, friend)
	}

}
