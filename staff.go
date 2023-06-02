package main

import (
	"fmt"
	"strconv"
	"strings"
)

func listStaff(C arrStaffs, n int) {
	var opt int

	for {
		clearScreen()
		fmt.Println("-------------------------------------------------")
		fmt.Println("           LIHAT DAFTAR STAFF")
		fmt.Println("-------------------------------------------------")
		fmt.Println()

		if n == 0 {
			fmt.Println("Tidak ada staff yang terdaftar")
		} else {
			fmt.Println("ID  | Nama       | Telp       | Username   | Password")
			fmt.Println("-----------------------------------------------------")
			for i := 0; i < n; i++ {
				fmt.Printf("%-3d | %-10s | %-10s | %-10s | %-5s\n", C[i].id, C[i].name, C[i].num_telp, C[i].username, strings.Repeat("*", len(C[i].password)))
			}
			fmt.Println("-----------------------------------------------------")
		}

		fmt.Println()
		fmt.Println("1. Sort 2. Kembali")
		fmt.Println("-------------------------------------------------")
		fmt.Print("Pilihan: ")
		fmt.Scanln(&opt)

		switch opt {
		case 1:
			sortStaff(&C, n)
		case 2:
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
			continue
		}
	}
}

func sortStaff(C *arrStaffs, n int) {
	var jenis, sort string
	fmt.Print("Masukkan [id/nama/telp/username] [a/A]:")
	fmt.Scanln(&jenis, &sort)
	switch jenis {
	case "id":
		if sort == "A" {
			sortStaffByIDDesc(C, n)
		} else {
			sortStaffByIDAsc(C, n)
		}
	case "nama":
		if sort == "A" {
			sortStaffByNameDesc(C, n)
		} else {
			sortStaffByNameAsc(C, n)
		}
	case "telp":
		if sort == "A" {
			sortStaffByTelpDesc(C, n)
		} else {
			sortStaffByTelpAsc(C, n)
		}
	case "username":
		if sort == "A" {
			sortStaffByUsernameDesc(C, n)
		} else {
			sortStaffByUsernameAsc(C, n)
		}
	default:
		fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
	}
}

func searchStaff(C *arrStaffs, n *int) {
	var opt int
	var keyword, message string
	var choosenStaff staff

	for {
		clearScreen()
		fmt.Println("-------------------------------------------------")
		fmt.Println("           CARI STAFF")
		fmt.Println("-------------------------------------------------")
		fmt.Println()

		if *n == 0 {
			fmt.Println("Tidak ada staff yang terdaftar")
		} else {
			fmt.Println("ID  | Nama       | Telp       | Username   | Password")
			fmt.Println("-----------------------------------------------------")
			for i := 0; i < *n; i++ {
				fmt.Printf("%-3d | %-10s | %-10s | %-10s | %-5s\n", C[i].id, C[i].name, C[i].num_telp, C[i].username, strings.Repeat("*", len(C[i].password)))
			}
			fmt.Println("-----------------------------------------------------")
		}

		fmt.Println(message)
		fmt.Print("Masukkan id/nama/telp | ` (exit): ")
		fmt.Scanln(&keyword)

		if keyword == "`" {
			return
		} else if keyword == "" {
			message = "**Keyword tidak boleh kosong**"
			continue
		}
		numkeyword, _ := strconv.Atoi(keyword)

		var found bool = false
		for i := 0; i < *n; i++ {
			if C[i].id == numkeyword || C[i].name == keyword || C[i].num_telp == keyword {
				clearScreen()
				fmt.Println("-------------------------------------------------")
				fmt.Println("           CARI STAFF asdf")
				fmt.Println("-------------------------------------------------")
				fmt.Println()
				fmt.Println("ID  | Nama       | Telp       | Username   | Password")
				fmt.Println("-----------------------------------------------------")
				fmt.Printf("%-3d | %-10s | %-10s | %-10s | %-5s\n", C[i].id, C[i].name, C[i].num_telp, C[i].username, strings.Repeat("*", len(C[i].password)))
				fmt.Println("-----------------------------------------------------")
				choosenStaff = C[i]
				found = true
				break
			}
		}

		message = ""
		keyword = ""

		if !found {
			clearScreen()
			fmt.Println("-------------------------------------------------")
			fmt.Println("           CARI STAFF")
			fmt.Println("-------------------------------------------------")
			fmt.Println()
			fmt.Println("Staff tidak ditemukan")
			fmt.Println()
			fmt.Println("-------------------------------------------------")
			fmt.Println("1. Cari lagi | 2. Kembali")
			fmt.Println("-------------------------------------------------")
			fmt.Print("Pilihan: ")
			fmt.Scanln(&opt)
			switch opt {
			case 1:
				continue
			case 2:
				return
			default:
				message = "**Pilihan tidak valid. Silakan coba lagi.**"
				continue
			}
		} else {
			fmt.Println()
			fmt.Println("1. Cari lagi | 2. Edit | 3. Hapus | 4. Kembali")
			fmt.Println("-------------------------------------------------")
			fmt.Print("Pilihan: ")
			fmt.Scanln(&opt)
			switch opt {
			case 1:
				continue
			case 2:
				editStaff(C, choosenStaff, n)
			case 3:
				deleteStaff(C, choosenStaff, n)
			case 4:
				return
			default:
				message = "**Pilihan tidak valid. Silakan coba lagi.**"
				continue
			}
		}
	}
}

