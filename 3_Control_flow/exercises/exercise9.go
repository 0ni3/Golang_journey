package main

import "fmt"

func main(){
	favSport := "boxing"
	switch favSport {
		case "skiing":
			fmt.Println("go to the mountain")
		case "surfing":
			fmt.Println("go to the sea")
		case "boxing":
			fmt.Println("go to the ring")
	}
}