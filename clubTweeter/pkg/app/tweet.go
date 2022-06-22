package app

import (
	"fmt"
	"time"

	pc "github.com/c-m-hunt/club-tweeter/pkg/playcricket"
	"github.com/c-m-hunt/club-tweeter/pkg/twitter"
)

func RunFixtureTweet(sendTweet bool) error {

	c := pc.NewClient(cfg.PlayCricket.ClubID, cfg.PlayCricket.APIToken)
	season_fixtures, err := c.GetFixtures(pc.GetCurrentSeason())
	if (err != nil) {
		return err
	}
	fixs := pc.ClubMatches(season_fixtures)
	fixs.PopulateTwitter(cfg.PlayCricketTwitterMap)

	for i := 0; i < 8; i++ {
		date := time.Now().Add(time.Duration(int(time.Hour) * 24 * i))
		fixFiltered := pc.ClubMatches(fixs).FilterByDate(date, cfg.PlayCricket.Teams)

		if len(fixFiltered) > 0 {
			tweetStr := twitter.GenerateFixtureTweet(fixFiltered, cfg)
			for _, fix := range fixFiltered {
				fmt.Printf("%v - %v \n", fix.Opposition.ClubID, fix.Opposition.ClubName)
			}
			fmt.Printf("----------------------\n%v\n----------------------\n", tweetStr)
			if sendTweet {
				if i == 0 {
					twitter.SendTweet(cfg.TwitterAuth, tweetStr)
				}
			}
		}
	}
	return nil
}
