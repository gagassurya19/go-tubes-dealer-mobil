package main

import (
	"fmt"
	"strconv"
	"time"
)

func placeOrder(O *arrOrders, n *int, cars *arrCars, n_car int, customers arrCustomers, n_customer int, staff staff, staffs arrStaffs) {
	var foundCar, foundCustomer, found bool = false, false, false
	var id_car_input, id_customer_input string
	var message string

	var carChoosen car
	var customerChoosen customer

	for !found {
		clearScreen()

		fmt.Println("-------------------------------------------------")
		fmt.Println("           PESAN MOBIL UNTUK PELANGGAN")
		fmt.Println("-------------------------------------------------")
		fmt.Println()
		fmt.Println("Silakan masukkan detail pemesanan:")
		fmt.Println(message)
		fmt.Println()

		if !foundCar {
			fmt.Printf("ID Mobil (%d - %d) | ` (exit): ", cars[0].id, cars[n_car-1].id)
			fmt.Scanln(&id_car_input)

			if id_car_input == "`" {
				return
			}

			id_car, err := strconv.Atoi(id_car_input)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			if id_car < 0 || id_car > cars[n_car-1].id {
				message = "ID Mobil tidak valid. Silakan coba lagi."
				continue
			}

			carFound, carData := searchCar(*cars, n_car, id_car)
			if carFound {
				fmt.Println("ID  | Merek    | Model    | Tahun    | Harga")
				fmt.Println("-------------------------------------------------")
				fmt.Printf("%-3d | %-8s | %-8s | %-8d | %-8d\n", carData.id, carData.brand, carData.model, carData.year, carData.price)
				fmt.Println("-------------------------------------------------")

				var opt string
				fmt.Print("Apakah anda yakin ingin memesan mobil ini? (y/n) ")
				fmt.Scanln(&opt)
				if opt == "y" {
					fmt.Println("**Berhasil masuk keranjang!")
					fmt.Println()
					foundCar = true
					carChoosen = carData
				} else {
					fmt.Println("Pemesanan dibatalkan.")
				}
			}
		} else {
			fmt.Println("Mobil yang dipilih:")
			fmt.Println("ID  | Merek    | Model    | Tahun    | Harga")
			fmt.Println("-------------------------------------------------")
			fmt.Printf("%-3d | %-8s | %-8s | %-8d | %-8d\n", carChoosen.id, carChoosen.brand, carChoosen.model, carChoosen.year, carChoosen.price)
			fmt.Println("-------------------------------------------------")
			fmt.Println()
		}

		if foundCar && !foundCustomer {
			fmt.Printf("ID Customer (1 - %d) | ` (exit): ", n_customer-1)
			fmt.Scanln(&id_customer_input)

			if id_customer_input == "`" {
				return
			}

			id_customer, err := strconv.Atoi(id_customer_input)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			if id_customer < 0 || id_customer >= n_customer {
				message = "ID Customer tidak valid. Silakan coba lagi."
				continue
			}

			customerFound, customerData := searchCustomerOrder(customers, n_customer, id_customer)
			if customerFound {
				fmt.Println("ID | Nama                | KTP        | Telp         | Alamat")
				fmt.Println("--------------------------------------------------------------------")
				fmt.Printf("%-3d| %-19s | %-10s | %-10s | %-10s\n", customerData.id, customerData.name, customerData.num_ktp, customerData.num_telp, customerData.address)
				fmt.Println("--------------------------------------------------------------------")

				var opt string
				fmt.Print("Apakah anda yakin benar? (y/n) ")
				fmt.Scanln(&opt)
				if opt == "y" {
					fmt.Println("**Berhasil!")
					foundCustomer = true
					customerChoosen = customerData
				} else {
					fmt.Println("Pemesanan dibatalkan.")
				}
			}
		} else {
			fmt.Println("Customer yang dipilih:")
			fmt.Println("ID  | Nama")
			fmt.Println("-------------------------------------------------")
			fmt.Println("ID | Nama                | KTP        | Telp         | Alamat")
			fmt.Println("--------------------------------------------------------------------")
			fmt.Printf("%-3d| %-19s | %-10s | %-10s | %-10s\n", customerChoosen.id, customerChoosen.name, customerChoosen.num_ktp, customerChoosen.num_telp, customerChoosen.address)
			fmt.Println("--------------------------------------------------------------------")
		}

		if foundCar && foundCustomer {
			var opt string
			fmt.Println()
			fmt.Print("Apakah anda yakin ingin memesan mobil pada customer ini? (y/n) ")
			fmt.Scanln(&opt)
			if opt == "y" {
				found = true
			} else {
				found, foundCar, foundCustomer = false, false, false
			}
		}
	}

	postToDatabase(O, n, cars, n_car, carChoosen, customerChoosen, staff)

	storeDataOrders(*O, *n, customers, *cars, staffs)
	loadDataOrders(O, n, customers, *cars, staffs)

	fmt.Println(cars)

	storeDataCars(*cars, n_car)
	loadDataCars(cars, &n_car)

	fmt.Println(cars)

	clearScreen()
	invoiceNumber := O[*n-1].invoice
	invoiceOrder(invoiceNumber, *O, *n)

	fmt.Println("Tekan ENTER untuk kembali ke menu utama...")
	fmt.Scanln()
}

