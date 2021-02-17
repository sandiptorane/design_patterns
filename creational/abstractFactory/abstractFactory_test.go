package abstractFactory

import "testing"

func TestMotorBikeFactory(t *testing.T) {
	motorBikeF, err := BuildFactory(MotorBikeFactoryType)    //build motorbike factory
	if err!=nil{
		t.Fatal(err)
	}

	motorBikeVehicle, err := motorBikeF.Build(SportMotorBikeType)  //build Sport bike
	if err!=nil{
		t.Fatal(err)
	}
	t.Logf("Motorbike vehicle has %d wheels\n", motorBikeVehicle.NumWheels())

    sportmotor, ok := motorBikeVehicle.(Motorbike)  //that the type we have received is correct.
    if !ok{
    	t.Fatal("Struct assertion failed")
	}
	t.Logf("Sport motorbike has type %d\n",sportmotor.GetMotorBikeType())
}

//Test for car factory
func TestCarFactory(t *testing.T) {
	carF,err := BuildFactory(CarFactoryType)   //build car factory
	if err!=nil{
		t.Fatal(err)
	}

	carVehicle,err := carF.Build(LuxuryCarType)  //build car of type luxury car
	if err!=nil{
		t.Fatal(err)
	}
	t.Logf("Car vehicle has %d wheels\n",carVehicle.NumWheels())

	luxurycar, ok := carVehicle.(Car)  //check type using assertion
	if !ok{
		t.Fatal("struct assertion failed")
	}
	t.Logf("luuxry car has %d doors\n",luxurycar.NumDoors())
}
