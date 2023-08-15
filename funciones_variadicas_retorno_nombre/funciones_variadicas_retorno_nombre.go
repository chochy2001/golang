package main

import "fmt"

// funcion normal
func sum(a, b int) int {
	return a + b
}

// funcion suma con parametros variables
func sum2(values ...int) int {
	total := 0
	for _, v := range values {
		total += v
	}
	return total
}

func printNames(names ...string) {
	for _, n := range names {
		fmt.Println(n)
	}
}

// devuelve 3 valores sumando 1 y 2 al valor pasado
func getValues(x int) (double int, triple int, quad int) {
	double = x * 2
	triple = x + 3
	quad = x * 4
	return
}

func main() {
	fmt.Println(sum(1, 2))
	fmt.Println(sum2(1, 2, 3, 4, 5))
	printNames("Nombre1", "Nombre2", "Nombre3")
	fmt.Println(getValues(2))
}
