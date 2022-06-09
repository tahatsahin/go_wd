package main

import (
	"fmt"
)

func main() {
	var person1 = &Employee{
		Id:     1001,
		Name:   "XXX",
		Salary: 1000,
	}

	var person2 = &Employee{
		Id:     1002,
		Name:   "YYY",
		Salary: 5000,
	}

	var employees []*Employee

	employees = append(employees, person1)
	employees = append(employees, person2)

	var capEmp = &Employees{Employees: employees}

	fmt.Println(capEmp)
}
