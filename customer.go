package main

import (
	"fmt"
	"strconv"
)

func listCustomer(C arrCustomers, n int) {
	var opt int

	for {
		clearScreen()
		fmt.Println("-------------------------------------------------")
		fmt.Println("           LIHAT DAFTAR PELANGGAN")
		fmt.Println("-------------------------------------------------")
		fmt.Println()

		if n == 0 {
			fmt.Println("Tidak ada pelanggan yang terdaftar")
		} else {
			fmt.Println("ID | Nama                | KTP        | Telp         | Alamat")
			fmt.Println("--------------------------------------------------------------------")

			for i := 0; i < n; i++ {
				fmt.Printf("%-3d| %-19s | %-10s | %-10s | %-10s\n", C[i].id, C[i].name, C[i].num_ktp, C[i].num_telp, C[i].address)
			}
		}

		fmt.Println("--------------------------------------------------------------------")
		fmt.Println("1. Sort 2. Kembali")
		fmt.Println("-------------------------------------------------")
		fmt.Print("Pilihan: ")
		fmt.Scanln(&opt)

		switch opt {
		case 1:
			sortCustomer(&C, n)
		case 2:
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
			continue
		}
	}
}

func sortCustomer(C *arrCustomers, n int) {
	var jenis, sort string
	fmt.Print("Masukkan [id/nama/ktp/telp/alamat] [a/A]:")
	fmt.Scanln(&jenis, &sort)
	switch jenis {
	case "id":
		if sort == "A" {
			sortCustomerByIDDesc(C, n)
		} else {
			sortCustomerByIDAsc(C, n)
		}
	case "nama":
		if sort == "A" {
			sortCustomerByNameDesc(C, n)
		} else {
			sortCustomerByNameAsc(C, n)
		}
	case "ktp":
		if sort == "A" {
			sortCustomerByKTPDesc(C, n)
		} else {
			sortCustomerByKTPAsc(C, n)
		}
	case "telp":
		if sort == "A" {
			sortCustomerByTelpDesc(C, n)
		} else {
			sortCustomerByTelpAsc(C, n)
		}
	case "alamat":
		if sort == "A" {
			sortCustomerByAddressDesc(C, n)
		} else {
			sortCustomerByAddressAsc(C, n)
		}
	default:
		fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
	}
}

func searchCustomer(C *arrCustomers, n *int) {
	var opt int
	var keyword, message string
	var choosenCustomer customer

	for {
		clearScreen()
		fmt.Println("-------------------------------------------------")
		fmt.Println("           CARI PELANGGAN")
		fmt.Println("-------------------------------------------------")
		fmt.Println()

		if *n == 0 {
			fmt.Println("Tidak ada pelanggan yang terdaftar")
		} else {
			fmt.Println("ID | Nama                | KTP        | Telp         | Alamat")
			fmt.Println("--------------------------------------------------------------------")

			for i := 0; i < *n; i++ {
				fmt.Printf("%-3d| %-19s | %-10s | %-10s | %-10s\n", C[i].id, C[i].name, C[i].num_ktp, C[i].num_telp, C[i].address)
			}
		}

		fmt.Println("--------------------------------------------------------------------")
		fmt.Println(message)
		fmt.Print("Masukkan id/nama/ktp/telp/alamat | ` (exit): ")
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
			if C[i].id == numkeyword || C[i].name == keyword || C[i].num_ktp == keyword || C[i].num_telp == keyword || C[i].address == keyword {
				clearScreen()
				fmt.Println("-------------------------------------------------")
				fmt.Println("           CARI PELANGGAN")
				fmt.Println("-------------------------------------------------")
				fmt.Println()
				fmt.Println("ID | Nama                | KTP        | Telp         | Alamat")
				fmt.Println("--------------------------------------------------------------------")
				fmt.Printf("%-3d| %-19s | %-10s | %-10s | %-10s\n", C[i].id, C[i].name, C[i].num_ktp, C[i].num_telp, C[i].address)
				fmt.Println("--------------------------------------------------------------------")
				choosenCustomer = C[i]
				found = true
				break
			}
		}

		message = ""
		keyword = ""

		if !found {
			clearScreen()
			fmt.Println("-------------------------------------------------")
			fmt.Println("           CARI PELANGGAN")
			fmt.Println("-------------------------------------------------")
			fmt.Println()
			fmt.Println("Pelanggan tidak ditemukan")
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
				editCustomer(C, choosenCustomer, n)
			case 3:
				deleteCustomer(C, choosenCustomer, n)
			case 4:
				return
			default:
				message = "**Pilihan tidak valid. Silakan coba lagi.**"
				continue
			}
		}
	}
}

