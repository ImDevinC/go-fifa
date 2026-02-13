package team

// Team represents a soccer team with a name and basic statistics.
type Team struct {
    Name         string
    Wins         int
    Draws        int
    Losses       int
}

// TeamStats returns total matches played and points for the team.
func (t *Team) TeamStats() (matchesPlayed int, points int) {
    matchesPlayed = t.Wins + t.Draws + t.Losses
    points = (t.Wins * 3) + t.Draws
    return
}

// NewTeam initializes a new team with the given name and statistics.
func NewTeam(name string, wins, draws, losses int) *Team {
    return &Team{
        Name:   name,
        Wins:   wins,
        Draws:  draws,
        Losses: losses,
    }
}

// To implement additional functionality such as storage, retrieval, etc., 
// consider integrating a database or external file storage.