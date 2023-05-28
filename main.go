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
	id                                 int
	name, num_telp, username, password string
}

type order struct {
	id            int
	data_customer customer
	data_car      car
	data_staff    staff
}

type arrCars [100]car
type arrCustomers [10]customer
type arrStaffs [10]staff
type arrOrders [10]order

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

	isSuccess, namaPetugas := login(staffs)

	for isSuccess {
		breakLoop = false
		for !breakLoop {
			home(namaPetugas) // HOME INTERFACE
			fmt.Scanln(&choice)
			switch choice {
			case 1: // CAR LIST
				var opt1 int
				var breaker bool = false
				for !breaker {
					carListInterface(cars, n_cars)
					fmt.Print("Pilihan: ")
					fmt.Scan(&opt1)
					switch opt1 {
					case 1:
						editCarInterface(&cars, n_cars)
					case 2:
						deleteCarInterface(&cars, &n_cars)
					case 3:
						breaker = true
					default:
						fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
						continue
					}
					fmt.Println()
				}
			case 2: // CAR SEARCH
				var opt2 int
				var breaker bool = false
				for !breaker {
					carSearchInterface()
					fmt.Print("Pilihan: ")
					fmt.Scanln(&opt2)
					switch opt2 {
					case 1:
						carSearchByBrand(cars, n_cars)
					case 2:
						carSearchByModel(cars, n_cars)
					case 3:
						carSearchByYear(cars, n_cars)
					case 4:
						breaker = true
					default:
						fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
						continue
					}
					fmt.Println()
				}
			case 3: // ADD NEW CAR
				var opt3 int
				var breaker bool = false
				for !breaker {
					addNewCarInterface()
					fmt.Print("Pilihan: ")
					fmt.Scanln(&opt3)
					switch opt3 {
					case 1:
						addNewCar(&cars, &n_cars)
					case 2:
						addBulkCar(&cars, &n_cars)
					case 3:
						breaker = true
					default:
						fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
						continue
					}
					fmt.Println()
				}
			case 4: // ORDER CAR
				orderCar()
			case 5: // SALES REPORT
				generateSalesReport()
			case 6: // EXIT APP
				exitApp()
				breakLoop = true
				isSuccess, namaPetugas = login(staffs)
			default:
				fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
			}
			fmt.Println()
		}
	}
}
