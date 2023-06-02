package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const databaseFolder = "./database"

func dataCars(C *arrCars, n *int) {
	*C = arrCars{
		{1, 20000, 2020, "toyota", "camry"},
		{2, 15000, 2018, "honda", "civic"},
		{3, 30000, 2021, "bmw", "x5"},
	}
	*n = 3

	storeDataCars(*C, *n)
	loadDataCars(C, n)
}

func dataCustomers(C *arrCustomers, n *int) {
	*C = arrCustomers{
		{1, "Surya", "3504020001", "082139456688", "jl. raya"},
		{2, "Yohanes", "9876543210", "082139456688", "no. 123"},
		{3, "Laksana", "6543210987", "082139456688", "jl. 123"},
	}
	*n = 3

	storeDataCustomers(*C, *n)
	loadDataCustomers(C, n)
	fmt.Println(*n)
}

func dataStaffs(C *arrStaffs, n *int) {
	*C = arrStaffs{
		{1, "Gagas", "1234567890", "gagas", "123"},
		{2, "Maurich", "9876543210", "maurich", "456"},
		{3, "1", "1", "1", "1"},
	}
	*n = 3

	storeDataStaffs(*C, *n)
	loadDataStaffs(C, n)
}

func dataOrders(C *arrOrders, n *int, customers arrCustomers, cars arrCars, staffs arrStaffs) {
	*C = arrOrders{
		{1, customers[0], cars[1], staffs[0]},
		{2, customers[1], cars[2], staffs[1]},
		{3, customers[2], cars[0], staffs[2]},
	}
	*n = 3

	storeDataOrders(*C, *n, customers, cars, staffs)
	loadDataOrders(C, n, customers, cars, staffs)
}

// STORE DATA TO FILE
func storeDataCars(C arrCars, n int) {
	// Check if there is no data in the array
	if n == 0 || n == 3 {
		fmt.Println("No data to store")
		return
	}

	// Create the database folder if it doesn't exist
	if _, err := os.Stat(databaseFolder); os.IsNotExist(err) {
		err := os.Mkdir(databaseFolder, 0755)
		if err != nil {
			fmt.Println("Error creating database folder:", err)
			return
		}
	}

	filePath := filepath.Join(databaseFolder, "dataCars.txt")

	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	for i := 0; i < n; i++ {
		fmt.Fprintf(file, "%d_%d_%d_%s_%s\n", C[i].id, C[i].price, C[i].year, C[i].brand, C[i].model)
	}
}

func storeDataCustomers(C arrCustomers, n int) {
	// Check if there is no data in the array
	if n == 0 || n == 3 {
		fmt.Println("No data to store")
		return
	}

	// Create the database folder if it doesn't exist
	if _, err := os.Stat(databaseFolder); os.IsNotExist(err) {
		err := os.Mkdir(databaseFolder, 0755)
		if err != nil {
			fmt.Println("Error creating database folder:", err)
			return
		}
	}

	filePath := filepath.Join(databaseFolder, "dataCustomers.txt")
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()
	for i := 0; i < n; i++ {
		fmt.Fprintf(file, "%d_%s_%s_%s_%s\n", C[i].id, C[i].name, C[i].num_ktp, C[i].num_telp, C[i].address)
	}
}

func storeDataStaffs(C arrStaffs, n int) {
	// Check if there is no data in the array
	if n == 0 || n == 3 {
		fmt.Println("No data to store")
		return
	}

	// Create the database folder if it doesn't exist
	if _, err := os.Stat(databaseFolder); os.IsNotExist(err) {
		err := os.Mkdir(databaseFolder, 0755)
		if err != nil {
			fmt.Println("Error creating database folder:", err)
			return
		}
	}

	filePath := filepath.Join(databaseFolder, "dataStaffs.txt")
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer file.Close()
	for i := 0; i < n; i++ {
		fmt.Fprintf(file, "%d_%s_%s_%s_%s\n", C[i].id, C[i].name, C[i].num_telp, C[i].username, C[i].password)
	}
}

