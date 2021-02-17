package flyweight

import (
	"fmt"
	"testing"
)

func TestTeamFlyweightFactory_GetTeam(t *testing.T) {
	factory := NewTeamFactory()

	teamA1 := factory.GetTeam(TEAM_A)
	if teamA1==nil{
		t.Error("The pointer to TEAM_A was nil")
	}

	teamA2 := factory.GetTeam(TEAM_A)

	if teamA2 == nil{
		t.Error("The pointer to TEAM_A was nil")
	}

	if teamA1 != teamA2{
		t.Error("The pointers of TEAM_A weren't same")
	}

	if factory.GetNumOfObjects()!=1{
		t.Errorf("The number of object created was not 1 : %d",factory.GetNumOfObjects())
	}
}


//We are going to create a million calls to the team  creation, representing a million calls from users.
//Then, we will simply check that the number of teams created is only two
func Test_HighVolume(t *testing.T) {
	factory := NewTeamFactory()

	teams := make([]*Team,500000*2)

	for i:=0;i<500000;i++{
		teams[i] = factory.GetTeam(TEAM_A)
	}

	for i:=500000;i<500000*2;i++{
		teams[i] = factory.GetTeam(TEAM_B)
	}

	if factory.GetNumOfObjects() != 2{
		t.Errorf("The number of objects created was not 2: %d\n",factory.GetNumOfObjects())
	}

	for i:=0; i<3; i++ {
		fmt.Printf("Pointer %d points to %p and is located in %p\n", i, teams[i], &teams[i])
	}
}

