package abstractFactory

import (
	"errors"
	"fmt"
)

type VehicleFactory interface {
	Build(v int) (Vehicle,error)
}

const (
	CarFactoryType = 1
	MotorBikeFactoryType = 2
)

func BuildFactory(f int) (VehicleFactory, error) {
	switch f {
	case CarFactoryType:
		return new(CarFactory),nil
	case MotorBikeFactoryType:
		return new(MotorBikeFactory),nil
	default:
		return nil, errors.New(fmt.Sprintf("Factory with id %d not recognized\n", f))
	}
}

const(
	LuxuryCarType = 1
	FamilyCarType = 2
)

type CarFactory struct {}

func (c *CarFactory) Build(v int) (Vehicle,error){
	switch v {
	case LuxuryCarType:
		return new(LuxuryCar),nil
	case FamilyCarType:
		return new(FamilyCar),nil
	default:
		return nil,errors.New(fmt.Sprintf("Vehicle of type %d not recognized",v))
	}
}

const(
	SportMotorBikeType = 1
	CruiseMotorBikeType = 2
)

type MotorBikeFactory struct {
}

func (m *MotorBikeFactory)Build(v int) (Vehicle,error){
	switch v {
	case SportMotorBikeType:
		return new(SportMotorBike),nil
	case CruiseMotorBikeType:
		return new(CruiseMotorBike),nil
	default:
		return nil,errors.New(fmt.Sprintf("Vehicle of type %d not recognized",v))
	}
}
