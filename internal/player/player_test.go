package player

import (
    "testing"
)

func TestAddPlayer(t *testing.T) {
    name := "Cristiano Ronaldo"
    age := 36
    id := AddPlayer(name, age)
    player, err := FindPlayerByID(id)
    if err != nil {
        t.Fatalf("expected player to be found, got error: %v", err)
    }
    if player.Name != name || player.Age != age {
        t.Errorf("expected player %s, age %d; got %s, age %d", name, age, player.Name, player.Age)
    }
}

func TestGetPlayers(t *testing.T) {
    AddPlayer("Lionel Messi", 34)
    AddPlayer("Neymar Jr", 29)
    players := GetPlayers()
    if len(players) != 3 { // Assuming 1 player added in previous test
        t.Errorf("expected 3 players, got %d", len(players))
    }
}