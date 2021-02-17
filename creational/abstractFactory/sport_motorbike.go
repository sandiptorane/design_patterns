package abstractFactory

type SportMotorBike struct {

}

func (s *SportMotorBike) NumWheels() int{
	return 2
}

func (s *SportMotorBike) NumSeats() int{
	return 1
}

func (s *SportMotorBike) GetMotorBikeType() int{
	return SportMotorBikeType
}
