package interfaces

type Equipements interface {
	ManualEquipements
	MachineEquipements
	AddWeights(weight float64)
}

type ManualEquipements interface {
	BicepCurl(reps int)
}

type MachineEquipements interface {
	LegPress(reps int)
}
