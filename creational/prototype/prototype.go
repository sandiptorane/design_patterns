package prototype

import (
	"errors"
	"fmt"
)

type ShirtCloner interface{
	GetClone(s int) (ItemInfoGetter,error)
}

const(
	White = 1
	Black = 2
	Blue = 3
)

func GetShirtCloner() ShirtCloner{
	return &ShirtCache{}
}

type ShirtCache struct{}

func (s *ShirtCache)GetClone(i int) (ItemInfoGetter,error){
	switch i {
	case White:
		newItem := *whitePrototype
		return &newItem,nil
	case Black:
		newItem := *blackPrototype
		return &newItem,nil
	case Blue:
		newItem := *bluePrototype
		return &newItem,nil
	default:
		return nil,errors.New("Shirt model not recognized\n")
	}

}

type ItemInfoGetter interface {
	GetInfo() string
}

type ShirtColor byte

type Shirt struct {
	Price float64
	SKU string    //SKU - Stock keeping unit
	Color ShirtColor
}

func (s *Shirt) GetInfo() string{
	return fmt.Sprintf("shirt with SKU %s and color id %d that costs %f",s.SKU,s.Color,s.Price)
}

var whitePrototype = &Shirt{
	Price: 500,
	SKU: "empty",
	Color: White,
}

var blackPrototype = &Shirt{
	Price: 600,
	SKU: "empty",
	Color: Black,
}
var bluePrototype = &Shirt{
	Price: 700,
	SKU: "empty",
	Color: Blue,
}

func (s *Shirt)GetPrice() float64{
	return s.Price
}