func addCustomer(C *arrCustomers, n *int) {
	var in customer
	var opt int

	for {
		var message string
		for i := 0; i < 2; i++ {
			clearScreen()
			fmt.Println("-------------------------------------------------")
			fmt.Println("           TAMBAH PELANGGAN")
			fmt.Println("-------------------------------------------------")
			fmt.Println()
			fmt.Println("LIST PELANGGAN:")
			fmt.Println("ID | Nama                | KTP        | Telp         | Alamat")
			fmt.Println("--------------------------------------------------------------------")
			for i := 0; i < *n; i++ {
				fmt.Printf("%-3d| %-19s | %-10s | %-10s | %-10s\n", C[i].id, C[i].name, C[i].num_ktp, C[i].num_telp, C[i].address)
			}
			fmt.Println("--------------------------------------------------------------------")
			fmt.Println(message)

			if i != 1 {
				fmt.Println("TAMBAH PELANGGAN:")
				fmt.Print("Nama: ")
				fmt.Scanln(&in.name)
				fmt.Print("No. KTP: ")
				fmt.Scanln(&in.num_ktp)
				fmt.Print("No. Telp: ")
				fmt.Scanln(&in.num_telp)
				fmt.Print("Alamat: ")
				fmt.Scanln(&in.address)

				in.id = *n + 1

				C[*n] = in
				*n++

				message = "**Data Berhasil Ditambahkan**"
			}
		}

		fmt.Println()
		fmt.Println("-------------------------------------------------")
		fmt.Println("1. Tambah lagi | 2. Kembali")
		fmt.Println("-------------------------------------------------")
		fmt.Print("Pilihan: ")
		fmt.Scanln(&opt)
		switch opt {
		case 1:
			continue
		case 2:
			return
		default:
			fmt.Println("**Pilihan tidak valid. Silakan coba lagi.**")
		}
	}
}

func editCustomer(C *arrCustomers, choosen customer, n *int) {
	var opt int
	var in customer

	for {
		clearScreen()
		fmt.Println("-------------------------------------------------")
		fmt.Println("           EDIT PELANGGAN")
		fmt.Println("-------------------------------------------------")
		fmt.Println()
		fmt.Println("ID | Nama                | KTP        | Telp         | Alamat")
		fmt.Println("--------------------------------------------------------------------")
		fmt.Printf("%-3d| %-19s | %-10s | %-10s | %-10s\n", choosen.id, choosen.name, choosen.num_ktp, choosen.num_telp, choosen.address)
		fmt.Println("--------------------------------------------------------------------")

		fmt.Println()
		fmt.Println("1. Edit nama | 2. Edit KTP | 3. Edit telp | 4. Edit alamat | 5. Kembali")
		fmt.Println("-------------------------------------------------")
		fmt.Print("Pilihan: ")
		fmt.Scanln(&opt)

		switch opt {
		case 1:
			fmt.Print("Input name:")
			fmt.Scanln(&in.name)
			choosen.name = in.name
		case 2:
			fmt.Print("Input KTP number:")
			fmt.Scanln(&in.num_ktp)
			choosen.num_ktp = in.num_ktp
		case 3:
			fmt.Print("Input Telphone number:")
			fmt.Scanln(&in.num_telp)
			choosen.num_telp = in.num_telp
		case 4:
			fmt.Print("Input address:")
			fmt.Scanln(&in.address)
			choosen.address = in.address
		case 5:
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
			continue
		}

		C[choosen.id-1] = choosen
	}
}

func deleteCustomer(C *arrCustomers, choosen customer, n *int) {
	var opt string

	fmt.Println("Apakah anda yakin ingin menghapus pelanggan ini? (y/n)")
	fmt.Scanln(&opt)

	if opt == "y" {
		for i := choosen.id - 1; i < *n-1; i++ {
			C[i] = C[i+1]
		}

		*n--
	}
}
