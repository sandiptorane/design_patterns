package flyweight

import "time"

type Team struct {
	ID  int
	Name string
	Shield []byte
	Players []Player
	HistoricalData []HistoricalData
}

const (
	TEAM_A = iota   //The iota value starts to reset to 0 when we declare TEAM_A , so TEAM_A is equal to 0. On the TEAM_B variable, iota is incremented by one so TEAM_B is equal to 1.
	TEAM_B
)

type Player struct {
	Name string
	Surname string
	PreviousTeam string
	Photo []byte
}

type HistoricalData struct{
	Year uint8
	LeagueResults []Match
}

type Match struct {
	Date time.Time
	VisitorID uint64
	LocalID uint64
	LocalScore byte
	VisitorScore byte
	LocalShoots uint16
	VisitorShoots uint16
}

type TeamFlyweightFactory struct {
	CreatedTeams map[int]*Team
}

func NewTeamFactory() TeamFlyweightFactory {
	return TeamFlyweightFactory{
		CreatedTeams: make(map[int]*Team),
	}
}

func (t *TeamFlyweightFactory)GetTeam(teamId int) *Team{
	if t.CreatedTeams[teamId]!=nil{
		return t.CreatedTeams[teamId]
	}
	team := getTeamFactory(teamId)
	t.CreatedTeams[teamId]= &team
	return t.CreatedTeams[teamId]
}

func (t *TeamFlyweightFactory)GetNumOfObjects() int{
	return len(t.CreatedTeams)
}

func getTeamFactory(team int) Team{
	switch team {
	case TEAM_B:
		return Team{
		ID :2,
		Name: "TEAM_B",
		}
	default:
		return Team{
		ID: 1,
		Name: "TEAM_A",
		}
	}
}
