// action: create
package team

import (
	"errors"
	"fmt"
)

// Team represents a participating team in the tournament
type Team struct {
	ID   int
	Name string
}

// TeamManager manages a list of teams
type TeamManager struct {
	teams []Team
	nextID int
}

// NewTeamManager creates a new instance of TeamManager
func NewTeamManager() *TeamManager {
	return &TeamManager{
		teams: []Team{},
		nextID: 1,
	}
}

// AddTeam adds a new team to the list
func (tm *TeamManager) AddTeam(name string) Team {
	team := Team{ID: tm.nextID, Name: name}
	tm.teams = append(tm.teams, team)
	tm.nextID++
	return team
}

// ListTeams returns all the teams
func (tm *TeamManager) ListTeams() []Team {
	return tm.teams
}

// RemoveTeam removes a team by ID
func (tm *TeamManager) RemoveTeam(id int) error {
	for i, t := range tm.teams {
		if t.ID == id {
			// Use slice tricks to remove item
			tm.teams = append(tm.teams[:i], tm.teams[i+1:]...)
			return nil
		}
	}
	return errors.New("team not found")
}