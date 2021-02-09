package main

import "fmt"

func main(){
	x := []int{3, 4, 2, 33, 76}
	fmt.Println(x[1:3])

	for i, v := range x {
		fmt.Println(i, v)
	}

	for i :=0; i <= 4; i++{
		fmt.Println(i, x[i])
	}
}