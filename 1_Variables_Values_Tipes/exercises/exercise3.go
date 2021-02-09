package main

import "fmt"

var x int = 42
var y string = "Jack"
var z bool = true

func main() {
	s := fmt.Sprintf("%v\t%v\t%v\t" , x, y, z) // \t is tab
	fmt.Println(s)
}