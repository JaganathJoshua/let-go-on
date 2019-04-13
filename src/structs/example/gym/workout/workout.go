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

	// picking a barbell of selective length and weight
	b := equipements.Barbell{
		Length: 7.2,
		Weight: 20.0,
	}
	// Adding more weights to barbell
	b.AddWeights(20.0)

	fmt.Println("\n Performing excercises with dumbell")
	performExercisesWithManualEquipements(&d, 15)
	fmt.Println("\n Performing excercises with barbell")
	performExercisesWithManualEquipements(&b, 15)
}

// Function to perfrom set of exercises using manual gym equipements
func performExercisesWithManualEquipements(equipment interfaces.ManualEquipements, reps int) {
	equipment.BicepCurl(reps)
	equipment.BenchPress(reps)
}
