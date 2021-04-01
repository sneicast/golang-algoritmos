package fibonacci

import "fmt"

func Fibonacci() {
	fmt.Println("¿Cuantos numeros de la serie Fibonacci quiere que se visualicen?")
	var maximumValue int = 0
	fmt.Scanln(&maximumValue)
	fmt.Println("*****************Serie Fibonacci********************************")

	var previousNumber int64 = 0
	var fibonacciNumber int64 = 1
	for i := 0; i < maximumValue; i++ {
		fmt.Println(fibonacciNumber)
		var tempNumber int64 = previousNumber + fibonacciNumber
		previousNumber = fibonacciNumber
		fibonacciNumber = tempNumber
	}
}
