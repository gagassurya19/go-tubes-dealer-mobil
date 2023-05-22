package main

import (
	"fmt"
)

func addCar(C *arrCars, n *int) {
	var in car

	fmt.Print("Input brand:")
	fmt.Scanln(&in.brand)
	fmt.Print("Input model:")
	fmt.Scanln(&in.model)
	fmt.Print("Input year:")
	fmt.Scanln(&in.year)
	fmt.Print("Input price:")
	fmt.Scanln(&in.price)

	in.id = *n
	C[*n] = in

	*n++
}

func addBulkCar(C *arrCars, n *int) {
	var in car
	var isBreak bool = false

	fmt.Println("Input with format: [Brand] [Model] [Year] [Price] // stop")
	for !isBreak {
		fmt.Scanln(&in.brand, &in.model, &in.year, &in.price)

		if in.brand == "stop" {
			isBreak = true
		} else {
			in.id = *n
			C[*n] = in
			*n++
		}
	}
}

// CREATE NEW CAR
func addNewCarInterface() {
	clearScreen()

	fmt.Println("-------------------------------------------------")
	fmt.Println("             TAMBAH MOBIL BARU")
	fmt.Println("-------------------------------------------------")
	fmt.Println()
	fmt.Println("Silakan masukkan detail mobil yang ingin ditambahkan:")
	fmt.Println()
	fmt.Print("Merek: ")
	fmt.Print("Model: ")
	fmt.Print("Tahun: ")
	fmt.Print("Harga: ")
	fmt.Println()
	fmt.Println("Mobil berhasil ditambahkan!")
}

// GET LIST OF CARS
func carListInterface(C arrCars, n int) {
	clearScreen()

	fmt.Println("-------------------------------------------------")
	fmt.Println("             LIHAT DAFTAR MOBIL DI DEALER")
	fmt.Println("-------------------------------------------------")
	fmt.Println()
	fmt.Println("Berikut adalah daftar mobil yang tersedia di dealer:")
	fmt.Println()
	fmt.Println("ID    | Merek      | Model      | Tahun | Harga")
	fmt.Println("-------------------------------------------------")
	for i := 0; i < n; i++ {
		fmt.Printf("%-6d | %-10s | %-10s | %-5d | %d\n", C[i].id, C[i].brand, C[i].model, C[i].year, C[i].price)
	}
	fmt.Println()
	fmt.Println("-------------------------------------------------")
	fmt.Println("Pilih: 1. Edit mobil | 2. Hapus mobil | 3. Kembali ke menu utama")
}

// EDIT CAR INTERFACE
func editCarInterface(C *arrCars, n int) {
	var opt11, idx int
	var editCarData car

	clearScreen()

	for {
		fmt.Println("-------------------------------------------------")
		fmt.Println("             EDIT MOBIL")
		fmt.Println("-------------------------------------------------")
		fmt.Println()
		fmt.Println("Silakan masukkan detail mobil yang ingin diedit:")
		fmt.Println()
		fmt.Printf("ID yang tersedia: %d - %d | 0 (keluar)\n", C[0].id, n)
		fmt.Print("ID Mobil: ")
		fmt.Scanln(&editCarData.id)
		if editCarData.id == 0 {
			return
		}
		for editCarData.id < 0 || editCarData.id > n {
			fmt.Println("ID mobil tidak valid. Silakan coba lagi.")
			fmt.Print("ID Mobil: ")
			fmt.Scanln(&editCarData.id)
		}

		fmt.Println()
		fmt.Println("Mobil yang ingin diedit:")
		fmt.Println("ID    | Merek      | Model      | Tahun | Harga")
		fmt.Println("-------------------------------------------------")
		for i := 0; i < n; i++ {
			if C[i].id == editCarData.id {
				idx = i
				fmt.Printf("%-6d | %-10s | %-10s | %-5d | %d\n", C[i].id, C[i].brand, C[i].model, C[i].year, C[i].price)
			}
		}
		fmt.Println()

		fmt.Println("Silakan masukkan detail mobil yang ingin diedit:")
		fmt.Print("Merek: ")
		fmt.Scanln(&editCarData.brand)
		fmt.Print("Model: ")
		fmt.Scanln(&editCarData.model)
		fmt.Print("Tahun: ")
		fmt.Scanln(&editCarData.year)
		fmt.Print("Harga: ")
		fmt.Scanln(&editCarData.price)
		fmt.Println()
		fmt.Println("Data mobil yang baru:")
		fmt.Println("ID    | Merek      | Model      | Tahun | Harga")
		fmt.Println("-------------------------------------------------")
		fmt.Printf("%-6d | %-10s | %-10s | %-5d | %d\n", editCarData.id, editCarData.brand, editCarData.model, editCarData.year, editCarData.price)
		fmt.Println()
		fmt.Println("-------------------------------------------------")
		fmt.Println("Pilih: 1. Save | 2. Kembali")
		fmt.Print("Pilihan: ")
		fmt.Scan(&opt11)

		switch opt11 {
		case 1:
			C[idx] = editCarData
			clearScreen()
			fmt.Println("-------------------------------------------------")
			fmt.Println("Car successfully edited!")
			fmt.Println("-------------------------------------------------")
			fmt.Println("ID    | Merek      | Model      | Tahun | Harga")
			fmt.Println("-------------------------------------------------")
			fmt.Printf("%-6d | %-10s | %-10s | %-5d | %d\n", editCarData.id, editCarData.brand, editCarData.model, editCarData.year, editCarData.price)
			fmt.Println()
			fmt.Println("-------------------------------------------------")
			fmt.Println("Press Enter to continue...")
			fmt.Scanln()
			return
		case 2:
			return
		default:
			clearScreen()
			fmt.Println("-------------------------------------------------")
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
			continue
		}
	}
}

