package main

import (
	"fmt"
	"os"
	"os/exec"
)

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func home() {
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
	fmt.Println("Halo, Petugas Dealer!\nSilakan pilih salah satu opsi di bawah ini:")
	fmt.Println("WIP : Work In Progress (DELETE AFTER DONE)")
	fmt.Println()
	fmt.Println("1. Lihat Daftar Mobil")
	fmt.Println("2. Cari Mobil (WIP)")
	fmt.Println("3. Tambah Mobil Baru")
	fmt.Println("4. Pesan Mobil untuk Pelanggan (WIP)")
	fmt.Println("5. Laporan Penjualan (WIP)")
	fmt.Println("6. Keluar ")
	fmt.Println()
	fmt.Print("Silakan masukkan nomor pilihan Anda: ")
}

func orderCar() {
	clearScreen()

	fmt.Println("-------------------------------------------------")
	fmt.Println("           PESAN MOBIL UNTUK PELANGGAN")
	fmt.Println("-------------------------------------------------")
	fmt.Println()
	fmt.Println("Silakan masukkan detail pemesanan:")
	fmt.Println()
	fmt.Print("ID Mobil: ")
	fmt.Print("ID Mobil: ")
	fmt.Print("Nama Pelanggan: ")
	fmt.Print("Alamat Pengiriman: ")

	fmt.Println()
	fmt.Println("Pesanan berhasil dibuat:")
	fmt.Println()
	fmt.Printf("ID Mobil: 1")
	fmt.Printf("Merek: Toyota")
	fmt.Printf("Model: Avanza")
	fmt.Printf("Nama Pelanggan: Budi")
	fmt.Printf("Alamat Pengiriman: Jl. Merdeka No. 1")
}

func generateSalesReport() {
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
