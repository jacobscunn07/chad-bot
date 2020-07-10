package main

import (
	"github.com/go-joe/joe"
  "github.com/nlopes/slack"
)

type ChadBot struct {
  *joe.Bot
  Slack *slack.Client
}
