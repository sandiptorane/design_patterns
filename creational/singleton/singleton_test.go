package singleton

import (
	"testing"
)

func TestGetInstance(t *testing.T) {
	counter1 := GetInstance()

	if counter1 == nil{
		//Test of acceptance criteria  failed When no counter has been created before, a new one is created with the value 0go
		t.Error("expected pointer to Singleton after calling GetInstance(), not nil")
	}

	expectedCounter := counter1

	currentCount := counter1.AddOne()

	if currentCount !=1{
		t.Errorf("After calling for the first time to count, the count must be\n1 but it is %d\n",currentCount)
	}

	counter2 := GetInstance()

	if counter2 != expectedCounter{
		//Test2 failed
		t.Errorf("Expected same instance in counter2 but it got different instance\n")
	}

	currentCount = counter2.AddOne()

	if currentCount !=2{
		t.Errorf("After calling AddOne using second counter, currentCount must be 2 but it is %d\n",currentCount)
	}

}
