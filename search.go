package go_fifa

type SearchOptions struct {
	Name string
}

// SearchTeam searches for a team based on the provided name
func (c *Client) SearchTeam(opts *SearchOptions) (*SearchTeamResponse, error) {
	var resp SearchTeamResponse
	_, err := c.get("/teams/search", &resp, opts)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SearchPlayer searches for a player based on the provided name
func (c *Client) SearchPlayer(opts *SearchOptions) (*SearchPlayerResponse, error) {
	var resp SearchPlayerResponse
	_, err := c.get("/players/search", &resp, opts)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SearchCompetition searches for a player based on the provided name
func (c *Client) SearchCompetition(opts *SearchOptions) (*SearchCompetitionResponse, error) {
	var resp SearchCompetitionResponse
	_, err := c.get("/competitions/search", &resp, opts)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
