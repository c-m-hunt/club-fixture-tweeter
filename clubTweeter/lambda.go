package main

import (
	"bytes"
	"context"
	_ "embed"
	"encoding/json"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
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

	app.RunFixtureTweet()

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

func SQSHandler(ctx context.Context, sqsEvent events.SQSEvent) (Response, error) {

}

func main() {
	handler, exists := os.LookupEnv("HANDLER")
	if exists {
		if handler == "SQS" {
			lambda.Start(SQSHandler())
		} else {
			lambda.Start(Handler)
		}
	} 
	lambda.Start(Handler)
}
