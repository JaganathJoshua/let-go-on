package equipements

import "fmt"

// Dumbell Equipement Structure
type Dumbell struct {
	Weight float32
}

// Function for BicepCurl exercise
func (d *Dumbell) BicepCurl(reps int) {
	fmt.Printf("Excercise Completed - %s :: %s :: %fkg :: %dreps\n", "Dumbell", "BicepCurl", d.Weight, reps)
}

// Func for BenchPress exercise usin dumbell
func (d *Dumbell) BenchPress(reps int) {
	fmt.Printf("Excercise Completed - %s :: %s :: %fkg :: %dreps\n", "Dumbell", "BenchPress", d.Weight, reps)
}
