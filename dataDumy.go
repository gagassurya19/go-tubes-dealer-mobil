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
		{1, 20000, 2020, "toyota", "camry", 1},
		{2, 15000, 2018, "honda", "civic", 1},
		{3, 30000, 2021, "bmw", "x5", 1},
	}
	*n = 3

	storeDataCars(*C, *n)
	loadDataCars(C, n)
	fmt.Println(C)
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
		{
			id:            1,
			tax:           10,
			subtotal:      20000,
			total:         22000,
			invoice:       "INV-001",
			date:          "2023-06-02",
			data_customer: customers[0],
			data_car:      cars[1],
			data_staff:    staffs[0],
		},
		{
			id:            2,
			tax:           10,
			subtotal:      30000,
			total:         33000,
			invoice:       "INV-002",
			date:          "2023-06-03",
			data_customer: customers[1],
			data_car:      cars[2],
			data_staff:    staffs[1],
		},
		{
			id:            3,
			tax:           10,
			subtotal:      25000,
			total:         27500,
			invoice:       "INV-003",
			date:          "2023-06-04",
			data_customer: customers[2],
			data_car:      cars[0],
			data_staff:    staffs[2],
		},
	}
	*n = 3

	storeDataOrders(*C, *n, customers, cars, staffs)
	loadDataOrders(C, n, customers, cars, staffs)
}

// STORE DATA TO FILE
func storeDataCars(C arrCars, n int) {
	// Check if there is no data in the array
	if n == 0 || n == 3 {
		fmt.Println("STORE CAR: No data to store")
		return
	}

	// Create the database folder if it doesn't exist
	if _, err := os.Stat(databaseFolder); os.IsNotExist(err) {
		err := os.Mkdir(databaseFolder, 0755)
		if err != nil {
			fmt.Println("STORE CAR: Error creating database folder:", err)
			return
		}
	}

	filePath := filepath.Join(databaseFolder, "dataCars.txt")

	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("STORE CAR: Error:", err)
		return
	}
	defer file.Close()

	for i := 0; i < n; i++ {
		fmt.Fprintf(file, "%d_%d_%d_%s_%s_%d\n", C[i].id, C[i].price, C[i].year, C[i].brand, C[i].model, C[i].totalTerjual)
	}
}

func storeDataCustomers(C arrCustomers, n int) {
	// Check if there is no data in the array
	if n == 0 || n == 3 {
		fmt.Println("STORE CUSTOMER: No data to store")
		return
	}

	// Create the database folder if it doesn't exist
	if _, err := os.Stat(databaseFolder); os.IsNotExist(err) {
		err := os.Mkdir(databaseFolder, 0755)
		if err != nil {
			fmt.Println("STORE CUSTOMER: Error creating database folder:", err)
			return
		}
	}

	filePath := filepath.Join(databaseFolder, "dataCustomers.txt")
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("STORE CUSTOMER: Error:", err)
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
		fmt.Println("STORE STAFF: No data to store")
		return
	}

	// Create the database folder if it doesn't exist
	if _, err := os.Stat(databaseFolder); os.IsNotExist(err) {
		err := os.Mkdir(databaseFolder, 0755)
		if err != nil {
			fmt.Println("STORE STAFF: Error creating database folder:", err)
			return
		}
	}

	filePath := filepath.Join(databaseFolder, "dataStaffs.txt")
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("STORE STAFF: Error: ", err)
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
		fmt.Println("STORE ORDERS: No data to store")
		return
	}

	// Create the database folder if it doesn't exist
	if _, err := os.Stat(databaseFolder); os.IsNotExist(err) {
		err := os.Mkdir(databaseFolder, 0755)
		if err != nil {
			fmt.Println("STORE ORDERS: Error creating database folder:", err)
			return
		}
	}

	filePath := filepath.Join(databaseFolder, "dataOrders.txt")
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("STORE ORDERS: Error: ", err)
		return
	}
	defer file.Close()
	for i := 0; i < n; i++ {
		order := O[i]
		customer := order.data_customer
		car := order.data_car
		staff := order.data_staff

		// Format the order data into a string
		orderData := fmt.Sprintf("%d_%d_%d_%d_%s_%s,%d_%s_%s_%s,%d_%s_%s_%d_%d,%d_%s_%s_%s\n",
			order.id,
			order.tax,
			order.subtotal,
			order.total,
			order.date,
			order.invoice,
			customer.id,
			customer.name,
			customer.num_ktp,
			customer.num_telp,
			car.id,
			car.brand,
			car.model,
			car.year,
			car.price,
			staff.id,
			staff.name,
			staff.num_telp,
			staff.username)

		// Write the order data to the file
		_, err := file.WriteString(orderData)
		if err != nil {
			fmt.Println("STORE ORDERS: Error writing order data:", err)
			return
		}
	}
}

