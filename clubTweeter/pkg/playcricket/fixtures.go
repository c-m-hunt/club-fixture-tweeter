package playcricket

import (
	"encoding/json"
	"fmt"
	"log"
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

const baseUrl = "http://play-cricket.com/api/v2/"

func NewClubMatch(m Match, cId string) ClubMatch {
	cm := ClubMatch{Match: m}
	ht := Team{ClubID: m.HomeClubID, ClubName: m.HomeClubName, TeamID: m.HomeTeamID, TeamName: m.HomeTeamName}
	at := Team{ClubID: m.AwayClubID, ClubName: m.AwayClubName, TeamID: m.AwayTeamID, TeamName: m.AwayTeamName}
	if m.HomeClubID == cId {
		cm.Team = ht
		cm.Opposition = at
		cm.Venue = "home"
	} else {
		cm.Team = at
		cm.Opposition = ht
		cm.Venue = "away"
	}
	return cm
}

func (c Client) GetFixtures(season string) (ClubMatches, error) {
	url := fmt.Sprintf("%smatches.json?site_id=%v&season=%v&api_token=%v", baseUrl, c.SiteID, season, c.APIToken)

	body, err:= c.getData(url)
	if err != nil {
		return nil, err
	}

	matches := MatchesResponse{}
	jsonErr := json.Unmarshal(body, &matches)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	cms := []ClubMatch{}
	for _, m := range matches.Matches {
		cm := NewClubMatch(m, c.SiteID)
		cms = append(cms, cm)
	}

	return cms, nil
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

func (cms ClubMatches) PopulateTwitter(twitterMap map[string]string) {
	for i, cm := range cms {
		if val, ok := twitterMap[cm.Opposition.ClubID]; ok {
			cms[i].Opposition.ClubTwitter = fmt.Sprintf("@%v", val)
		}
	}
}


func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}


func (cm *ClubMatches) FilterByStatus(status string) ClubMatches {
	filtered := []ClubMatch{}
	for _, f := range *cm {
		if f.Match.Status == status {
			filtered = append(filtered, f)
		}
	}
	return filtered
}