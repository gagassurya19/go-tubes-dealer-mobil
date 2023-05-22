package main

import (
	"fmt"
)

type car struct {
	id, price, year int
	brand, model    string
}

type customer struct {
	id, num_ktp             int
	name, num_telp, address string
}

type staff struct {
	id             int
	name, num_telp string
}

type order struct {
	id            int
	data_customer customer
	data_car      car
	data_staff    staff
}

type arrCars [10]car
type arrCustomers [10]customer
type arrStaffs [10]staff
type arrOrders [10]order

func dataCars(C *arrCars, n *int) {
	*C = arrCars{
		{1, 20000, 2020, "Toyota", "Camry"},
		{2, 15000, 2018, "Honda", "Civic"},
		{3, 30000, 2021, "BMW", "X5"},
	}
	*n = 3
}

func dataCustomers(C *arrCustomers, n *int) {
	*C = arrCustomers{
		{1, 123456, "John Doe", "1234567890", "123 Main St"},
		{2, 987654, "Jane Smith", "9876543210", "456 Elm St"},
		{3, 654321, "Michael Johnson", "6543210987", "789 Oak St"},
	}
	*n = 3
}

func dataStaffs(C *arrStaffs, n *int) {
	*C = arrStaffs{
		{1, "Emily Brown", "1234567890"},
		{2, "James Wilson", "9876543210"},
		{3, "Sophia Davis", "6543210987"},
	}
	*n = 3
}

func dataOrders(C *arrOrders, n *int, customers arrCustomers, cars arrCars, staffs arrStaffs) {
	*C = arrOrders{
		{1, customers[0], cars[1], staffs[0]},
		{2, customers[1], cars[2], staffs[1]},
		{3, customers[2], cars[0], staffs[2]},
	}
	*n = 3
}

func main() {
	var cars arrCars
	var customers arrCustomers
	var staffs arrStaffs
	var orders arrOrders
	var choice, n_cars, n_customers, n_staffs, n_orders int
	var breakLoop bool = false
	// DATA DUMMY DEFAULT
	dataCars(&cars, &n_cars)
	dataCustomers(&customers, &n_customers)
	dataStaffs(&staffs, &n_staffs)
	dataOrders(&orders, &n_orders, customers, cars, staffs)
	// TEMPORARY - DELETE AFTER DONE
	fmt.Println(n_customers)
	fmt.Println(n_staffs)
	fmt.Println(n_orders)

	for {
		home() // HOME INTERFACE
		fmt.Scanln(&choice)
		switch choice {
		case 1: // CAR LIST
			var opt1 int
			for {
				carListInterface(cars, n_cars)
				fmt.Print("Pilihan: ")
				fmt.Scan(&opt1)
				switch opt1 {
				case 1:
					editCarInterface(&cars, n_cars)
				case 2:
					deleteCarInterface(&cars, &n_cars)
				case 3:
					breakLoop = true
				default:
					fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
					continue
				}
				fmt.Println()
				if breakLoop {
					breakLoop = false
					break
				}
			}
		case 2: // CAR SEARCH
			var opt2 int
			for {
				carSearchInterface()
				fmt.Print("Pilihan: ")
				fmt.Scanln(&opt2)
				switch opt2 {
				case 1:
					carSearchByBrandInterface()
				case 2:
					carSearchByModelInterface()
				case 3:
					carSearchByYearInterface()
				case 4:
					breakLoop = true
				default:
					fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
					continue
				}
				fmt.Println()
				if breakLoop {
					breakLoop = false
					break
				}
			}
		case 3: // ADD NEW CAR
			addNewCarInterface()
		case 4: // ORDER CAR
			orderCar()
		case 5: // SALES REPORT
			generateSalesReport()
		case 6: // EXIT APP
			exitApp()
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
		fmt.Println()
	}
}
