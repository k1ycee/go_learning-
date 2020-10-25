package main

import "fmt"

func main(){
	fmt.Print("Input temprature in Fahrenheit: ")
	var fahr float64
	fmt.Scanf("%f", &fahr)

	celsius := (fahr - 32) * 5/9

	fmt.Println(celsius)
}