package main

import "fmt"

type Person struct {
	name string
	age  uint8
}

type Employee struct {
	id int
}

type FullTimeEmployee struct {
	Person
	Employee
	endDate string
}

func (ftEmployee FullTimeEmployee) getMessage() string {
	return "full time employee"
}

func (temporaryEmployee TemporaryEmployee) getMessage() string {
	return "temporary employee"
}

type TemporaryEmployee struct {
	Person
	Employee
	taxRate int
}

type PrintInfo interface {
	getMessage() string
}

func getMessage(p PrintInfo) {
	fmt.Println(p.getMessage())
}

func main() {

	ftEmployee := FullTimeEmployee{}
	ftEmployee.Person.age = 21
	ftEmployee.name = "Nombre"
	ftEmployee.id = 1

	fmt.Println(ftEmployee)

	tEmployee := TemporaryEmployee{}
	getMessage(tEmployee)
	getMessage(ftEmployee)

}
