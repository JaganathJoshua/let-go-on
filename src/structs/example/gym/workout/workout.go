package main

import (
	"fmt"
	"structs/example/gym/equipements"
	"structs/example/gym/interfaces"
)

// Main function
func main() {
	fmt.Println("---Starting Workout Session---")
	// picking a dumbell of selective weight
	d := equipements.Dumbell{
		Weight: 22.5,
	}
	// perform excercises with this equipement
	performExercisesWithManualEquipements(&d, 15)
}

// Function to perfrom set of exercises using manual gym equipements
func performExercisesWithManualEquipements(equipment interfaces.ManualEquipements, reps int) {
	equipment.BicepCurl(reps)
}
