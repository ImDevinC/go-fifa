package go_fifa

import "fmt"

type GetTeamOptions struct {
	TeamId string
}

func (c *Client) GetTeam(options *GetTeamOptions) (*TeamResponse, error) {
	var team TeamResponse
	url := fmt.Sprintf("/teams/%s", options.TeamId)
	_, err := c.get(url, &team, nil)
	if err != nil {
		return nil, err
	}
	return &team, nil
}
