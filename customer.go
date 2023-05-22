package main

import "fmt"

func addCustomer(C *arrCustomers, n *int) {
	var in customer

	fmt.Print("Input name:")
	fmt.Scanln(&in.name)
	fmt.Print("Input KTP number:")
	fmt.Scanln(&in.num_ktp)
	fmt.Print("Input Telphone number:")
	fmt.Scanln(&in.num_telp)
	fmt.Print("Input address:")
	fmt.Scanln(&in.address)

	in.id = *n
	C[*n] = in

	*n++
}

func editCustomer(C *arrCustomers, n int) {
	var id int

	fmt.Print("Input customer ID to edit:")
	fmt.Scanln(&id)

	if id >= 0 && id < n {
		var in customer

		fmt.Print("Input name:")
		fmt.Scanln(&in.name)
		fmt.Print("Input KTP number:")
		fmt.Scanln(&in.num_ktp)
		fmt.Print("Input Telphone number:")
		fmt.Scanln(&in.num_telp)
		fmt.Print("Input address:")
		fmt.Scanln(&in.address)

		in.id = id
		C[id] = in

		fmt.Println("Customer successfully edited!")
	} else {
		fmt.Println("Invalid customer ID!")
	}
}

func deleteCustomer(C *arrCustomers, n *int) {
	var id int

	fmt.Print("Input customer ID to delete:")
	fmt.Scanln(&id)

	if id >= 0 && id < *n {
		for i := id; i < *n-1; i++ {
			C[i] = C[i+1]
		}

		*n--
		fmt.Println("Customer successfully deleted!")
	} else {
		fmt.Println("Invalid customer ID!")
	}
}
