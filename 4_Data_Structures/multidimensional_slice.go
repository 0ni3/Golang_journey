package main

import "fmt"

func main(){

	x := []string{"James", "Bond"}
	fmt.Println(x)
	y := []string{"Miss", "Moneypenny", "peasy"}
	fmt.Println(y)

	//slicing string
	xy := [][]string{x, y}
	fmt.Println(xy)
}