Superfantasy

Data structure
// UserLeague is a group of UserTeams
type UserLeague struct {
	UserTeams []UserTeam
}

// UserTeam is is the teams that a user picks
type UserTeam struct {
	Teams []models.Team
}

// League has many teams
type League struct {
	Name string
}