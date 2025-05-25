package go_fifa_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	fifa "github.com/imdevinc/go-fifa"
	"github.com/stretchr/testify/assert"
)

var expectedSearchPlayer = fifa.SearchPlayerResponse{
	Results: []fifa.PlayerResponse{{
		Id: "390267",
		Name: []fifa.LocaleDescription{{
			Locale:      "en-GB",
			Description: "Christian  PULISIC",
		}},
		Alias: []fifa.LocaleDescription{{
			Locale:      "en-GB",
			Description: "Christian PULISIC",
		}},
		Weight:                   72.0,
		Height:                   179.0,
		Birthplace:               "Pennsylvania",
		CountryId:                "USA",
		InternationalCaps:        56,
		InternationalDebut:       "2016-03-29T00:00:00Z",
		TopCompetitionDebut:      nil,
		PictureUrl:               "",
		ThumbnailUrl:             "",
		TwitterAccount:           nil,
		PreferredFoot:            float64(9999),
		MediaContent:             []any{},
		LocalizedTwitterAccounts: nil,
		Goals:                    22,
		PlayerPicture:            nil,
		Properties: fifa.Properties{
			StatsPerformIfesId: "390267",
			IdStatsPerform:     "93soe2rjn0z874cc9ofcesgkp",
			IdIFES:             "390267",
		},
		IsUpdateable: false,
	}},
}

func TestSearchPlayer(t *testing.T) {
	t.Parallel()
	client := fifa.Client{
		Client: &TestClient{},
	}
	resp, err := client.SearchPlayer(&fifa.SearchOptions{"Pulisic"})
	if ok := assert.Nil(t, err, "expected no error with SearchPlayer, got: %s", err); !ok {
		t.FailNow()
	}
	diff := cmp.Diff(*resp, expectedSearchPlayer)
	if ok := assert.Equal(t, diff, "", diff); !ok {
		t.Fail()
	}
}

var expectedSearchCompetition = fifa.SearchCompetitionResponse{
	PaginationResponse: fifa.PaginationResponse{
		ContinuationHash:  "50",
		ContinuationToken: "50",
	},
	Results: []fifa.Competition{{
		Id: "521",
		Name: []fifa.LocaleDescription{{
			Locale:      "en-GB",
			Description: "FIFA Women's World Cup Qualifier",
		}},
		ConfederationIds:     []string{},
		MemberAssociationIds: []string{},
		OwnerId:              "FIFA",
		Gender:               fifa.FemaleGender,
		FootballType:         0,
		TeamType:             1,
		Type:                 3,
		AgeType:              7,
		Properties: fifa.Properties{
			IdIFES: "FWWC_Q",
		},
		IsUpdateable: false,
	}},
}

func TestSearchCompetition(t *testing.T) {
	t.Parallel()
	client := fifa.Client{
		Client: &TestClient{},
	}
	resp, err := client.SearchCompetition(&fifa.SearchOptions{"FIFA World Cup"})
	if ok := assert.Nil(t, err, "expected no error with SearchCompetition, got: %s", err); !ok {
		t.FailNow()
	}
	diff := cmp.Diff(*resp, expectedSearchCompetition)
	if ok := assert.Equal(t, diff, "", diff); !ok {
		t.Fail()
	}
}

var expectedSearchTeam = fifa.SearchTeamResponse{
	PaginationResponse: fifa.PaginationResponse{
		ContinuationToken: "",
		ContinuationHash:  "",
	},
	Results: []fifa.TeamResponse{{
		Id:              "2000019635",
		ConfederationId: "CAF",
		Type:            0,
		AgeType:         7,
		FootballType:    0,
		Gender:          fifa.MaleGender,
		Name: []fifa.LocaleDescription{{
			Locale:      "en-GB",
			Description: "Paradou AC",
		}},
		AssociationId:  "ALG",
		CityId:         "",
		Headquarters:   nil,
		TrainingCenter: nil,
		OfficialSite:   nil,
		City:           "al-Jaza’ir (Algiers)",
		CountryId:      "ALG",
		PostalCode:     "",
		RegionName:     nil,
		ShortClubName:  "Paradou",
		Abbreviation:   "PAC",
		Street:         "",
		FoundationYear: 1994,
		Stadium:        nil,
		PictureUrl:     "https://api.fifa.com/api/v3/picture/teams-{format}-{size}/2000019635",
		DisplayName: []fifa.LocaleDescription{{
			Locale:      "en-gb",
			Description: "PAC",
		}},
		Content: []any{},
		Properties: fifa.Properties{
			IdStatsPerform: "9bvywjkxdag5aye9rvzedglat",
		},
		IsUpdateable: false,
	}},
}

func TestSearchTeam(t *testing.T) {
	t.Parallel()
	client := fifa.Client{
		Client: &TestClient{},
	}
	resp, err := client.SearchTeam(&fifa.SearchOptions{Name: "Paradou"})
	if ok := assert.Nil(t, err, "expected no error with SearchTeam, got: %s", err); !ok {
		t.FailNow()
	}
	diff := cmp.Diff(*resp, expectedSearchTeam)
	if ok := assert.Equal(t, diff, "", diff); !ok {
		t.Fail()
	}
}
