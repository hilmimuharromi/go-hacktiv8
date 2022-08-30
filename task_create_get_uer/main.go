package main

import (
	"fmt"
	"task_get_user/service"
)

func main() {
	var db []*service.User
	userSvc := service.NewUserService(db)
	listUsers := []string{"Bayu", "Ahmed", "Mail", "Bintang"}
	for _, v := range listUsers {
		res := userSvc.Register(&service.User{Nama: v})
		fmt.Println(res, "Berhasil didaftarkan")
	}
	resUsers := userSvc.GetUser()

	fmt.Println("-----------Hasil get user-------------")
	for _, v := range resUsers {
		fmt.Println(v.Nama)
	}

}
