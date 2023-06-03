package main

type car struct {
	id, price, year int
	brand, model    string
}

type customer struct {
	id                               int
	name, num_ktp, num_telp, address string
}

type staff struct {
	id                                 int
	name, num_telp, username, password string
}

type order struct {
	id            int
	tax           int
	subtotal      int
	total         int
	invoice       string
	date          string
	data_customer customer
	data_car      car
	data_staff    staff
}

type arrCars [100]car
type arrCustomers [10]customer
type arrStaffs [10]staff
type arrOrders [10]order
