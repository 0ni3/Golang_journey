package main

import "fmt"

func main(){
	x := 42
	if x == 40 {
		fmt.Println("our values is 40")
	} else if x == 41 {
		fmt.Println("out value is 41")
	} else {
		fmt.Println("our value is not 40")
	}
}