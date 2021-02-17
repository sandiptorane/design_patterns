package creational

type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats()  BuildProcess
	SetStructure() BuildProcess
	GetVehicle() VehicleProduct
}

type ManufacturingDirector struct {
	Builder BuildProcess
}

func (m *ManufacturingDirector) SetBuilder(b BuildProcess){
	m.Builder =b
}

func (m *ManufacturingDirector) Construct(){
     m.Builder.SetSeats()
     m.Builder.SetWheels()
     m.Builder.SetStructure()
}

type VehicleProduct struct {
	Wheels int
	Seats int
	Structure string
}

//car builder
type CarBuilder struct {
	v VehicleProduct
}

func (c *CarBuilder) SetWheels() BuildProcess{
	c.v.Wheels =4
	return c
}

func (c *CarBuilder) SetSeats() BuildProcess{
	c.v.Seats = 5
	return c
}

func (c *CarBuilder) SetStructure() BuildProcess{
	c.v.Structure = "car"
	return c
}

func (c *CarBuilder) GetVehicle() VehicleProduct{
	return c.v
}

//Bike builder
type BikeBuilder struct {
	v VehicleProduct
}

func (b *BikeBuilder) SetWheels() BuildProcess{
	b.v.Wheels = 2
	return b
}

func (b *BikeBuilder) SetSeats() BuildProcess{
	b.v.Seats = 2
	return b
}

func (b *BikeBuilder) SetStructure() BuildProcess{
	b.v.Structure = "bike"
	return b
}

func (b *BikeBuilder) GetVehicle() VehicleProduct{
	return b.v
}