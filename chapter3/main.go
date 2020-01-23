package main

import "fmt"

func main() {
	fmt.Println("1 + 1 =", 1+1)
	fmt.Println(len("Hello World"))
	fmt.Println("Hello World"[1]) /* this will return 101 byte representation of e) */
	fmt.Println("Hello " + "World")
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
	fmt.Println((true && false) || (false && true) || !(false && false))
}
