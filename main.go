package main

import (
 "github.com/go-joe/joe"
 joeSlack "github.com/go-joe/slack-adapter/v2"
 "github.com/nlopes/slack"
)

func main() {
 slackToken := "" //TODO: Get from a configuration source
 b := &ChadBot{
   Bot: joe.New("chad", joeSlack.Adapter(slackToken)),
   Slack: slack.New(slackToken),
 }

 b.Respond("dog pile", b.DogPile)
 b.Respond("dog me", b.DogMe)

 err := b.Run()
 if err != nil {
   b.Logger.Fatal(err.Error())
 }
}

