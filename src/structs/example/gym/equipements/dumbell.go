package equipements

import "fmt"

// Dumbell structure
type Dumbell struct {
	Weight float64
}

// Interface to define list of exercises
type Exercise interface {
	BicepCurl(reps int)
}

// Function for BicepCurl exercise
func (d *Dumbell) BicepCurl(reps int) {
	fmt.Printf("Excercise Completed - %s :: %fkg :: %dreps\n", "BicepCurl", d.Weight, reps)
}
