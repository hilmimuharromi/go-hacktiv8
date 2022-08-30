package main

import (
	"fmt"
	"sync"
	"task_get_user/service"
)

func main() {
	var db []*service.User
	userSvc := service.NewUserService(db)
	listUsers := []string{"Bayu", "Ahmed", "Mail", "Bintang", "kadek", "Anjas", "kevin", "hugo", "Syendi", "Nazera"}
	var wg sync.WaitGroup
	wg.Add(len(listUsers))
	for _, v := range listUsers {
		go func(name string) {
			res := userSvc.Register(&service.User{Nama: name})
			fmt.Println(res, "Berhasil didaftarkan")
			wg.Done()
		}(v)
	}
	wg.Wait()

	var wg1 sync.WaitGroup
	resUsers := userSvc.GetUser()
	wg1.Add(len(resUsers))
	fmt.Println("-----------Hasil get user-------------")
	for _, v := range resUsers {
		go func(name string) {
			fmt.Println(name)
			wg1.Done()
		}(v.Nama)
	}
	wg1.Wait()

}
