package twitter

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/c-m-hunt/club-tweeter/pkg/config"
	pc "github.com/c-m-hunt/club-tweeter/pkg/playcricket"
)

type tweetData struct {
	Fixtures string
}

func GenerateFixtureTweet(ms []pc.ClubMatch, cfg config.Config) string {
	fixturesPart := ""
	tweetTemplate, err := template.New("tweet").Parse(cfg.Templates.Fixtures)
	if err != nil {
		panic(err)
	}
	fixTemplate, err := template.New("fixture").Parse(cfg.Templates.FixtureLine)
	if err != nil {
		panic(err)
	}

	for _, t := range cfg.Teams {
		for _, f := range ms {
			if t == f.Team.TeamID {
				var fixLineBuff bytes.Buffer
				err = fixTemplate.Execute(&fixLineBuff, f)
				if err != nil {
					panic(err)
				}
				fixturesPart += fmt.Sprintf("%v\n", fixLineBuff.String())
			}
		}
	}

	var tweetBuff bytes.Buffer
	err = tweetTemplate.Execute(&tweetBuff, tweetData{fixturesPart})
	if err != nil {
		panic(err)
	}
	return tweetBuff.String()
}
