package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/nlopes/slack"
)

// TODO
const channelID = "C02EFSW4Y"

type SlackAttachment struct {
	Pretext string
	Text    string
}

type Input struct {
	Text string `json:"text"`
}

type SendMessage interface {
	Send(arg string) string
}

// Send -- execute shell command
func Send(slackAttachment SlackAttachment, message string, channelID string) string {
	api := slack.New(os.Getenv("SLACK_TOKEN"))
	params := slack.PostMessageParameters{AsUser: true}
	attachment := slack.Attachment{
		Pretext: slackAttachment.Pretext,
		Text:    slackAttachment.Text,
	}
	params.Attachments = []slack.Attachment{attachment}
	channelName, timestamp, err := api.PostMessage(channelID, message, params)
	if err != nil {
		return err.Error()
	}
	return fmt.Sprintf("Message successfully sent to channel %s at %s", channelName, timestamp)
}

func execute(text string) {
	command := string(text[5:])
	if strings.Contains(command, "/music") {
		ProcessSearchAppleMusicEvent(command, channelID)
	}
}

func request(ctx context.Context, input Input) {
	execute(input.Text)
	//log := send(SlackAttachment{Pretext: "Pretext", Text: "Dummy Pretext Text"}, fmt.Sprintf("Hello %s!!", text), "C02EFSW4Y")
}

func main() {
	lambda.Start(request)
}