func postToDatabase(O *arrOrders, n *int, carList *arrCars, n_car int, carChoosen car, customerChoosen customer, staff staff) {
	// Generate invoice number
	invoiceNumber := generateInvoiceNumber()
	// Get current date
	date := time.Now().Format("2006-01-02")

	// Update the car data total terjual
	for i := 0; i < n_car; i++ {
		if carList[i].id == carChoosen.id {
			carList[i].totalTerjual++
		}
	}

	// Create the order and add it to the array
	O[*n] = order{
		id:            *n + 1,
		invoice:       invoiceNumber,
		date:          date,
		tax:           carChoosen.price * 10 / 100,
		subtotal:      carChoosen.price,
		total:         carChoosen.price + carChoosen.price*10/100,
		data_car:      carChoosen,
		data_customer: customerChoosen,
		data_staff:    staff,
	}
	*n++
}

func generateInvoiceNumber() string {
	// Generate a unique invoice number based on the current timestamp
	timestamp := time.Now().UnixNano()
	invoiceNumber := fmt.Sprintf("INV-%d", timestamp)

	return invoiceNumber
}

func getOrderData(invoiceNumber string, O arrOrders, n int) order {
	for i := 0; i < n; i++ {
		if O[i].invoice == invoiceNumber {
			return O[i]
		}
	}

	return order{}
}

func invoiceOrder(invoiceNumber string, O arrOrders, n int) {
	// Get the order data from the database based on the invoice number
	orderData := getOrderData(invoiceNumber, O, n)
	admin := orderData.data_staff
	customer := orderData.data_customer
	car := orderData.data_car

	// Total pembelian
	subtotal := orderData.data_car.price
	pajak := orderData.data_car.price * 10 / 100
	total := orderData.total

	// Tampilan invoice pembelian
	fmt.Println("---------------------------------------------------------------------")
	fmt.Println("                        INVOICE PEMBELIAN")
	fmt.Println("---------------------------------------------------------------------")
	fmt.Printf("Tanggal Pembelian: %s\n", orderData.date)
	fmt.Printf("Nomor Invoice: %s\n", orderData.invoice)
	fmt.Println()

	// Maksimal lebar kolom berdasar panjang nomor telepon
	maxWidth := len(customer.num_telp) + 2

	fmt.Println("---------------------------------------------------------------------")
	fmt.Printf("%-*s  %-*s\n", maxWidth*2+7, "Pembeli:", maxWidth, "Admin:")
	fmt.Println("---------------------------------------------------------------------")
	fmt.Printf("%-*s : %-*s  |   %-*s : %-s\n", maxWidth, "Nama", maxWidth, customer.name, maxWidth, "Nama", admin.name)
	fmt.Printf("%-*s : %-*s  |   %-*s : %-s\n", maxWidth, "Alamat", maxWidth, customer.address, maxWidth, "No. Telepon", admin.num_telp)
	fmt.Printf("%-*s : %-*s\n", maxWidth, "No. Telepon", maxWidth, customer.num_telp)
	fmt.Println("---------------------------------------------------------------------")
	fmt.Println()

	fmt.Println("---------------------------------------------------------------------")
	fmt.Println("No. | Mobil    | Jenis    | Tahun    | Harga")
	fmt.Println("---------------------------------------------------------------------")
	fmt.Printf("%-3d | %-8s | %-8s | %-8d | %-8d\n", 1, car.brand, car.model, car.year, car.price)
	fmt.Println("---------------------------------------------------------------------")
	fmt.Println()
	fmt.Printf("Subtotal     : %d\n", subtotal)
	fmt.Printf("Pajak (10%%)  : %d\n", pajak)
	fmt.Printf("Total        : %d\n", total)
	fmt.Println()
	fmt.Println("---------------------------------------------------------------------")
	fmt.Println("              Terima kasih atas pembelian Anda!")
	fmt.Println("---------------------------------------------------------------------")
}

func searchCar(C arrCars, n int, id_car int) (found bool, car car) {
	found = false
	i := 0
	for i < n && !found {
		if C[i].id == id_car {
			found = true
		}
		i++
	}
	return found, C[i-1]
}

func searchCustomerOrder(C arrCustomers, n int, id_customer int) (found bool, customer customer) {
	found = false
	i := 0
	for i < n && !found {
		if C[i].id == id_customer {
			found = true
		}
		i++
	}
	return found, C[i-1]
}
