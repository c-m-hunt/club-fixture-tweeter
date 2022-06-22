package main

import (
	"bytes"
	"context"
	_ "embed"
	"encoding/json"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/c-m-hunt/club-tweeter/pkg/app"
)

type Response events.APIGatewayProxyResponse


func Handler(ctx context.Context) (Response, error) {
	var buf bytes.Buffer

	body, err := json.Marshal(map[string]interface{}{
		"message": "Run successful",
	})
	if err != nil {
		return Response{StatusCode: 404}, err
	}
	json.HTMLEscape(&buf, body)

	whichApp, exists := os.LookupEnv("APP")
	if exists {
		if whichApp == "FIX_TWEETER" {
			app.RunFixtureTweet(true)
		} else if whichApp == "SCORE_UPDATER" {
			app.RunFixtureTweet(true)
		}
	}

	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}

	return resp, nil
}

func main() {
	// lambda.Start(Handler)
	se := app.ScoreEvent{1, app.BatFifty}
	se.GenerateImage("imgs/out/result.jpg")
	app.RunScoreTweet()
}
