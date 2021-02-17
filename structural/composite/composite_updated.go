package main

import "fmt"

type swimmer interface {
	swim()
}

type trainer interface {
	train()
}

type swimmerImpl struct{}

func (s *swimmerImpl) swim(){
	println("Swimming!")
}
type compositeSwimmerB struct{
	trainer
	swimmer
}

type athlete struct {

}

func (a *athlete) train() {
	fmt.Println("training")
}

func main(){
	swimmer:= compositeSwimmerB{
		&athlete{},
		&swimmerImpl{},
	}

	swimmer.train()
	swimmer.swim()
}