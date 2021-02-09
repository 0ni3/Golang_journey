package main

import "fmt"

const (
	a = iota
	b = iota
)
func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
}