// LOAD DATA FROM FILE
func loadDataCars(C *arrCars, n *int) {
	file, err := os.Open("./database/dataCars.txt")
	if err != nil {
		fmt.Println("LOAD CAR: Error:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() && i < len(C) {
		line := scanner.Text()
		fields := strings.Split(line, "_")
		if len(fields) != 6 {
			continue
		}

		id, _ := strconv.Atoi(fields[0])
		price, _ := strconv.Atoi(fields[1])
		year, _ := strconv.Atoi(fields[2])
		brand := fields[3]
		model := fields[4]
		totalTerjual, _ := strconv.Atoi(fields[5])

		c := car{
			id:           id,
			price:        price,
			year:         year,
			brand:        brand,
			model:        model,
			totalTerjual: totalTerjual,
		}

		C[i] = c
		i++
	}
	*n = i
}

func loadDataCustomers(C *arrCustomers, n *int) {
	file, err := os.Open("./database/dataCustomers.txt")
	if err != nil {
		fmt.Println("LOAD CUSTOMER: Error:", err)
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
		fmt.Println("LOAD STAFF: Error:", err)
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

func loadDataOrders(O *arrOrders, n *int, customers arrCustomers, cars arrCars, staffs arrStaffs) {
	file, err := os.Open("./database/dataOrders.txt")
	if err != nil {
		fmt.Println("LOAD ORDER: Error:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	i := 0

	for scanner.Scan() && i < len(O) {
		line := scanner.Text()

		// Extract the fields from the line
		fields := strings.Split(line, ",")
		if len(fields) != 4 {
			continue
		}

		orderFields := strings.Split(fields[0], "_")
		if len(orderFields) != 6 {
			continue
		}
		id, _ := strconv.Atoi(orderFields[0])
		orderTax, _ := strconv.Atoi(orderFields[1])
		orderSubtotal, _ := strconv.Atoi(orderFields[2])
		orderTotal, _ := strconv.Atoi(orderFields[3])
		orderDate := orderFields[4]
		orderInvoice := orderFields[5]

		// Extract the customer data from the line
		customerFields := strings.Split(fields[1], "_")
		if len(customerFields) != 4 {
			continue
		}
		customerID, _ := strconv.Atoi(customerFields[0])

		// Extract the car data from the line
		carFields := strings.Split(fields[2], "_")
		if len(carFields) != 5 {
			continue
		}
		carID, _ := strconv.Atoi(carFields[0])

		// Extract the staff data from the line
		staffFields := strings.Split(fields[3], "_")
		if len(staffFields) != 4 {
			continue
		}
		staffID, _ := strconv.Atoi(staffFields[0])

		// Find the corresponding customer, car, and staff objects based on their IDs
		var customer customer
		for _, c := range customers {
			if c.id == customerID {
				customer = c
				break
			}
		}

		var car car
		for _, c := range cars {
			if c.id == carID {
				car = c
				break
			}
		}

		var staff staff
		for _, s := range staffs {
			if s.id == staffID {
				staff = s
				break
			}
		}

		// Create the order object and assign the customer, car, and staff references
		o := order{
			id:            id,
			tax:           orderTax,
			subtotal:      orderSubtotal,
			total:         orderTotal,
			date:          orderDate,
			invoice:       orderInvoice,
			data_customer: customer,
			data_car:      car,
			data_staff:    staff,
		}

		O[i] = o
		i++
	}

	*n = i
}
