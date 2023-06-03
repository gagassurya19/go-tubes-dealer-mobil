package main

import (
	"fmt"
)

func login(S arrStaffs) (bool, staff) {
	username, password, message := "", "", ""
	percobaan, idx_found, index := 1, 0, 0
	loginSuccessful := false

	for !loginSuccessful {
		clearScreen()
		fmt.Println(message)
		fmt.Println("-------------------------------------------------")
		fmt.Println("           PORTAL LOGIN DEALER MOBIL")
		fmt.Println("-------------------------------------------------")
		fmt.Printf("Username (%s, %s): ", S[0].username, S[1].username)
		fmt.Scanln(&username)
		fmt.Printf("Password (%s, %s): ", S[0].password, S[1].password)
		fmt.Scanln(&password)

		for index < len(S) && !loginSuccessful {
			if S[index].username == username && S[index].password == password {
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

	return loginSuccessful, S[idx_found]
}
