package main

import "fmt"

func main(){
	x := []int{4, 7, 5 , 43}
	fmt.Println(x)
	x = append(x, 23, 33, 55, 44)
	fmt.Println(x)

	y := []int{23, 24, 54, 53}
	x = append(x, y...)
	fmt.Println(x)
}