package player

import (
    "fmt"
    "errors"
)

// Player represents a soccer player
type Player struct {
    ID   int
    Name string
    Age  int
}

var players []Player
var nextID int

// AddPlayer adds a new player to the in-memory list
func AddPlayer(name string, age int) int {
    nextID++
    newPlayer := Player{
        ID:   nextID,
        Name: name,
        Age:  age,
    }
    players = append(players, newPlayer)
    return newPlayer.ID
}

// GetPlayers returns all players
func GetPlayers() []Player {
    return players
}

// FindPlayerByID finds a player by ID
func FindPlayerByID(id int) (Player, error) {
    for _, player := range players {
        if player.ID == id {
            return player, nil
        }
    }
    return Player{}, errors.New("player not found")
}