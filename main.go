package main

import (
  "fmt"
  "github.com/go-joe/joe"
 joeSlack "github.com/go-joe/slack-adapter/v2"
 "github.com/nlopes/slack"
  "github.com/spf13/viper"
)

func main() {
  viper.SetConfigName("config")
  viper.AddConfigPath(".")
  viper.SetEnvPrefix("chad")
  viper.AutomaticEnv()
  viper.SetConfigType("yaml")

  if err := viper.ReadInConfig(); err != nil {
    fmt.Printf("Error reading config file, %s", err)
  }

  var configuration Configuration

  if err := viper.Unmarshal(&configuration); err != nil {
    fmt.Printf("Unable to decode into struct, %v", err)
  }

 slackToken := configuration.SlackToken
 b := &ChadBot{
   Bot: joe.New("chad", joeSlack.Adapter(slackToken)),
   Slack: slack.New(slackToken),
 }

 b.Respond("dog pile", b.DogPile)
 b.Respond("dog me", b.DogMe)
 b.RespondRegex("roll\\s(\\d+)?d(\\d+)((\\+|\\-)\\d+)?", b.RollDieRegex)

 if err := b.Run(); err != nil {
   b.Logger.Fatal(err.Error())
 }
}

type Configuration struct {
  SlackToken string `mapstructure:"slack_token"`
}
