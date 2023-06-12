package main

import "fmt"

func viewHistoryOrders(O arrOrders, n int) {
	if n == 0 {
		fmt.Println("Belum ada riwayat pemesanan.")
	} else {
		clearScreen()
		fmt.Println("-------------------------------------------------")
		fmt.Println("           LAPORAN PENJUALAN DEALER")
		fmt.Println("-------------------------------------------------")
		fmt.Println("Riwayat pemesanan:")
		fmt.Println("----------------------------------------------------------------------------------------------------------------------------------------------------------------------")
		fmt.Println("ID | Invoice            | Tanggal Pesan | Pajak | Total | Merek Mobil | Tahun Mobil | Nama Customer | KTP Customer | Telp Customer | Nama Staff | Telp Staff")
		fmt.Println("----------------------------------------------------------------------------------------------------------------------------------------------------------------------")
		for i := 0; i < n; i++ {
			fmt.Printf("%-3d| %-19s | %-13s | %-5d | %-8d | %-5s | %-8d | %-11s | %-10s | %-11s | %-14s | %-12s\n",
				O[i].id,
				O[i].invoice,
				O[i].date,
				O[i].tax,
				O[i].total,
				O[i].data_car.brand,
				O[i].data_car.year,
				O[i].data_customer.name,
				O[i].data_customer.num_ktp,
				O[i].data_customer.num_telp,
				O[i].data_staff.name,
				O[i].data_staff.num_telp)
		}
		fmt.Println("---------------------------------------------------------------------------------------------------------------------------------------------------------------------")
	}
	fmt.Println("Tekan ENTER untuk kembali...")
	fmt.Scanln()
}

func findHistoryOrder(O arrOrders, n int) {
	if n == 0 {
		fmt.Println("Belum ada riwayat pemesanan.")
	} else {
		var invoice string
		var found bool
		fmt.Print("Masukkan nomor invoice: ")
		fmt.Scanln(&invoice)

		for i := 0; i < n; i++ {
			if O[i].invoice == invoice {
				found = true
				clearScreen()
				invoiceOrder(invoice, O, n)
			}
		}
		if !found {
			fmt.Println("Nomor invoice tidak ditemukan.")
		}
		fmt.Println("Tekan ENTER untuk kembali...")
		fmt.Scanln()
	}
}

func generateTopThree(C arrCars, n int) {
	if n == 0 {
		fmt.Println("Belum ada riwayat pemesanan.")
	} else {
		var topList [MAXCARS]car
		var n_topList int
		for i := 0; i < n; i++ {
			if C[i].totalTerjual > 0 {
				topList[n_topList] = C[i]
				n_topList++
			}
		}

		fmt.Println("tess")

		// Sorting by totalTerjual
		for i := 0; i < n_topList; i++ {
			for j := i + 1; j < n_topList; j++ {
				if topList[i].totalTerjual < topList[j].totalTerjual {
					temp := topList[i]
					topList[i] = topList[j]
					topList[j] = temp
				}
			}
		}

		clearScreen()
		fmt.Println("-----------------------------------------------------------")
		fmt.Println("           LAPORAN PENJUALAN DEALER")
		fmt.Println("-----------------------------------------------------------")
		fmt.Printf("Top %d mobil terlaris:\n", MAXTOPS)
		fmt.Println("-----------------------------------------------------------")
		fmt.Println("No. | Merek      | Tipe       | Tahun | Qty | Total Omset")
		fmt.Println("-----------------------------------------------------------")
		for i := 0; i < MAXTOPS; i++ {
			fmt.Printf("%-3d | %-10s | %-10s | %-5d | %-3d | %-5d\n",
				i+1,
				topList[i].brand,
				topList[i].model,
				topList[i].year,
				topList[i].totalTerjual,
				topList[i].totalTerjual*topList[i].price)
		}
		fmt.Println("-----------------------------------------------------------")
	}
	fmt.Println("Tekan ENTER untuk kembali...")
	fmt.Scanln()
}
