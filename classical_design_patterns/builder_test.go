package main

import (
	"fmt"
	"testing"
)
// BUILDER PATTERN
// The construction process is always abstracted from the user of the product.
// Builder resolves the telescoping constructor anti-pattern
// Builder is handy when the parameters are boolean, as the builder result would be addA(), addB(), addC() etc to the entity
// The key difference from the factory pattern is that factory pattern is to be used when
// the creation is a one step process while builder pattern is to be used when the creation is a multi step process.

// Trade-off
// The BuildProcess interface specifies what he must comply to be part of the possible builders (good for self documentation)
// Very useful for stable and predictable algorithm, as any small change in this interface will affect
// all your builders and it could be awkward if you add a new method that some of your
// builders need and others Builders do not




// could be a singleton
type ManufacturingDirector struct{
	builder BuildProcess
}
func (m *ManufacturingDirector) Construct(){
	m.builder.SetSeats().SetStructure().SetWheels()
}
func (m *ManufacturingDirector) SetBuilder(b BuildProcess){
	m.builder = b
}

type BuildProcess interface{
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	GetVehicle() VehicleProduct
}

// BASE ENTITY VEHICLE
type VehicleProduct struct{
	Wheels int
	Seats int
	Structure string
}

// CAR IMPLEMENTATION
type CarBuilder struct {
	v VehicleProduct
}
func (c *CarBuilder) SetWheels() BuildProcess {
	c.v.Wheels = 4
	return c
}

func (c *CarBuilder) SetSeats() BuildProcess {
 	c.v.Seats = 5
 	return c
}
func (c *CarBuilder) SetStructure() BuildProcess {
	c.v.Structure = "Car"
	return  c
}
func (c *CarBuilder) GetVehicle() VehicleProduct {
	return c.v
}
// BIKE IMPLEMENTATION
type BikeBuilder struct {
	v VehicleProduct
}
func (b *BikeBuilder) SetWheels() BuildProcess {
	b.v.Wheels = 2
	return b
}
func (b *BikeBuilder) SetSeats() BuildProcess {
	b.v.Seats = 2
	return b
}
func (b *BikeBuilder) SetStructure() BuildProcess {
 	b.v.Structure = "Motorbike"
 	return b
}
func (b *BikeBuilder) GetVehicle() VehicleProduct {
	return b.v
}

// BUS IMPLEMENTATION
type BusBuilder struct {
	v VehicleProduct
}
func (b *BusBuilder) SetWheels() BuildProcess {
	b.v.Wheels = 4*2
	return b
}
func (b *BusBuilder) SetSeats() BuildProcess {
	b.v.Seats = 30
	return b
}
func (b *BusBuilder) SetStructure() BuildProcess {
	b.v.Structure = "Bus"
	return b
}
func (b *BusBuilder) GetVehicle() VehicleProduct {
	return b.v
}


//go test -v -run=TestBuilder .
func TestBuilderPattern(t *testing.T) {
	manufacturingComplex := ManufacturingDirector{}

	// CAR TEST
	carBuilder := &CarBuilder{}
	manufacturingComplex.SetBuilder(carBuilder)
	manufacturingComplex.Construct()
	car := carBuilder.GetVehicle()

	fmt.Print(car)
	if car.Wheels != 4 {
		t.Errorf("Wheels on a car must be 4 and they were %d\n", car.Wheels)
	}
	if car.Structure != "Car" {
		t.Errorf("Structure on a car must be 'Car' and was %s\n",
			car.Structure)
	}
	if car.Seats != 5 {
		t.Errorf("Seats on a car must be 5 and they were %d\n", car.Seats)
	}

	//BIKE TEST
	bikeBuilder := &BikeBuilder{}
	manufacturingComplex.SetBuilder(bikeBuilder)
	manufacturingComplex.Construct()
	motorbike := bikeBuilder.GetVehicle()
	motorbike.Seats = 1
	if motorbike.Wheels != 2 {
		t.Errorf("Wheels on a motorbike must be 2 and they were %d\n",
			motorbike.Wheels)
	}
	if motorbike.Structure != "Motorbike" {
		t.Errorf("Structure on a motorbike must be 'Motorbike' and was %s\n",
			motorbike.Structure)
	}

}

func main() {
	fmt.Print("OK")
}
