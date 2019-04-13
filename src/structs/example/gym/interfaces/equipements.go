package interfaces

type Equipements interface {
	ManualEquipements
	MachineEquipements
	AddWeights(weight float32)
}

type ManualEquipements interface {
	BicepCurl(reps int)
	BenchPress(reps int)
}

type MachineEquipements interface {
	LegPress(reps int)
}
