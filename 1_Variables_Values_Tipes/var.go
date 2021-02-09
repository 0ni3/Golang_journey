package main

import "fmt"

var z = "====================="
var a int // is zero as a default
type newtype int
var b newtype


func main() {
	a = 42
	b := 100
	a = int(b) // this is a convertion, a it is converts to newtype
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(z)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	// fmt.Printf("%T\n", y)
	// fmt.Printf("%#x\n%b\n%x\n", y, y, y)


	second()
}
func second() {
	fmt.Println("fine:", a)
}