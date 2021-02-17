package main

import (
	"fmt"
)


type Athlete struct {

}

func (a *Athlete) Train() {
	fmt.Println("training")
}

type SwimmerA struct {
	MyAthlete Athlete
	MySwim func()
}

//The Swim() function takes no arguments and returns nothing, so it can be used
//as the MySwim field in SwimmerA struct:

func Swim() {
	fmt.Println("Swimming")
}

//for Animal
type Animal struct {

}

func (a *Animal)Eat(){
	fmt.Println("Eating")
}

type Shark struct {
	Animal
	MySwim func()
}

func main() {
	swimmer := SwimmerA{
		MySwim: Swim,
	}

	swimmer.MyAthlete.Train()
	swimmer.MySwim()

	fmt.Println("----for shark animal----")
	fish := Shark{
		MySwim : Swim,
	}

	fish.Eat()
	fish.MySwim()

}


