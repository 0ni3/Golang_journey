package main

import "fmt"

func main(){
	x := []int{23, 32, 87, 93, 43, 55}
	fmt.Println(x, "\n")
	fmt.Println("the lenght of the array is equal to:", len(x))

	// fmt.Println(x)
	// fmt.Println(x[0])
	// fmt.Println(x[1])
	// fmt.Println(x[1])
	// fmt.Println(x[1])

	//this is index position

	for i, v := range x {
		fmt.Println(i, v)
	}
}	