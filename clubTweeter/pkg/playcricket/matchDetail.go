package playcricket

import (
	"encoding/json"
	"fmt"
	"log"
)

type MatchDetailsResponse struct {
	MatchDetails []MatchDetail  `json:"match_details"`
}

type MatchDetail struct {
	Match
	TossWonByTeamID   string        `json:"toss_won_by_team_id"`
	Toss              string        `json:"toss"`
	BattedFirst       string        `json:"batted_first"`
	NoOfOvers         string        `json:"no_of_overs"`
	BallsPerInnings   string        `json:"balls_per_innings"`
	NoOfInnings       string        `json:"no_of_innings"`
	NoOfDays          string        `json:"no_of_days"`
	NoOfPlayers       string        `json:"no_of_players"`
	NoOfReserves      string        `json:"no_of_reserves"`
	Result            string        `json:"result"`
	ResultDescription string        `json:"result_description"`
	ResultAppliedTo   string        `json:"result_applied_to"`
	MatchNotes        string        `json:"match_notes"`
	Points            []interface{} `json:"points"`
	MatchResultTypes  [][]string    `json:"match_result_types"`
	Players           []struct {
		HomeTeam []struct {
			Position     int    `json:"position"`
			PlayerName   string `json:"player_name"`
			PlayerID     int    `json:"player_id"`
			Captain      bool   `json:"captain"`
			WicketKeeper bool   `json:"wicket_keeper"`
		} `json:"home_team,omitempty"`
		AwayTeam []struct {
			Position     int    `json:"position"`
			PlayerName   string `json:"player_name"`
			PlayerID     int    `json:"player_id"`
			Captain      bool   `json:"captain"`
			WicketKeeper bool   `json:"wicket_keeper"`
		} `json:"away_team,omitempty"`
	} `json:"players"`
	Innings []struct {
		TeamBattingName                    string `json:"team_batting_name"`
		TeamBattingID                      string `json:"team_batting_id"`
		InningsNumber                      int    `json:"innings_number"`
		ExtraByes                          string `json:"extra_byes"`
		ExtraLegByes                       string `json:"extra_leg_byes"`
		ExtraWides                         string `json:"extra_wides"`
		ExtraNoBalls                       string `json:"extra_no_balls"`
		ExtraPenaltyRuns                   string `json:"extra_penalty_runs"`
		PenaltiesRunsAwardedInOtherInnings string `json:"penalties_runs_awarded_in_other_innings"`
		TotalExtras                        string `json:"total_extras"`
		Runs                               string `json:"runs"`
		Wickets                            string `json:"wickets"`
		Overs                              string `json:"overs"`
		Balls                              string `json:"balls"`
		Declared                           bool   `json:"declared"`
		ForfeitedInnings                   bool   `json:"forfeited_innings"`
		RevisedTargetRuns                  string `json:"revised_target_runs"`
		RevisedTargetOvers                 string `json:"revised_target_overs"`
		RevisedTargetBalls                 string `json:"revised_target_balls"`
		Bat                                []struct {
			Position    string `json:"position"`
			BatsmanName string `json:"batsman_name"`
			BatsmanID   string `json:"batsman_id"`
			HowOut      string `json:"how_out"`
			FielderName string `json:"fielder_name"`
			FielderID   string `json:"fielder_id"`
			BowlerName  string `json:"bowler_name"`
			BowlerID    string `json:"bowler_id"`
			Runs        string `json:"runs"`
			Fours       string `json:"fours"`
			Sixes       string `json:"sixes"`
			Balls       string `json:"balls"`
		} `json:"bat"`
		Fow []struct {
			Runs           string `json:"runs"`
			Wickets        int    `json:"wickets"`
			BatsmanOutName string `json:"batsman_out_name"`
			BatsmanOutID   string `json:"batsman_out_id"`
			BatsmanInName  string `json:"batsman_in_name"`
			BatsmanInID    string `json:"batsman_in_id"`
			BatsmanInRuns  string `json:"batsman_in_runs"`
		} `json:"fow"`
		Bowl []struct {
			BowlerName string `json:"bowler_name"`
			BowlerID   string `json:"bowler_id"`
			Overs      string `json:"overs"`
			Maidens    string `json:"maidens"`
			Runs       string `json:"runs"`
			Wides      string `json:"wides"`
			Wickets    string `json:"wickets"`
			NoBalls    string `json:"no_balls"`
		} `json:"bowl"`
	} `json:"innings"`
}

func (c Client) GetMatchDetail() (*MatchDetail, error) {
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

// type MatchEvents struct {
// 	MatchId string
// 	Fifties []string
// 	Hundreds []string 
// 	FiveWickets []string 
// }

// func (cms ClubMatches) GetEvents(d time.Time, teams []string) []MatchEvents {
// 	fixFiltered := []ClubMatch{}
// 	date := d.Format("02/01/2006")

// 	for _, f := range cms {
// 		if f.Match.MatchDate == date && contains(teams, f.Team.TeamID) {
// 			fixFiltered = append(fixFiltered, f)
// 		}
// 	}
// 	fmt.Printf("Filtered from %v fixtures to %v for date %v\n", len(cms), len(fixFiltered), date)
// 	return fixFiltered
// }