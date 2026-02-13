package team

import (
    "testing"
)

func TestTeamStats(t *testing.T) {
    team := NewTeam("FIFA Stars", 10, 5, 3)

    matchesPlayed, points := team.TeamStats()

    if matchesPlayed != 18 {
        t.Errorf("expected matchesPlayed to be 18, got %d", matchesPlayed)
    }

    if points != 35 {
        t.Errorf("expected points to be 35, got %d", points)
    }
}

func TestNewTeam(t *testing.T) {
    team := NewTeam("Soccer Warriors", 8, 6, 4)
    
    if team.Name != "Soccer Warriors" {
        t.Errorf("expected team name to be 'Soccer Warriors', got '%s'", team.Name)
    }

    if team.Wins != 8 || team.Draws != 6 || team.Losses != 4 {
        t.Errorf("expected stats to be wins: 8, draws: 6, losses: 4, got wins: %d, draws: %d, losses: %d",
			team.Wins, team.Draws, team.Losses)
    }
}