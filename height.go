package main

import "fmt"

func main()  {
	fmt.Print("Input height in feet: ")
	var feet float64
	fmt.Scanf("%f", &feet)
	meter := feet * 0.3048

	fmt.Println(meter)
}