func addStaff(C *arrStaffs, n *int) {
	var opt int
	var in staff

	for {
		clearScreen()
		fmt.Println("-------------------------------------------------")
		fmt.Println("           TAMBAH STAFF")
		fmt.Println("-------------------------------------------------")
		fmt.Println()
		fmt.Println("ID  | Nama       | Telp       | Username   | Password")
		fmt.Println("-----------------------------------------------------")
		for i := 0; i < *n; i++ {
			fmt.Printf("%-3d | %-10s | %-10s | %-10s | %-5s\n", C[i].id, C[i].name, C[i].num_telp, C[i].username, strings.Repeat("*", len(C[i].password)))
		}
		fmt.Println("-----------------------------------------------------")
		fmt.Println()
		fmt.Println("-------------------------------------------------")
		fmt.Println("1. Tambah | 2. Kembali")
		fmt.Println("-------------------------------------------------")
		fmt.Print("Pilihan: ")
		fmt.Scanln(&opt)
		switch opt {
		case 1:
			fmt.Print("Nama Staff: ")
			fmt.Scanln(&in.name)
			fmt.Print("Telp Staff: ")
			fmt.Scanln(&in.num_telp)
			fmt.Print("Username Staff: ")
			fmt.Scanln(&in.username)
			fmt.Print("Password Staff: ")
			fmt.Scanln(&in.password)
			in.id = *n + 1
			C[in.id-1] = in
			*n++

			storeDataStaffs(*C, *n)
			loadDataStaffs(C, n)
		case 2:
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
			continue
		}
	}
}

func editStaff(C *arrStaffs, choosenStaff staff, n *int) {
	var opt int

	for {
		clearScreen()
		fmt.Println("-------------------------------------------------")
		fmt.Println("           EDIT STAFF")
		fmt.Println("-------------------------------------------------")
		fmt.Println()
		fmt.Println("ID  | Nama       | Telp       | Username   | Password")
		fmt.Println("-----------------------------------------------------")
		fmt.Printf("%-3d | %-10s | %-10s | %-10s | %-5s\n", choosenStaff.id, choosenStaff.name, choosenStaff.num_telp, choosenStaff.username, strings.Repeat("*", len(choosenStaff.password)))
		fmt.Println("-----------------------------------------------------")
		fmt.Println()
		fmt.Println("-------------------------------------------------")
		fmt.Println("1. Edit nama | 2. Edit telp | 3. Edit Username | 4. Edit Password | 5. Kembali")
		fmt.Println("-------------------------------------------------")
		fmt.Print("Pilihan: ")
		fmt.Scanln(&opt)
		switch opt {
		case 1:
			fmt.Print("Masukkan nama baru: ")
			fmt.Scanln(&choosenStaff.name)
		case 2:
			fmt.Print("Masukkan telp baru: ")
			fmt.Scanln(&choosenStaff.num_telp)
		case 3:
			fmt.Print("Masukkan username baru: ")
			fmt.Scanln(&choosenStaff.username)
		case 4:
			fmt.Print("Masukkan password baru: ")
			fmt.Scanln(&choosenStaff.password)
		case 5:
			return
		default:
			fmt.Println("**Pilihan tidak valid. Silakan coba lagi.**")
			continue
		}
		C[choosenStaff.id-1] = choosenStaff
		storeDataStaffs(*C, *n)
		loadDataStaffs(C, n)
	}
}

func deleteStaff(C *arrStaffs, choosen staff, n *int) {
	var opt string

	fmt.Println("Apakah anda yakin ingin menghapus pelanggan ini? (y/n)")
	fmt.Scanln(&opt)

	if opt == "y" {
		for i := choosen.id - 1; i < *n-1; i++ {
			C[i] = C[i+1]
		}

		*n--
		storeDataStaffs(*C, *n)
		loadDataStaffs(C, n)
	}
}
