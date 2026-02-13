// action: create
package main

import (
	"fmt"
	"log"

	"github.com/ImDevinC/go-fifa/internal/team"
)

func main() {
	tm := team.NewTeamManager()

	// Add teams
	t1 := tm.AddTeam("Team Alpha")
	t2 := tm.AddTeam("Team Beta")
	fmt.Println("Added teams:", t1, t2)

	// List teams
	teams := tm.ListTeams()
	fmt.Println("List of teams:", teams)

	// Remove a team
	err := tm.RemoveTeam(t1.ID)
	if err != nil {
		log.Fatalf("Error removing team: %v", err)
	}
	fmt.Println("After removal, list of teams:", tm.ListTeams())
}