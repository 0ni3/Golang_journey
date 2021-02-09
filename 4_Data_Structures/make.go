package main

import "fmt"

func main(){

	x := make([]int, 10, 12)
	x[0] = 23
	x[9] = 111
	x = append(x, 1212)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
}