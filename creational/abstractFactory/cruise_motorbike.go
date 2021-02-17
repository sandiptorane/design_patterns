package abstractFactory

type CruiseMotorBike struct {

}

func (s *CruiseMotorBike) NumWheels() int{
	return 2
}

func (s *CruiseMotorBike) NumSeats() int{
	return 1
}

func (s *CruiseMotorBike) GetMotorBikeType() int{
	return CruiseMotorBikeType
}
