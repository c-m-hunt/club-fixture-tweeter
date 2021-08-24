package playcricket

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type ClubMatches []ClubMatch

type Team struct {
	ClubID      string
	ClubName    string
	ClubTwitter string
	TeamID      string
	TeamName    string
}

type ClubMatch struct {
	Match
	Team
	Opposition Team
	Venue      string
}

type Match struct {
	ID              int    `json:"id"`
	Status          string `json:"status"`
	Published       string `json:"published"`
	LastUpdated     string `json:"last_updated"`
	LeagueName      string `json:"league_name"`
	LeagueID        string `json:"league_id"`
	CompetitionName string `json:"competition_name"`
	CompetitionID   string `json:"competition_id"`
	CompetitionType string `json:"competition_type"`
	MatchType       string `json:"match_type"`
	GameType        string `json:"game_type"`
	Season          string `json:"season"`
	MatchDate       string `json:"match_date"`
	MatchTime       string `json:"match_time"`
	GroundName      string `json:"ground_name"`
	GroundID        string `json:"ground_id"`
	GroundLatitude  string `json:"ground_latitude"`
	GroundLongitude string `json:"ground_longitude"`
	HomeClubName    string `json:"home_club_name"`
	HomeTeamName    string `json:"home_team_name"`
	HomeTeamID      string `json:"home_team_id"`
	HomeClubID      string `json:"home_club_id"`
	AwayClubName    string `json:"away_club_name"`
	AwayTeamName    string `json:"away_team_name"`
	AwayTeamID      string `json:"away_team_id"`
	AwayClubID      string `json:"away_club_id"`
	Umpire1Name     string `json:"umpire_1_name"`
	Umpire1ID       string `json:"umpire_1_id"`
	Umpire2Name     string `json:"umpire_2_name"`
	Umpire2ID       string `json:"umpire_2_id"`
	Umpire3Name     string `json:"umpire_3_name"`
	Umpire3ID       string `json:"umpire_3_id"`
	RefereeName     string `json:"referee_name"`
	RefereeID       string `json:"referee_id"`
	Scorer1Name     string `json:"scorer_1_name"`
	Scorer1ID       string `json:"scorer_1_id"`
	Scorer2Name     string `json:"scorer_2_name"`
	Scorer2ID       string `json:"scorer_2_id"`
}

type MatchesResponse struct {
	Matches []Match `json:"matches"`
}

func (c Client) GetFixtures(season string) []ClubMatch {
	url := fmt.Sprintf("http://play-cricket.com/api/v2/matches.json?site_id=%v&season=%v&api_token=%v", c.SiteID, season, c.APIToken)

	spaceClient := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	matches := MatchesResponse{}
	jsonErr := json.Unmarshal(body, &matches)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	cms := []ClubMatch{}
	for _, m := range matches.Matches {
		cm := ClubMatch{Match: m}
		ht := Team{ClubID: m.HomeClubID, ClubName: m.HomeClubName, TeamID: m.HomeTeamID, TeamName: m.HomeTeamName}
		at := Team{ClubID: m.AwayClubID, ClubName: m.AwayClubName, TeamID: m.AwayTeamID, TeamName: m.AwayTeamName}
		if m.HomeClubID == c.SiteID {
			cm.Team = ht
			cm.Opposition = at
			cm.Venue = "home"
		} else {
			cm.Team = at
			cm.Opposition = ht
			cm.Venue = "away"
		}
		cms = append(cms, cm)
	}

	return cms
}

func (cms ClubMatches) FilterByDate(d time.Time, teams []string) ClubMatches {
	fixFiltered := []ClubMatch{}
	date := d.Format("02/01/2006")

	for _, f := range cms {
		if f.Match.MatchDate == date && contains(teams, f.Team.TeamID) {
			fixFiltered = append(fixFiltered, f)
		}
	}
	fmt.Printf("Filtered from %v fixtures to %v for date %v\n", len(cms), len(fixFiltered), date)
	return fixFiltered
}

func (cms *ClubMatches) PopulateTwitter(twitterMap map[string]string) {
	cmsOut := ClubMatches{}
	for _, cm := range *cms {
		if val, ok := twitterMap[cm.Opposition.ClubID]; ok {
			cm.Opposition.ClubTwitter = fmt.Sprintf("@%v", val)
		}
		cmsOut = append(cmsOut, cm)
	}
	*cms = cmsOut
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
