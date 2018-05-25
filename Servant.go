package main

import (
	"log"
	"os"
	"strings"

	"github.com/nlopes/slack"
)

// ToBotID -- bot ID
const ToBotID = "<@BOTID>"

// SlackAPIToken -- slack API token
const SlackAPIToken = "<Slack API Token>"

var api = slack.New(SlackAPIToken)
var rtm = api.NewRTM()

func run(api *slack.Client) int {
	go rtm.ManageConnection()

	for {
		select {
		case msg := <-rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.HelloEvent:
				log.Print("Slack BOT Initialized.")

			case *slack.MessageEvent:
				log.Printf("Message: %v\n", ev)
				if strings.Contains(ev.Msg.Text, "/music") {
					ProcessSearchAppleMusicEvent(ev)
				}

			case *slack.InvalidAuthEvent:
				log.Print("Invalid credentials")
				return 1
			}
		}
	}
}

func main() {
	os.Exit(run(api))
}
