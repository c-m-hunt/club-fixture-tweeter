package app

import (
	_ "embed"
	"fmt"
	"time"

	"github.com/c-m-hunt/club-tweeter/pkg/config"
	pc "github.com/c-m-hunt/club-tweeter/pkg/playcricket"
	"github.com/c-m-hunt/club-tweeter/pkg/twitter"
)

type ClubMatches []pc.ClubMatch

func PostTweet() {
	cfg := config.NewConfig()

	c := pc.NewClient(cfg.PlayCricket.ClubID, cfg.PlayCricket.APIToken)
	fixs := c.GetFixtures(pc.GetCurrentSeason())

	for i := 0; i < 8; i++ {
		date := time.Now().Add(time.Duration(int(time.Hour) * 24 * i))
		fixFiltered := ClubMatches(fixs).FilterByDate(date, cfg.PlayCricket.Teams)
		if len(fixFiltered) > 0 {
			tweetStr := twitter.GenerateFixtureTweet(fixFiltered, cfg)
			fmt.Printf("----------------------\n%v\n----------------------\n", tweetStr)
			if i == 0 {
				twitter.SendTweet(cfg.TwitterAuth, tweetStr)
			}
		}
	}
}

func (cms ClubMatches) FilterByDate(d time.Time, teams []string) ClubMatches {
	fixFiltered := []pc.ClubMatch{}
	date := d.Format("02/01/2006")

	for _, f := range cms {
		if f.Match.MatchDate == date && contains(teams, f.Team.TeamID) {
			fixFiltered = append(fixFiltered, f)
		}
	}
	fmt.Printf("Filtered from %v fixtures to %v for date %v\n", len(cms), len(fixFiltered), date)
	return fixFiltered
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
