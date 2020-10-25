package main

import "fmt"

func main() {
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i = 1 + i
	}
	actualFor()
}

func actualFor()  {
	for i := 0; i < 10; i++{
		fmt.Print(i)
	}
}