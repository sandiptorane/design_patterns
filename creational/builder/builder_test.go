package creational

import "testing"

func TestBuilder(t *testing.T){
	manufacturingComplex := &ManufacturingDirector{}

	//Test for Build a car
	t.Run("Test for Build a car",func(t *testing.T) {
		carBuilder := &CarBuilder{}

		manufacturingComplex.SetBuilder(carBuilder)
		manufacturingComplex.Construct()

		car := carBuilder.GetVehicle()

		if car.Wheels != 4 {
			t.Errorf("wheels on car must be 4 and they were %d\n", car.Wheels)
		}

		if car.Structure != "car" {
			t.Errorf("Structure on car must be 'car'but was %s \n", car.Structure)
		}

		if car.Seats != 5 {
			t.Errorf("seats on car must be 5 and they were %d\n", car.Seats)
		}

	})

	//Test for Build a bike
	t.Run("Test for Build a bike",func(t *testing.T) {
		bikeBuilder := &BikeBuilder{}

		manufacturingComplex.SetBuilder(bikeBuilder)
		manufacturingComplex.Construct()

		bike := bikeBuilder.GetVehicle()

		if bike.Wheels != 2 {
			t.Errorf("wheels on bike must be 2 and they were %d\n", bike.Wheels)
		}

		if bike.Structure != "bike" {
			t.Errorf("Structure on bike must be 'bike' but was %s \n", bike.Structure)
		}

		if bike.Seats != 2 {
			t.Errorf("seats on bike must be 2 and they were %d\n", bike.Seats)
		}
	})

}