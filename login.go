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
		fmt.Println("-------------------------------------------------")
		fmt.Println("           PORTAL LOGIN DEALER MOBIL")
		fmt.Println("-------------------------------------------------")
		fmt.Println(message)
		fmt.Printf("Username (%s, %s): ", S[0].username, S[1].username)
		fmt.Scanln(&username)

		if username == "" {
			message = "Username tidak boleh kosong. Silakan coba lagi."
			continue
		}

		fmt.Printf("Password (%s, %s): ", S[0].password, S[1].password)
		fmt.Scanln(&password)

		if password == "" {
			username = ""
			message = "Password tidak boleh kosong. Silakan coba lagi."
			continue
		}

		for index < len(S) && !loginSuccessful {
			if S[index].username == username && S[index].password == password {
				loginSuccessful = true
				idx_found = index
			}
			index++
		}

		if !loginSuccessful {
			username = ""
			password = ""
			index = 0
			message = "Login gagal. Silakan coba lagi. Percobaan ke-" + fmt.Sprint(percobaan)
			percobaan++
		}
	}

	return loginSuccessful, S[idx_found]
}
