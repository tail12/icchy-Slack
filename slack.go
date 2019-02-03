package main

import (
	"fmt"
	"github.com/nlopes/slack"
	"log"
	"strings"
)

type slackListener struct {
	client    *slack.Client
	botID     string
	channelID string
}

func (s slackListener) ListenAndResponse() {
	rtm := s.client.NewRTM()

	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.HelloEvent:
			if err := s.handleHelloEvent(ev); err != nil {
				log.Printf("error: %s", err)
			}
		case *slack.MessageEvent:
			if err := s.handleMessageEvent(ev); err != nil {
				log.Printf("error: %s", err)
			}
		}
	}
}

func (s *slackListener) handleHelloEvent(ev *slack.HelloEvent) error {
	log.Printf("info: Bot is online")

	return nil
}

func (s *slackListener) handleMessageEvent(ev *slack.MessageEvent) error {
	if !strings.HasPrefix(ev.Msg.Text, fmt.Sprintf("<@%s>", s.botID)) {
		return nil
	}
	a := slack.Attachment{
		Title: "ご用件は何でしょうか",
		Color: "good",
	}
	_, _, err := s.client.PostMessage(s.channelID, slack.MsgOptionAttachments(a))
	return err
}
