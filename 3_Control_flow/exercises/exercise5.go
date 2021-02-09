package main

import "fmt"

func main(){
	for i := 10; i < 15; i++ {
		fmt.Printf("when %v is divided by 4, the remainder is %v\n", i, i%4)
	}
}