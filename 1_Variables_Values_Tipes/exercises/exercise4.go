package main

import "fmt"

var y int
var x newtype

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	var x int = 42
	fmt.Println(x)
}