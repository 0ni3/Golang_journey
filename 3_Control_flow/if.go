package main

import "fmt"

func main(){
	if true {
		fmt.Println("hello")
	}

	if false {
		fmt.Println("goodbye")
	}

	if 2 == 2 {
		fmt.Println("hello there")
	}

	if 2 != 2 {
		fmt.Println("hey")
	}

	if !(2 == 2){
		fmt.Println("mmm..")
	}

	if !(2 != 2){
		fmt.Println("ok")
	}
}