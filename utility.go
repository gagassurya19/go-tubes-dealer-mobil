package main

import (
	"os"
	"os/exec"
)

// THIS FILE CONTAINS UTILITY FUNCTIONS
// SUCH AS CLEAR SCREEN, SORTING, ETC.
// THAT CAN BE USED IN ANY OTHER FILE

// CLEAR SCREEN
func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// CAR
func sortByIDAscending(c *arrCars, n int) {
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if c[j].id > c[j+1].id {
				temp := c[j]
				c[j] = c[j+1]
				c[j+1] = temp
			}
		}
	}
}

func sortByIDDescending(c *arrCars, n int) {
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if c[j].id < c[j+1].id {
				temp := c[j]
				c[j] = c[j+1]
				c[j+1] = temp
			}
		}
	}
}

func sortByBrandAscending(c *arrCars, n int) {
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if c[j].brand > c[j+1].brand {
				temp := c[j]
				c[j] = c[j+1]
				c[j+1] = temp
			}
		}
	}
}

func sortByBrandDescending(c *arrCars, n int) {
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if c[j].brand < c[j+1].brand {
				temp := c[j]
				c[j] = c[j+1]
				c[j+1] = temp
			}
		}
	}
}

func sortByModelAscending(c *arrCars, n int) {
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if c[j].model > c[j+1].model {
				temp := c[j]
				c[j] = c[j+1]
				c[j+1] = temp
			}
		}
	}
}

func sortByModelDescending(c *arrCars, n int) {
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if c[j].model < c[j+1].model {
				temp := c[j]
				c[j] = c[j+1]
				c[j+1] = temp
			}
		}
	}
}

func sortByYearAscending(c *arrCars, n int) {
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if c[j].year > c[j+1].year {
				temp := c[j]
				c[j] = c[j+1]
				c[j+1] = temp
			}
		}
	}
}

func sortByYearDescending(c *arrCars, n int) {
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if c[j].year < c[j+1].year {
				temp := c[j]
				c[j] = c[j+1]
				c[j+1] = temp
			}
		}
	}
}

func sortByPriceAscending(c *arrCars, n int) {
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if c[j].price > c[j+1].price {
				temp := c[j]
				c[j] = c[j+1]
				c[j+1] = temp
			}
		}
	}
}

func sortByPriceDescending(c *arrCars, n int) {
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if c[j].price < c[j+1].price {
				temp := c[j]
				c[j] = c[j+1]
				c[j+1] = temp
			}
		}
	}
}

// CUSTOMER
func sortCustomerByIDAsc(C *arrCustomers, n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if C[j].id > C[j+1].id {
				temp := C[j]
				C[j] = C[j+1]
				C[j+1] = temp
			}
		}
	}
}

func sortCustomerByIDDesc(C *arrCustomers, n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if C[j].id < C[j+1].id {
				temp := C[j]
				C[j] = C[j+1]
				C[j+1] = temp
			}
		}
	}
}

func sortCustomerByNameAsc(C *arrCustomers, n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if C[j].name > C[j+1].name {
				temp := C[j]
				C[j] = C[j+1]
				C[j+1] = temp
			}
		}
	}
}

func sortCustomerByNameDesc(C *arrCustomers, n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if C[j].name < C[j+1].name {
				temp := C[j]
				C[j] = C[j+1]
				C[j+1] = temp
			}
		}
	}
}

func sortCustomerByKTPAsc(C *arrCustomers, n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if C[j].num_ktp > C[j+1].num_ktp {
				temp := C[j]
				C[j] = C[j+1]
				C[j+1] = temp
			}
		}
	}
}

func sortCustomerByKTPDesc(C *arrCustomers, n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if C[j].num_ktp < C[j+1].num_ktp {
				temp := C[j]
				C[j] = C[j+1]
				C[j+1] = temp
			}
		}
	}
}

func sortCustomerByTelpAsc(C *arrCustomers, n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if C[j].num_telp > C[j+1].num_telp {
				temp := C[j]
				C[j] = C[j+1]
				C[j+1] = temp
			}
		}
	}
}

func sortCustomerByTelpDesc(C *arrCustomers, n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if C[j].num_telp < C[j+1].num_telp {
				temp := C[j]
				C[j] = C[j+1]
				C[j+1] = temp
			}
		}
	}
}

func sortCustomerByAddressAsc(C *arrCustomers, n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if C[j].address > C[j+1].address {
				temp := C[j]
				C[j] = C[j+1]
				C[j+1] = temp
			}
		}
	}
}

func sortCustomerByAddressDesc(C *arrCustomers, n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if C[j].address < C[j+1].address {
				temp := C[j]
				C[j] = C[j+1]
				C[j+1] = temp
			}
		}
	}
}

// STAFF
func sortStaffByIDAsc(S *arrStaffs, n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n-1; j++ {
			if S[j].id > S[j+1].id {
				temp := S[j]
				S[j] = S[j+1]
				S[j+1] = temp
			}
		}
	}
}

func sortStaffByIDDesc(S *arrStaffs, n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n-1; j++ {
			if S[j].id < S[j+1].id {
				temp := S[j]
				S[j] = S[j+1]
				S[j+1] = temp
			}
		}
	}
}

func sortStaffByNameAsc(S *arrStaffs, n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n-1; j++ {
			if S[j].name > S[j+1].name {
				temp := S[j]
				S[j] = S[j+1]
				S[j+1] = temp
			}
		}
	}
}

func sortStaffByNameDesc(S *arrStaffs, n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n-1; j++ {
			if S[j].name < S[j+1].name {
				temp := S[j]
				S[j] = S[j+1]
				S[j+1] = temp
			}
		}
	}
}

func sortStaffByUsernameAsc(S *arrStaffs, n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n-1; j++ {
			if S[j].username > S[j+1].username {
				temp := S[j]
				S[j] = S[j+1]
				S[j+1] = temp
			}
		}
	}
}

func sortStaffByUsernameDesc(S *arrStaffs, n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n-1; j++ {
			if S[j].username < S[j+1].username {
				temp := S[j]
				S[j] = S[j+1]
				S[j+1] = temp
			}
		}
	}
}

func sortStaffByTelpAsc(S *arrStaffs, n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n-1; j++ {
			if S[j].num_telp > S[j+1].num_telp {
				temp := S[j]
				S[j] = S[j+1]
				S[j+1] = temp
			}
		}
	}
}

func sortStaffByTelpDesc(S *arrStaffs, n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n-1; j++ {
			if S[j].num_telp < S[j+1].num_telp {
				temp := S[j]
				S[j] = S[j+1]
				S[j+1] = temp
			}
		}
	}
}
