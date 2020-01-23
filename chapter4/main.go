package main

import "fmt"

func main() {
/*	var x string
	x = "first"
	fmt.Println(x)
	x = "second"
	fmt.Println(x)


	var x string
	x = "first"
	fmt.Println(x)
	x += "second"
	fmt.Println(x)
*/
	
	x := "Hello World"
	fmt.Println(x)

	var (
		a = 5
		b = 10
		c = 15
	)
	
	fmt.Println(a,b,c)
	
	fmt.Println("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)
	
	output := input * 2

	fmt.Println(output)
}

