package main

import (
	"fmt"

	"github.com/sneicast/golang-algoritmos/algoritmos/fibonacci"
)

func main() {
	menu :=
		`
Seleccione una de las siguientes opciones:

	  MENU
[ 1 ] Fibonacci

`
	fmt.Print(menu)

	var eleccion int
	fmt.Scanln(&eleccion)

	fmt.Println("****************************************************************")

	switch eleccion {
	case 1:
		fibonacci.Fibonacci()
	default:
		fmt.Println("Opcion no disponible")
	}

}
