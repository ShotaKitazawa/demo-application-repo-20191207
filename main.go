package main

import (
	"fmt"
	"os"

	"github.com/nlopes/slack"
)

func main() {
	// Get Slack BotAppID & Token from EnvironmentVariable
	botAppID := os.Getenv("SLACK_BOT_APP_ID")
	if botAppID == "" {
		panic("SLACK_BOT_APP_ID is not defined")
	}
	token := os.Getenv("SLACK_TOKEN")
	if token == "" {
		panic("SLACK_TOKEN is not defined")
	}

	// Slack Client & RTM
	client := slack.New(token)
	rtm := client.NewRTM()
	go rtm.ManageConnection()

	// getBotID
	botID, err := getBotID(client, botAppID)
	if err != nil {
		panic(err)
	}

	// Handle slack events
	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.MessageEvent:
			if ev.Msg.BotID != botID {
				client.PostMessage(ev.Channel, slack.MsgOptionText(ev.Msg.Text, false))
				client.PostMessage(ev.Channel, slack.MsgOptionText(ev.Msg.Text, false))
			}
		}
	}
}

func getBotID(client *slack.Client, botAppID string) (botID string, err error) {
	users, err := client.GetUsers()
	if err != nil {
		return
	}
	for _, user := range users {
		if user.Profile.ApiAppID == botAppID {
			return user.Profile.BotID, nil
		}
	}
	return "", fmt.Errorf("None of Bot Users")
}
