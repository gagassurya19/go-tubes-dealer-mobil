package main

import (
	"fmt"
)

// CREATE NEW CAR
func addNewCarInterface() {
	clearScreen()

	fmt.Println("-------------------------------------------------")
	fmt.Println("             TAMBAH MOBIL BARU")
	fmt.Println("-------------------------------------------------")
	fmt.Println()
	fmt.Println("1. Tambah satu mobil")
	fmt.Println("2. Tambah banyak mobil (auto save)")
	fmt.Println("3. Kembali ke menu utama")
	fmt.Println()
}

func addNewCar(C *arrCars, n *int) {
	var in car
	var opt31 int

	clearScreen()

	for {
		fmt.Println("-------------------------------------------------")
		fmt.Println("         TAMBAH MOBIL BARU (SATU MOBIL)")
		fmt.Println("-------------------------------------------------")
		fmt.Println()
		in.id = C[*n-1].id + 1 // add id (auto increment)
		fmt.Print("Masukan merk: ")
		fmt.Scanln(&in.brand)
		fmt.Print("Masukan model: ")
		fmt.Scanln(&in.model)
		fmt.Print("Masukan year: ")
		fmt.Scanln(&in.year)
		fmt.Print("Masukan price: ")
		fmt.Scanln(&in.price)
		fmt.Println()
		fmt.Println("-------------------------------------------------")
		fmt.Println()
		fmt.Println("Data mobil yang akan ditambahkan:")
		fmt.Println("-------------------------------------------------")
		fmt.Println("ID    | Merek      | Model      | Tahun | Harga")
		fmt.Println("-------------------------------------------------")
		fmt.Printf("%-6d | %-10s | %-10s | %-5d | %d\n", in.id, in.brand, in.model, in.year, in.price)
		fmt.Println()
		fmt.Println("-------------------------------------------------")
		fmt.Println("Pilih: 1. Simpan | 2. Kembali")
		fmt.Print("Pilihan: ")
		fmt.Scan(&opt31)
		switch opt31 {
		case 1:
			C[*n] = in
			*n++
			storeDataCars(*C, *n)
			loadDataCars(C, n)

			clearScreen()

			fmt.Println("-------------------------------------------------")
			fmt.Println("MOBIL BERHASIL DITAMBAHKAN!")
			fmt.Println("-------------------------------------------------")
			fmt.Println()
			fmt.Println("ID    | Merek      | Model      | Tahun | Harga")
			fmt.Println("-------------------------------------------------")
			fmt.Printf("%-6d | %-10s | %-10s | %-5d | %d\n", in.id, in.brand, in.model, in.year, in.price)
			fmt.Println()
			fmt.Println("-------------------------------------------------")
			fmt.Println("Tekan Enter untuk melanjutkan...")
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

func addBulkCar(C *arrCars, n *int) {
	var in car
	var breaker bool = false

	clearScreen()

	fmt.Println("-------------------------------------------------")
	fmt.Println("         TAMBAH MOBIL BARU (BANYAK MOBIL)")
	fmt.Println("-------------------------------------------------")
	fmt.Println()
	fmt.Println("Masukan dengan format: [Merk] [Model] [Tahun] [Harga] // stop")
	fmt.Println()

	for !breaker {
		in.id = *n + 1 // add id (auto increment)
		fmt.Scanln(&in.brand, &in.model, &in.year, &in.price)

		if in.brand == "stop" {
			breaker = true
			fmt.Println()
			fmt.Println("-------------------------------------------------")
			fmt.Println("MOBIL BERHASIL DITAMBAHKAN!")
			fmt.Println("-------------------------------------------------")
			fmt.Println("Tekan Enter untuk melanjutkan...")
			fmt.Scanln()
		} else {
			C[*n] = in
			*n++
			storeDataCars(*C, *n)
			loadDataCars(C, n)
		}
	}
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
	fmt.Println("Pilih: 1. Edit | 2. Hapus | 3. Kembali ke menu utama")
}

func isFoundCar(C arrCars, n, id int) bool {
	for i := 0; i < n; i++ {
		if id == C[i].id {
			return true
		}
	}
	return false
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

		for (editCarData.id < 0 || editCarData.id > n) || !isFoundCar(*C, n, editCarData.id) {
			fmt.Println("ID mobil tidak valid/telah dihapus. Silakan coba lagi.")
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
		fmt.Println("Pilih: 1. Simpan | 2. Kembali")
		fmt.Print("Pilihan: ")
		fmt.Scan(&opt11)

		switch opt11 {
		case 1:
			C[idx] = editCarData
			storeDataCars(*C, n)
			loadDataCars(C, &n)

			clearScreen()
			fmt.Println("-------------------------------------------------")
			fmt.Println("MOBIL BERHASIL DI UBAH!")
			fmt.Println("-------------------------------------------------")
			fmt.Println("ID    | Merek      | Model      | Tahun | Harga")
			fmt.Println("-------------------------------------------------")
			fmt.Printf("%-6d | %-10s | %-10s | %-5d | %d\n", editCarData.id, editCarData.brand, editCarData.model, editCarData.year, editCarData.price)
			fmt.Println()
			fmt.Println("-------------------------------------------------")
			fmt.Println("Tekan Enter untuk melanjutkan...")
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
		fmt.Printf("ID yang tersedia: %d - %d | 0 (keluar)\n", C[0].id, C[*n-1].id)
		fmt.Print("ID Mobil: ")
		fmt.Scan(&idx)
		if idx == 0 {
			return
		}
		for idx < 0 || idx > C[*n-1].id {
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
		fmt.Println("Pilih: 1. Simpan | 2. Kembali")
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

			storeDataCars(*C, *n)
			loadDataCars(C, n)

			clearScreen()
			fmt.Println("-------------------------------------------------")
			fmt.Println("MOBIL BERHASIL DI HAPUS!")
			fmt.Println("-------------------------------------------------")
			fmt.Println("Tekan Enter untuk melanjutkan...")
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
}

func carSearchByBrand(C arrCars, n int) {
	var brand string
	clearScreen()
	for {
		fmt.Println("-------------------------------------------------")
		fmt.Println("             CARI MOBIL BERDASARKAN MEREK")
		fmt.Println("-------------------------------------------------")
		fmt.Println()
		fmt.Print("Masukkan merek mobil dicari: ")
		fmt.Scan(&brand)
		fmt.Println()
		fmt.Println("Mobil yang Anda cari:")
		fmt.Println("ID    | Merek      | Model      | Tahun | Harga")
		fmt.Println("-------------------------------------------------")
		for i := 0; i < n; i++ {
			if C[i].brand == brand {
				fmt.Printf("%-6d | %-10s | %-10s | %-5d | %d\n", C[i].id, C[i].brand, C[i].model, C[i].year, C[i].price)
			}
		}
		fmt.Println()
		fmt.Println("-------------------------------------------------")

		var opt13 int
		fmt.Println("Pilih: 1. sorting | 2. Kembali")
		fmt.Print("Pilihan: ")
		fmt.Scan(&opt13)
		switch opt13 {
		case 1:
			for {
				var jenis, order string
				fmt.Print("Masukan [id/model/tahun/harga] [a/A]: ")
				fmt.Scan(&jenis, &order)
				fmt.Println()
				fmt.Println("Mobil yang Anda cari:")
				fmt.Println("ID    | Merek      | Model      | Tahun | Harga")
				fmt.Println("-------------------------------------------------")

				switch jenis {
				case "id":
					if order == "a" {
						sortByIDAscending(&C, n)
					} else {
						sortByIDDescending(&C, n)
					}
				case "model":
					if order == "a" {
						sortByModelAscending(&C, n)
					} else {
						sortByModelDescending(&C, n)
					}
				case "tahun":
					if order == "a" {
						sortByYearAscending(&C, n)
					} else {
						sortByYearDescending(&C, n)
					}
				case "harga":
					if order == "a" {
						sortByPriceAscending(&C, n)
					} else {
						sortByPriceDescending(&C, n)
					}
				default:
					fmt.Println("Jenis sorting tidak valid.")
					continue
				}

				for i := 0; i < n; i++ {
					if C[i].brand == brand {
						fmt.Printf("%-6d | %-10s | %-10s | %-5d | %d\n", C[i].id, C[i].brand, C[i].model, C[i].year, C[i].price)
					}
				}

				fmt.Println()
				fmt.Println("-------------------------------------------------")
				fmt.Println("Pilih: 1. sorting | 2. Kembali")
				fmt.Print("Pilihan: ")
				fmt.Scan(&opt13)
				switch opt13 {
				case 1:
					continue
				case 2:
					return
				default:
					clearScreen()
					fmt.Println("-------------------------------------------------")
					fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
					continue
				}
			}
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

func carSearchByModel(C arrCars, n int) {
	var model string
	clearScreen()

	for {
		fmt.Println("-------------------------------------------------")
		fmt.Println("             CARI MOBIL BERDASARKAN MODEL")
		fmt.Println("-------------------------------------------------")
		fmt.Println()
		fmt.Print("Silakan masukkan model mobil yang Anda cari: ")
		fmt.Scan(&model)
		fmt.Println()
		fmt.Println("Mobil yang Anda cari:")
		fmt.Println("ID    | Merek      | Model      | Tahun | Harga")
		fmt.Println("-------------------------------------------------")
		for i := 0; i < n; i++ {
			if C[i].model == model {
				fmt.Printf("%-6d | %-10s | %-10s | %-5d | %d\n", C[i].id, C[i].brand, C[i].model, C[i].year, C[i].price)
			}
		}
		fmt.Println()
		fmt.Println("-------------------------------------------------")

		var opt13 int
		fmt.Println("Pilih: 1. sorting | 2. Kembali")
		fmt.Print("Pilihan: ")
		fmt.Scan(&opt13)
		switch opt13 {
		case 1:
			for {
				var jenis, order string
				fmt.Print("Masukan [id/merek/tahun/harga] [a/A]: ")
				fmt.Scan(&jenis, &order)
				fmt.Println()
				fmt.Println("Mobil yang Anda cari:")
				fmt.Println("ID    | Merek      | Model      | Tahun | Harga")
				fmt.Println("-------------------------------------------------")

				switch jenis {
				case "id":
					if order == "a" {
						sortByIDAscending(&C, n)
					} else {
						sortByIDDescending(&C, n)
					}
				case "merek":
					if order == "a" {
						sortByBrandAscending(&C, n)
					} else {
						sortByBrandDescending(&C, n)
					}
				case "tahun":
					if order == "a" {
						sortByYearAscending(&C, n)
					} else {
						sortByYearDescending(&C, n)
					}
				case "harga":
					if order == "a" {
						sortByPriceAscending(&C, n)
					} else {
						sortByPriceDescending(&C, n)
					}
				default:
					fmt.Println("Jenis sorting tidak valid.")
					continue
				}

				for i := 0; i < n; i++ {
					if C[i].model == model {
						fmt.Printf("%-6d | %-10s | %-10s | %-5d | %d\n", C[i].id, C[i].brand, C[i].model, C[i].year, C[i].price)
					}
				}

				fmt.Println()
				fmt.Println("-------------------------------------------------")
				fmt.Println("Pilih: 1. sorting | 2. Kembali")
				fmt.Print("Pilihan: ")
				fmt.Scan(&opt13)
				switch opt13 {
				case 1:
					continue
				case 2:
					return
				default:
					clearScreen()
					fmt.Println("-------------------------------------------------")
					fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
					continue
				}
			}
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

func carSearchByYear(C arrCars, n int) {
	var year int
	clearScreen()

	for {
		fmt.Println("-------------------------------------------------")
		fmt.Println("             CARI MOBIL BERDASARKAN TAHUN")
		fmt.Println("-------------------------------------------------")
		fmt.Println()
		fmt.Print("Silakan masukkan tahun mobil yang Anda cari: ")
		fmt.Scan(&year)
		fmt.Println()
		fmt.Println("Mobil yang Anda cari:")
		fmt.Println("ID    | Merek      | Model      | Tahun | Harga")
		fmt.Println("-------------------------------------------------")
		for i := 0; i < n; i++ {
			if C[i].year == year {
				fmt.Printf("%-6d | %-10s | %-10s | %-5d | %d\n", C[i].id, C[i].brand, C[i].model, C[i].year, C[i].price)
			}
		}
		fmt.Println()
		fmt.Println("-------------------------------------------------")

		var opt13 int
		fmt.Println("Pilih: 1. sorting | 2. Kembali")
		fmt.Print("Pilihan: ")
		fmt.Scan(&opt13)
		switch opt13 {
		case 1:
			for {
				var jenis, order string
				fmt.Print("Masukan [id/merek/model/harga] [a/A]: ")
				fmt.Scan(&jenis, &order)
				fmt.Println()
				fmt.Println("Mobil yang Anda cari:")
				fmt.Println("ID    | Merek      | Model      | Tahun | Harga")
				fmt.Println("-------------------------------------------------")

				switch jenis {
				case "id":
					if order == "a" {
						sortByIDAscending(&C, n)
					} else {
						sortByIDDescending(&C, n)
					}
				case "merek":
					if order == "a" {
						sortByBrandAscending(&C, n)
					} else {
						sortByBrandDescending(&C, n)
					}
				case "model":
					if order == "a" {
						sortByModelAscending(&C, n)
					} else {
						sortByModelDescending(&C, n)
					}
				case "harga":
					if order == "a" {
						sortByPriceAscending(&C, n)
					} else {
						sortByPriceDescending(&C, n)
					}
				default:
					fmt.Println("Jenis sorting tidak valid.")
					continue
				}

				for i := 0; i < n; i++ {
					if C[i].year == year {
						fmt.Printf("%-6d | %-10s | %-10s | %-5d | %d\n", C[i].id, C[i].brand, C[i].model, C[i].year, C[i].price)
					}
				}

				fmt.Println()
				fmt.Println("-------------------------------------------------")
				fmt.Println("Pilih: 1. sorting | 2. Kembali")
				fmt.Print("Pilihan: ")
				fmt.Scan(&opt13)
				switch opt13 {
				case 1:
					continue
				case 2:
					return
				default:
					clearScreen()
					fmt.Println("-------------------------------------------------")
					fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
					continue
				}
			}
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
