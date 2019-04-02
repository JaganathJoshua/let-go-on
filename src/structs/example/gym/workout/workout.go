package main

import (
	"fmt"
	"structs/example/gym/equipements"
)

func main() {
	fmt.Println("---Starting Workout Session---")
	// picking a dumbell of selective weight
	d := equipements.Dumbell{
		22.5,
	}
	// doing the excersice
	d.BicepCurl(15)
}
