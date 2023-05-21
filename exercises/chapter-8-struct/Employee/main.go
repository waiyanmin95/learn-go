package main

import (
	"fmt"

	"github.com/ivan-amity/learn-go/exercises/chapter-8-struct/Employee/magazine"
)

func main() {
	//address := magazine.Address{"Street111", "Yangon", "Myanmar", "11041"}
	var employee magazine.Employee
	employee.Name = "Ivan"
	employee.Salary = 89000
	//employee.HomeAddress = magazine.Address{Street: "Narnattaw Street", City: "Yangon", State: "Myanmar", PostalCode: "11041"}

	employee.City = "Mandalay"
	employee.PostalCode = "10250"

	fmt.Println("Employee Name:", employee.Name)
	fmt.Println("Salary:", employee.Salary)
	fmt.Println("City:", employee.City)
	fmt.Println("Postal Code:", employee.PostalCode)
}
