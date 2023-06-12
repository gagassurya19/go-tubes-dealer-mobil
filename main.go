package main

import (
	"fmt"
)

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

	// LOGIN INTERFACE
	isSuccess, dataStaff := login(staffs)

	for isSuccess {
		breakLoop = false
		for !breakLoop {
			home(dataStaff.name) // HOME INTERFACE
			fmt.Scanln(&choice)
			switch choice {
			case 1: // CAR LIST
				var opt int
				var breaker bool = false
				for !breaker {
					carListInterface(cars, n_cars)
					fmt.Print("Pilihan: ")
					fmt.Scan(&opt)
					switch opt {
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
				var opt int
				var breaker bool = false
				for !breaker {
					carSearchInterface()
					fmt.Print("Pilihan: ")
					fmt.Scanln(&opt)
					switch opt {
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
				var opt int
				var breaker bool = false
				for !breaker {
					addNewCarInterface()
					fmt.Print("Pilihan: ")
					fmt.Scanln(&opt)
					switch opt {
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
				placeOrder(&orders, &n_orders, &cars, n_cars, customers, n_customers, dataStaff, staffs)
			case 5: // CRUD CUSTOMER
				var opt int
				var breaker bool = false
				for !breaker {
					viewCustomer()
					fmt.Scanln(&opt)
					switch opt {
					case 1:
						listCustomer(customers, n_customers)
					case 2:
						searchCustomer(&customers, &n_customers)
					case 3:
						addCustomer(&customers, &n_customers)
					case 4:
						breaker = true
					default:
						fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
						continue
					}
					fmt.Println()
				}
			case 6: // CRUD STAFF
				var opt int
				var breaker bool = false
				for !breaker {
					viewStaff()
					fmt.Scanln(&opt)
					switch opt {
					case 1:
						listStaff(staffs, n_staffs)
					case 2:
						searchStaff(&staffs, &n_staffs)
					case 3:
						addStaff(&staffs, &n_staffs)
					case 4:
						breaker = true
					default:
						fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
						continue
					}
					fmt.Println()
				}
			case 7: // SALES REPORT
				var opt int
				var breaker bool = false
				for !breaker {
					viewGenerateSalesReport()
					fmt.Scanln(&opt)
					switch opt {
					case 1:
						generateTopThree(cars, n_cars)
					case 2:
						viewHistoryOrders(orders, n_orders)
					case 3:
						findHistoryOrder(orders, n_orders)
					case 4:
						breaker = true
					default:
						fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
						continue
					}
					fmt.Println()
				}
			case 8: // EXIT APP
				exitApp()
				breakLoop = true
				isSuccess, dataStaff = login(staffs)
			default:
				fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
			}
			fmt.Println()
		}
	}
}
