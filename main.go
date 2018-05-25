package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/nlopes/slack"
)

type SlackAttachment struct {
	Pretext string
	Text    string
}

const SlackBotToken = "<Slack Bot Token>"

// Send -- execute shell command
func send(slackAttachment SlackAttachment, message string, channelID string) string {
	api := slack.New(SlackBotToken)
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

func request() (string, error) {
	log := send(SlackAttachment{Pretext: "Pretext", Text: "Dummy Pretext Text"}, "This is a test.", "<Channel ID>")
	return log, nil
}

func main() {
	lambda.Start(request)
}
