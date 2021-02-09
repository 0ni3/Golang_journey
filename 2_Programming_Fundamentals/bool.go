package main

import "fmt"

var x bool //zero value is false

func main() {
	a := 7
	b := 42
	fmt.Println(a <= b)
	fmt.Println(x)
	x = true
	fmt.Println(x)
}