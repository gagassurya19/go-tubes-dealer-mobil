package main

import (
	"fmt"
)

func home(namaPetugas string) {
	clearScreen()

	fmt.Println("-------------------------------------------------")
	fmt.Println("        TUGAS BESAR ALGORITMA DAN PEMROGRAMAN")
	fmt.Println("                      IF-46-10")
	fmt.Println("-------------------------------------------------")
	fmt.Println("           SELAMAT DATANG DI DEALER MOBIL")
	fmt.Println("-------------------------------------------------")
	fmt.Println("Kelompok:")
	fmt.Println("1. Gagas Surya Laksana (1301223164)")
	fmt.Println("2. Yohanes Maurich Pangkey (1301220203)")
	fmt.Println("-------------------------------------------------")
	fmt.Printf("Halo, %s! - petugas dealer\nSilakan pilih salah satu opsi di bawah ini:\n", namaPetugas)
	fmt.Println("WIP : Work In Progress (DELETE AFTER DONE)")
	fmt.Println()
	fmt.Println("1. Lihat Daftar Mobil")
	fmt.Println("2. Cari Mobil")
	fmt.Println("3. Tambah Mobil Baru")
	fmt.Println("4. PESAN Mobil untuk Pelanggan (WIP)")
	fmt.Println("5. CRUD Customer")
	fmt.Println("6. CRUD Staff")
	fmt.Println("7. Laporan Penjualan (WIP)")
	fmt.Println("8. Keluar ")
	fmt.Println()
	fmt.Print("Pilihan: ")
}

func viewCustomer() {
	clearScreen()

	fmt.Println("-------------------------------------------------")
	fmt.Println("           LIHAT DAFTAR PELANGGAN")
	fmt.Println("-------------------------------------------------")
	fmt.Println()
	fmt.Println("1. Lihat semua pelanggan")
	fmt.Println("2. Cari pelanggan")
	fmt.Println("3. Tambah pelanggan baru")
	fmt.Println("4. Kembali ke menu utama")
	fmt.Println()
	fmt.Print("Pilihan: ")
}

func viewStaff() {
	clearScreen()

	fmt.Println("-------------------------------------------------")
	fmt.Println("           LIHAT DAFTAR STAFF")
	fmt.Println("-------------------------------------------------")
	fmt.Println()
	fmt.Println("1. Lihat semua staff")
	fmt.Println("2. Cari staff")
	fmt.Println("3. Tambah staff baru")
	fmt.Println("4. Kembali ke menu utama")
	fmt.Println()
	fmt.Print("Pilihan: ")
}

func viewGenerateSalesReport() {
	clearScreen()

	fmt.Println("-------------------------------------------------")
	fmt.Println("           LAPORAN PENJUALAN DEALER")
	fmt.Println("-------------------------------------------------")
	fmt.Println()
	fmt.Println("Silakan masukkan periode laporan penjualan:")
	fmt.Println()
	fmt.Print("Bulan: ")
	fmt.Print("Tahun: ")
	fmt.Println()
	fmt.Println("Laporan penjualan:")
	fmt.Println()
	fmt.Printf("Periode: Bulan juni, Tahun 2021")
	fmt.Printf("Total Penjualan: $ 1000000")
}

func exitApp() {
	clearScreen()

	fmt.Println("-------------------------------------------------")
	fmt.Println("                 KELUAR DARI APLIKASI")
	fmt.Println("-------------------------------------------------")
	fmt.Println()
	fmt.Println("Terima kasih telah menggunakan Dealer Mobil App. Sampai jumpa lagi!")
}
