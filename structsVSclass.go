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

func main() {
	e := Employee{}
	fmt.Printf("%v", e)
	e.id = 1
	e.name = "Nombre"
	fmt.Println(e)
	e.SetId(5)
	fmt.Println(e)

}
