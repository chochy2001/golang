package main

import "fmt"

type Person struct {
	name string
	age  uint8
}

type Employee struct {
	id int
}

// FullTimeEmployee Se pone el tipo de manera anonima
// para acceder de manera directa a las propiedades
type FullTimeEmployee struct {
	Person
	Employee
}

func GetMessage(p Person) {
	fmt.Printf("%s with age %d\n", p.name, p.age)
}

func main() {

	ftEmployee := FullTimeEmployee{}
	ftEmployee.Person.age = 21
	ftEmployee.name = "Nombre"
	ftEmployee.id = 1

	fmt.Println(ftEmployee)

	//No se puede hacer por que no extiende,
	//ya que un fulltime employee no es una persona

	//GetMessage(ftEmployee)

}
