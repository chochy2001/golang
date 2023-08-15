package main

import "fmt"

type Employee struct {
	id   int
	name string
}

// SetId Receiver functions
func (e *Employee) SetId(id int) {
	e.id = id
}
func (e *Employee) SetName(name string) {
	e.name = name
}

// GetId hay que especificar el valor de retorno
func (e *Employee) GetId() int {
	return e.id
}

func (e *Employee) GetName() string {
	return e.name
}

func main() {
	e := Employee{}
	fmt.Printf("%v", e)
	e.id = 1
	e.name = "Nombre"
	fmt.Println(e)
	e.SetId(5)
	e.SetName("Nombre 2")
	fmt.Println(e)
	fmt.Println(e.GetId())
	fmt.Println(e.GetName())

}
