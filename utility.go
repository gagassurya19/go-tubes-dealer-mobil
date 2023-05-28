package main

import (
	"os"
	"os/exec"
)

// clears the screen
func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// Function to sort arrCars by ID in ascending order
func sortByIDAscending(c *arrCars, n int) {
	for i := 0; i < len(c)-1; i++ {
		for j := 0; j < len(c)-i-1; j++ {
			if c[j].id > c[j+1].id {
				temp := c[j]
				c[j] = c[j+1]
				c[j+1] = temp
			}
		}
	}
}

// Function to sort arrCars by ID in descending order
func sortByIDDescending(c *arrCars, n int) {
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if (*c)[j].id < (*c)[j+1].id {
				temp := (*c)[j]
				(*c)[j] = (*c)[j+1]
				(*c)[j+1] = temp
			}
		}
	}
}

// Function to sort arrCars by brand in ascending order
func sortByBrandAscending(c *arrCars, n int) {
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if (*c)[j].brand > (*c)[j+1].brand {
				temp := (*c)[j]
				(*c)[j] = (*c)[j+1]
				(*c)[j+1] = temp
			}
		}
	}
}

// Function to sort arrCars by brand in descending order
func sortByBrandDescending(c *arrCars, n int) {
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if (*c)[j].brand < (*c)[j+1].brand {
				temp := (*c)[j]
				(*c)[j] = (*c)[j+1]
				(*c)[j+1] = temp
			}
		}
	}
}

// Function to sort arrCars by model in ascending order
func sortByModelAscending(c *arrCars, n int) {
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if (*c)[j].model > (*c)[j+1].model {
				temp := (*c)[j]
				(*c)[j] = (*c)[j+1]
				(*c)[j+1] = temp
			}
		}
	}
}

// Function to sort arrCars by model in descending order
func sortByModelDescending(c *arrCars, n int) {
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if (*c)[j].model < (*c)[j+1].model {
				temp := (*c)[j]
				(*c)[j] = (*c)[j+1]
				(*c)[j+1] = temp
			}
		}
	}
}

// Function to sort arrCars by year in ascending order
func sortByYearAscending(c *arrCars, n int) {
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if (*c)[j].year > (*c)[j+1].year {
				temp := (*c)[j]
				(*c)[j] = (*c)[j+1]
				(*c)[j+1] = temp
			}
		}
	}
}

// Function to sort arrCars by year in descending order
func sortByYearDescending(c *arrCars, n int) {
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if (*c)[j].year < (*c)[j+1].year {
				temp := (*c)[j]
				(*c)[j] = (*c)[j+1]
				(*c)[j+1] = temp
			}
		}
	}
}

// Function to sort arrCars by price in ascending order
func sortByPriceAscending(c *arrCars, n int) {
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if (*c)[j].price > (*c)[j+1].price {
				temp := (*c)[j]
				(*c)[j] = (*c)[j+1]
				(*c)[j+1] = temp
			}
		}
	}
}

// Function to sort arrCars by price in descending order
func sortByPriceDescending(c *arrCars, n int) {
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if (*c)[j].price < (*c)[j+1].price {
				temp := (*c)[j]
				(*c)[j] = (*c)[j+1]
				(*c)[j+1] = temp
			}
		}
	}
}
