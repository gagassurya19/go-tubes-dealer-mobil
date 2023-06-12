package main

type car struct {
	id, price, year int
	brand, model    string
	totalTerjual    int
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

const MAXCARS = 100
const MAXCUSTOMERS = 10
const MAXSTAFFS = 10
const MAXORDERS = 100
const MAXTOPS = 3

type arrCars [MAXCARS]car
type arrCustomers [MAXCUSTOMERS]customer
type arrStaffs [MAXSTAFFS]staff
type arrOrders [MAXORDERS]order
