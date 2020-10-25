package main

import "fmt"

func main(){
	fmt.Print("Enter a Number: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := 2 * input
	fmt.Println(output)
}
