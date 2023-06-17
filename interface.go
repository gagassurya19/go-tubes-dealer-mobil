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
	fmt.Println("        🚘 SELAMAT DATANG DI DEALER MOBIL 🚘")
	fmt.Println("-------------------------------------------------")
	fmt.Println("Kelompok:")
	fmt.Println("1️⃣  Gagas Surya Laksana (1301223164)")
	fmt.Println("2️⃣  Yohanes Maurich Pangkey (1301220203)")
	fmt.Println("-------------------------------------------------")
	fmt.Printf("Halo, %s! - petugas dealer\nSilakan pilih salah satu opsi di bawah ini:\n", namaPetugas)
	fmt.Println()
	fmt.Println("1. LIHAT DAFTAR MOBIL")
	fmt.Println("2. CARI MOBIL")
	fmt.Println("3. TAMBAH MOBIL BARU")
	fmt.Println("4. ORDER MOBIL")
	fmt.Println("5. CRUD CUSTOMER")
	fmt.Println("6. CRUD STAFF")
	fmt.Println("7. LAPORAN PENJUALAN")
	fmt.Println("8. KELUAR")
	fmt.Println("                                              🐈")
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
	fmt.Println("1. Top 3 mobil terlaris")
	fmt.Println("2. Lihat history penjualan")
	fmt.Println("3. Cari history penjualan")
	fmt.Println("4. Kembali ke menu utama")
	fmt.Println()
	fmt.Print("Pilihan: ")
}

func exitApp() {
	clearScreen()

	fmt.Println("-------------------------------------------------")
	fmt.Println("                 KELUAR DARI APLIKASI")
	fmt.Println("-------------------------------------------------")
	fmt.Println()
	fmt.Println("Terima kasih telah menggunakan Dealer Mobil App. Sampai jumpa lagi!")
}
