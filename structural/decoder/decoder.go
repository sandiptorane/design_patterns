package decoder

import (
	"errors"
	"fmt"
)

type IngredientAdd interface {
	AddIngredient() (string, error)
}

type PizzaDecorator struct {
	Ingredient IngredientAdd
}

func (p *PizzaDecorator)AddIngredient() (string, error){
	return "The Pizza with the following ingredients:",nil
}

type Meat struct {
	Ingredient IngredientAdd
}

func (m *Meat)AddIngredient() (string, error){
	if m.Ingredient==nil{
		return "",errors.New("an IngredientAdd should be in the Ingredient field of Meat")
	}

	s,err := m.Ingredient.AddIngredient()
	if err!=nil{
		return "", err
	}
	word := "meat"
	return fmt.Sprintf("%s %s,",s,word),nil
}

type Onion struct {
	Ingredient IngredientAdd
}

func (o *Onion)AddIngredient() (string, error){
	if o.Ingredient==nil{
		return "",errors.New("an IngredientAdd should be in the Ingredient field of onion")
	}

	s,err := o.Ingredient.AddIngredient()
	if err!=nil{
		return "", err
	}
	word := "onion"
	return fmt.Sprintf("%s %s,",s,word),nil
}





