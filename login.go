package main

import (
	"fmt"
)

type Employee struct {
	Name     string
	Username string
	Password string
}

func login(E arrStaffs) (bool, string) {
	employees := E
	loginSuccessful := false
	message := ""
	percobaan, idx_found := 1, 0

	for !loginSuccessful {
		clearScreen()
		fmt.Println(message)
		fmt.Println("-------------------------------------------------")
		fmt.Println("           PORTAL LOGIN DEALER MOBIL")
		fmt.Println("-------------------------------------------------")

		var username string
		var password string

		fmt.Printf("Username (%s, %s): ", employees[0].username, employees[1].username)
		fmt.Scanln(&username)

		fmt.Printf("Password (%s, %s): ", employees[0].password, employees[1].password)
		fmt.Scanln(&password)

		index := 0
		for index < len(employees) && !loginSuccessful {
			if employees[index].username == username && employees[index].password == password {
				loginSuccessful = true
				idx_found = index
			}
			index++
		}

		if !loginSuccessful {
			message = "Login gagal. Silakan coba lagi. Percobaan ke-" + fmt.Sprint(percobaan)
			percobaan++
		}
	}

	return loginSuccessful, employees[idx_found].name
}
