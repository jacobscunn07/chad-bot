package main

import (
  "github.com/go-joe/joe"
  "github.com/go-resty/resty/v2"
  "github.com/nlopes/slack"
  "strings"
)

func (b *ChadBot) DogMe(msg joe.Message) error {
  b.Slack.PostMessageContext(msg.Context, msg.Channel,
   slack.MsgOptionText(getDogPictureUrls(1)[0], false),
   slack.MsgOptionPostMessageParameters(
     slack.PostMessageParameters{
       UnfurlLinks: true,
       UnfurlMedia: true,
     },
   ),
  )

  return nil
}

func (b *ChadBot) DogPile(msg joe.Message) error {
  for _, url := range getDogPictureUrls(5) {
    b.Slack.PostMessageContext(msg.Context, msg.Channel,
      slack.MsgOptionText(url, false),
      slack.MsgOptionPostMessageParameters(
        slack.PostMessageParameters{
          UnfurlLinks: true,
          UnfurlMedia: true,
        },
      ),
    )
  }

  return nil
}

type Dog struct {
  Url string
}

func getDogPictureUrls(count int) []string {
  if count < 1 {
    count = 1
  }

  client := resty.New()
  urls := make([]string, 0)

  for i := 0; i < count; i++ {
    resp := &resty.Response{}
    url := ""
    for ok := true; ok; ok = strings.HasSuffix(strings.ToLower(url), ".mp4") {
      resp, _ = client.
        R().
        SetResult(&Dog{}).
        Get("https://random.dog/woof.json")
      url = resp.Result().(*Dog).Url
    }

    urls = append(urls, url)
  }

  return urls
}
