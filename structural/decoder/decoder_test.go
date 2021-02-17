package decoder

import (
	"strings"
	"testing"
)

func TestPizzaDecorator_AddIngredient(t *testing.T) {
	t.Run("AddIngredient method of the pizza decorator object",func(t *testing.T){
		pizza := &PizzaDecorator{}

		pizzaResult, _ := pizza.AddIngredient()
		expectedResult := "The Pizza with the following ingredients:"

		assertResult(t, pizzaResult, expectedResult)

	})
}

func TestOnion_AddIngredient(t *testing.T) {
	onion := &Onion{}

	onionResult,err :=onion.AddIngredient()

	if err==nil{
		t.Errorf("When calling AddIngredient on the onion decorator without an IngredientAdd on its Ingredient field must return an error, not a string with '%s'", onionResult)
	}

    t.Run("calling AddIngredient on the onion decorator ",func(t *testing.T) {
		onion = &Onion{&PizzaDecorator{}}
		onionResult, err = onion.AddIngredient()
		if err != nil {
			t.Error(err)
		}
		expectedResult := "onion"
		assertResult(t, onionResult, expectedResult)
	})
}

func TestMeat_AddIngredient(t *testing.T) {
	meat := &Meat{}

	meatResult,err :=meat.AddIngredient()

	if err==nil{
		t.Errorf("When calling AddIngredient on the meat decorator without an IngredientAdd on its Ingredient field must return an error, not a string with %s", meatResult)
	}

	t.Run("calling AddIngredient on the onion decorator ",func(t *testing.T) {
		meat = &Meat{&PizzaDecorator{}}
		meatResult, err = meat.AddIngredient()
		if err != nil {
			t.Error(err)
		}
		expectedResult := "meat"
		assertResult(t, meatResult, expectedResult)
	})
}

func TestPizzaDecorator_FullStack(t *testing.T) {
	t.Run("asking for a pizza with onion and meat",func(t *testing.T){
		pizza := &Onion{&Meat{&PizzaDecorator{}}}

		pizzaResult,err := pizza.AddIngredient()
		if err != nil {
			t.Error(err)
		}
		expectedText := "The Pizza with the following ingredients: meat, onion,"

		assertResult(t,pizzaResult,expectedText)
		//t.Log(pizzaResult)
	})
}

func assertResult(t *testing.T,actual,expected string){
	t.Helper()
	if !strings.Contains(actual,expected){
		t.Errorf("returned string must contains text/word: '%s' \n but got: '%s' \n",expected,actual)
	}
}




