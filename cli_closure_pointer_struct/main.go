package main

import (
	"fmt"
	"os"
	"strconv"
)

type Person struct {
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}

var PrintPerson = func(idx int) {
	listPerson := []*Person{
		{nama: "Nugraha Muthus", alamat: "Jalan Kebon Sirih", pekerjaan: "Developer", alasan: "Nambah Ilmu"},
		{nama: "Bayu Fajar Sidik", alamat: "Jalan Kebon Sirih", pekerjaan: "Developer", alasan: "Nambah Ilmu"},
		{nama: "I Putu Eka", alamat: "Jalan Kebon Sirih", pekerjaan: "Developer", alasan: "Nambah Ilmu"},
		{nama: "Satrio Utomo", alamat: "Jalan Kebon Sirih", pekerjaan: "Developer", alasan: "Nambah Ilmu"},
		{nama: "Agus Supriyatna", alamat: "Jalan Kebon Sirih", pekerjaan: "Developer", alasan: "Nambah Ilmu"},
		{nama: "Barru Kurniawan", alamat: "Jalan Kebon Sirih", pekerjaan: "Developer", alasan: "Nambah Ilmu"},
		{nama: "Hasanudin", alamat: "Jalan Kebon Sirih", pekerjaan: "Developer", alasan: "Nambah Ilmu"},
		{nama: "Yudha Nugraha", alamat: "Jalan Kebon Sirih", pekerjaan: "Developer", alasan: "Nambah Ilmu"},
		{nama: "Eka Muhendra", alamat: "Jalan Kebon Sirih", pekerjaan: "Developer", alasan: "Nambah Ilmu"},
		{nama: "Fatur Ewing Fadh", alamat: "Jalan Kebon Sirih", pekerjaan: "Developer", alasan: "Nambah Ilmu"},
	}

	if idx > len(listPerson) {
		fmt.Println("Person is not found")
		return
	}

	fmt.Println("nama:", listPerson[idx].nama)
	fmt.Println("alamat:", listPerson[idx].alamat)
	fmt.Println("Pekerjaan:", listPerson[idx].pekerjaan)
	fmt.Println("alasan:", listPerson[idx].alasan)

}

func main() {
	idx := os.Args[1]
	idxString, _ := strconv.Atoi(idx)
	PrintPerson(idxString)
}
