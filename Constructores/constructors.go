package main

import "fmt"

type Employee2 struct {
	id       int
	name     string
	vacation bool
}

func NewEmployee(id int, name string, vacation bool) *Employee2 {
	return &Employee2{
		id:       id,
		name:     name,
		vacation: vacation,
	}
}

func main() {

	// 1
	e := Employee2{}
	fmt.Printf("%v\n", e)

	// 2
	e2 := Employee2{id: 2, name: "nombre"}
	fmt.Printf("%v\n", e2)

	//3
	e3 := new(Employee2) // devuelve un apuntador de la instancia
	// de employee2
	fmt.Printf("%v\n", *e3)

	//4
	e4 := NewEmployee(12, "Nombre4", true)
	fmt.Println(*e4)

}
