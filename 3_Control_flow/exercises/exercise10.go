package main

import "fmt"

func main(){
	switch {
	case (true && true):
		fmt.Println("migliorare in go")
	case (true && false):
		fmt.Println("procrastinare")
	case (true || true):
		fmt.Println("migliorare in python")
	case (true || false):
		fmt.Println("rimandare")
	case (!true):
		fmt.Println("studiare pentesting.. daje!")
	}
}