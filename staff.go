package main

import "fmt"

func addStaff(C *arrStaffs, n *int) {
	var in staff

	fmt.Print("Input name:")
	fmt.Scanln(&in.name)
	fmt.Print("Input Telphone number:")
	fmt.Scanln(&in.num_telp)

	in.id = *n
	C[*n] = in

	*n++
}

func editStaff(C *arrStaffs, n int) {
	var id int

	fmt.Print("Input staff ID to edit:")
	fmt.Scanln(&id)

	if id >= 0 && id < n {
		var in staff

		fmt.Print("Input name:")
		fmt.Scanln(&in.name)
		fmt.Print("Input Telphone number:")
		fmt.Scanln(&in.num_telp)

		in.id = id
		C[id] = in

		fmt.Println("Staff successfully edited!")
	} else {
		fmt.Println("Invalid staff ID!")
	}
}

func deleteStaff(C *arrStaffs, n *int) {
	var id int

	fmt.Print("Input staff ID to delete:")
	fmt.Scanln(&id)

	if id >= 0 && id < *n {
		for i := id; i < *n-1; i++ {
			C[i] = C[i+1]
		}

		*n--
		fmt.Println("Staff successfully deleted!")
	} else {
		fmt.Println("Invalid staff ID!")
	}
}
