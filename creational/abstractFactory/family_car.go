package abstractFactory

type FamilyCar struct{}

func (l *FamilyCar) NumWheels() int{
	return 4
}

func (l *FamilyCar) NumSeats() int{
	return 5
}

func (l *FamilyCar) NumDoors() int{
	return 5
}
