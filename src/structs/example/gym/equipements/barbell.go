package equipements

import "fmt"

// Barbell Equipement Structure
type Barbell struct {
	Length       float32
	Weight       float32
	LoadedWeight float32
}

// Function to add given weights to the barbell
func (b *Barbell) AddWeights(weight float32) {
	b.LoadedWeight = weight
	fmt.Printf("Added weight %f to Barbell-%v \n", b.LoadedWeight, b)
}

// Function to do bicep curl with barbell
func (b *Barbell) BicepCurl(reps int) {
	fmt.Printf("Excercise Completed - %s %fft :: %s :: %fkg :: %dreps\n", "Barbell", b.Length, "BicepCurl", getTotalWeight(b), reps)
}

// Function for BenchPress using barbell
func (b *Barbell) BenchPress(reps int) {
	fmt.Printf("Excercise Completed - %s %fft :: %s :: %fkg :: %dreps\n", "Barbell", b.Length, "BenchPress", getTotalWeight(b), reps)
}

// Func to calculate the total weight of loaded barbell
func getTotalWeight(b *Barbell) float32 {
	return b.Weight + b.LoadedWeight
}