// DELETE CAR INTERFACE
func deleteCarInterface(C *arrCars, n *int) {
	var idx, opt12 int

	clearScreen()

	for {
		fmt.Println("-------------------------------------------------")
		fmt.Println("             HAPUS MOBIL")
		fmt.Println("-------------------------------------------------")
		fmt.Println()
		fmt.Println("Silakan masukkan ID mobil yang ingin dihapus:")
		fmt.Println()
		fmt.Printf("ID yang tersedia: %d - %d | 0 (keluar)\n", C[0].id, *n)
		fmt.Print("ID Mobil: ")
		fmt.Scan(&idx)
		if idx == 0 {
			return
		}
		for idx < 0 || idx > *n {
			fmt.Println("ID mobil tidak valid. Silakan coba lagi.")
			fmt.Print("ID Mobil: ")
			fmt.Scan(&idx)
		}

		fmt.Println()
		fmt.Println("Mobil yang ingin dihapus:")
		fmt.Println("ID    | Merek      | Model      | Tahun | Harga")
		fmt.Println("-------------------------------------------------")
		for i := 0; i < *n; i++ {
			if C[i].id == idx {
				fmt.Printf("%-6d | %-10s | %-10s | %-5d | %d\n", C[i].id, C[i].brand, C[i].model, C[i].year, C[i].price)
			}
		}

		fmt.Println()
		fmt.Println("-------------------------------------------------")
		fmt.Println("Pilih: 1. Save | 2. Kembali")
		fmt.Print("Pilihan: ")
		fmt.Scan(&opt12)

		switch opt12 {
		case 1:
			// shift to left
			for i := idx - 1; i < *n-1; i++ {
				C[i] = C[i+1]
			}
			// Clear the last element
			C[*n-1] = car{}

			*n--

			clearScreen()
			fmt.Println("-------------------------------------------------")
			fmt.Println("Car successfully deleted!")
			fmt.Println("-------------------------------------------------")
			fmt.Println("Press Enter to continue...")
			fmt.Scanln()
			return
		case 2:
			return
		default:
			clearScreen()
			fmt.Println("-------------------------------------------------")
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
			continue
		}
	}
}

// SEARCH CAR INTERFACE
func carSearchInterface() {
	clearScreen()

	fmt.Println("-------------------------------------------------")
	fmt.Println("             CARI MOBIL DI DEALER")
	fmt.Println("-------------------------------------------------")
	fmt.Println()
	fmt.Println("Silakan masukkan kriteria pencarian:")
	fmt.Println()
	fmt.Println("1. Cari berdasarkan merek")
	fmt.Println("2. Cari berdasarkan model")
	fmt.Println("3. Cari berdasarkan tahun")
	fmt.Println("4. Kembali ke menu utama")
	fmt.Println()
	fmt.Print("Silakan masukkan nomor kriteria pencarian Anda: ")
}

func carSearchByBrandInterface() {
	clearScreen()

	fmt.Println("-------------------------------------------------")
	fmt.Println("             CARI MOBIL BERDASARKAN MEREK")
	fmt.Println("-------------------------------------------------")
	fmt.Println()
	fmt.Print("Silakan masukkan merek mobil yang Anda cari: ")
}

func carSearchByModelInterface() {
	clearScreen()

	fmt.Println("-------------------------------------------------")
	fmt.Println("             CARI MOBIL BERDASARKAN MODEL")
	fmt.Println("-------------------------------------------------")
	fmt.Println()
	fmt.Print("Silakan masukkan model mobil yang Anda cari: ")
}

func carSearchByYearInterface() {
	fmt.Println("-------------------------------------------------")
	fmt.Println("             CARI MOBIL BERDASARKAN TAHUN")
	fmt.Println("-------------------------------------------------")
	fmt.Println()
	fmt.Print("Silakan masukkan tahun mobil yang Anda cari: ")
}
