package main

func dataCars(C *arrCars, n *int) {
	*C = arrCars{
		{1, 20000, 2020, "toyota", "camry"},
		{2, 15000, 2018, "honda", "civic"},
		{3, 30000, 2021, "bmw", "x5"},
		{4, 25000, 2019, "ford", "mustang"},
		{5, 18000, 2017, "chevrolet", "corvette"},
		{6, 22000, 2016, "nissan", "altima"},
		{7, 19000, 2015, "audi", "a4"},
		{8, 28000, 2014, "mercedes-benz", "c-class"},
		{9, 23000, 2013, "volkswagen", "golf"},
		{10, 17000, 2012, "subaru", "impreza"},
		{11, 21000, 2011, "mazda", "cx-5"},
		{12, 26000, 2010, "lexus", "rx"},
		{13, 24000, 2019, "porsche", "911"},
		{14, 29000, 2018, "jeep", "wrangler"},
		{15, 20000, 2017, "kia", "soul"},
		{16, 17000, 2016, "hyundai", "elantra"},
		{17, 18000, 2015, "chrysler", "300"},
		{18, 22000, 2014, "cadillac", "escalade"},
		{19, 19000, 2013, "buick", "encore"},
		{20, 25000, 2012, "lincoln", "navigator"},
		{21, 20000, 2020, "toyota", "corolla"}, // Toyota Corolla
		{22, 18000, 2020, "toyota", "camry"},   // Toyota Camry
		{23, 22000, 2020, "toyota", "rav4"},    // Toyota RAV4
		{24, 24000, 2018, "honda", "civic"},    // Honda Civic
		{25, 23000, 2018, "honda", "accord"},   // Honda Accord
		{26, 21000, 2018, "honda", "cr-v"},     // Honda CR-V
		{27, 30000, 2021, "bmw", "x5"},         // BMW X5
		{28, 32000, 2021, "bmw", "3-series"},   // BMW 3 Series
		{29, 28000, 2021, "bmw", "5-series"},   // BMW 5 Series
		{30, 25000, 2019, "ford", "mustang"},   // Ford Mustang
	}
	*n = 30
}

func dataCustomers(C *arrCustomers, n *int) {
	*C = arrCustomers{
		{1, "Surya", "3504020001", "082139456688", "jl. raya"},
		{2, "Yohanes", "9876543210", "082139456688", "no. 123"},
		{3, "Laksana", "6543210987", "082139456688", "jl. 123"},
	}
	*n = 3
}

func dataStaffs(C *arrStaffs, n *int) {
	*C = arrStaffs{
		{1, "Gagas", "1234567890", "gagas", "123"},
		{2, "Maurich", "9876543210", "maurich", "456"},
		{3, "1", "1", "1", "1"},
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
