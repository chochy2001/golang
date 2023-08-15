package main

func main() {
	x := 5
	y := func() int {
		return x * 2
	}() //Sino se ponen los parametros se asigna la funcion
	// y con los parentesis se ejecuta la funcion y regresa el valor
	println(y)

	//Se usan cuando se hacen go routines

}
