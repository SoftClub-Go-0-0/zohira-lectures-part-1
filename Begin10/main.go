package main

import "fmt"

func main() {
	var firstNumber float64
	var secoundNumber float64
	fmt.Scan(&firstNumber)
	fmt.Scan(&secoundNumber)
	var sum = (firstNumber * firstNumber) + (secoundNumber * secoundNumber)
	fmt.Println("sum is", sum)
	var difference = (firstNumber * firstNumber) - (secoundNumber * secoundNumber)
	fmt.Println("difference is", difference)
	var product = (firstNumber * firstNumber) * (secoundNumber * secoundNumber)
	fmt.Println("product is", product)
	var quotientOfFirstNumberSquares = (firstNumber * firstNumber) / (secoundNumber * secoundNumber)
	fmt.Println("quotientOfNumberSquarse is", quotientOfFirstNumberSquares)

}
