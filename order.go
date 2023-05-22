package main

import "fmt"

func newOrder(C *arrOrders, n *int, customers arrCustomers, cars arrCars, staffs arrStaffs) {
	var in order
	var id_customer, id_car, id_staff int

	fmt.Print("Input id-customer:")
	fmt.Scanln(&id_customer)
	fmt.Print("Input id-car:")
	fmt.Scanln(&id_car)
	fmt.Print("Input id-staff:")
	fmt.Scanln(&id_staff)

	in = order{
		id:            *n,
		data_customer: customers[id_customer],
		data_car:      cars[id_car],
		data_staff:    staffs[id_staff],
	}

	C[*n] = in

	*n++
}