func storeDataOrders(O arrOrders, n int, C arrCustomers, Car arrCars, S arrStaffs) {
	// Check if there is no data in the array
	if n == 0 || n == 3 {
		fmt.Println("No data to store")
		return
	}

	// Create the database folder if it doesn't exist
	if _, err := os.Stat(databaseFolder); os.IsNotExist(err) {
		err := os.Mkdir(databaseFolder, 0755)
		if err != nil {
			fmt.Println("Error creating database folder:", err)
			return
		}
	}

	filePath := filepath.Join(databaseFolder, "dataOrders.txt")
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer file.Close()
	for i := 0; i < n; i++ {
		fmt.Fprintln(file, O[i].id, C[i], Car[i], S[i])
	}
}

// LOAD DATA FROM FILE
func loadDataCars(C *arrCars, n *int) {
	file, err := os.Open("./database/dataCars.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() && i < len(C) {
		line := scanner.Text()
		fields := strings.Split(line, "_")
		if len(fields) != 5 {
			continue
		}

		id, _ := strconv.Atoi(fields[0])
		price, _ := strconv.Atoi(fields[1])
		year, _ := strconv.Atoi(fields[2])
		brand := fields[3]
		model := fields[4]

		c := car{
			id:    id,
			price: price,
			year:  year,
			brand: brand,
			model: model,
		}

		C[i] = c
		i++
	}
	*n = i
}

func loadDataCustomers(C *arrCustomers, n *int) {
	file, err := os.Open("./database/dataCustomers.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() && i < len(C) {
		line := scanner.Text()
		fields := strings.Split(line, "_")
		if len(fields) != 5 {
			continue
		}

		id, _ := strconv.Atoi(fields[0])
		name := fields[1]
		num_ktp := fields[2]
		num_telp := fields[3]
		address := fields[4]

		c := customer{
			id:       id,
			name:     name,
			num_ktp:  num_ktp,
			num_telp: num_telp,
			address:  address,
		}

		C[i] = c
		i++
	}
	*n = i
}

func loadDataStaffs(C *arrStaffs, n *int) {
	file, err := os.Open("./database/dataStaffs.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() && i < len(C) {
		line := scanner.Text()
		fields := strings.Split(line, "_")
		if len(fields) != 5 {
			continue
		}

		id, _ := strconv.Atoi(fields[0])
		name := fields[1]
		num_telp := fields[2]
		username := fields[3]
		password := fields[4]

		c := staff{
			id:       id,
			name:     name,
			num_telp: num_telp,
			username: username,
			password: password,
		}

		C[i] = c
		i++
	}
	*n = i
}

func loadDataOrders(O *arrOrders, n *int, C arrCustomers, Car arrCars, S arrStaffs) {
	file, err := os.Open("./database/dataOrders.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() && i < len(O) {
		line := scanner.Text()
		fields := strings.Split(line, " ")
		if len(fields) != 4 {
			continue
		}

		id, _ := strconv.Atoi(fields[0])
		customerID, _ := strconv.Atoi(fields[1])
		carID, _ := strconv.Atoi(fields[2])
		staffID, _ := strconv.Atoi(fields[3])

		// Find the corresponding customer, car, and staff objects based on their IDs
		var customer customer
		for _, c := range C {
			if c.id == customerID {
				customer = c
				break
			}
		}

		var car car
		for _, c := range Car {
			if c.id == carID {
				car = c
				break
			}
		}

		var staff staff
		for _, s := range S {
			if s.id == staffID {
				staff = s
				break
			}
		}

		// Create the order object and assign the customer, car, and staff references
		o := order{
			id:            id,
			data_customer: customer,
			data_car:      car,
			data_staff:    staff,
		}

		O[i] = o
		i++
	}
	*n = i
}
