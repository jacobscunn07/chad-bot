package main

import (
  "fmt"
  "github.com/go-joe/joe"
  "math/rand"
  "strconv"
  "time"
)

func (b *ChadBot) RollDieRegex(msg joe.Message) error {
  qty, _ := strconv.Atoi(msg.Matches[0])
  faces, _ := strconv.Atoi(msg.Matches[1])
  modifier, _ := strconv.Atoi(msg.Matches[2])

  if qty <= 1 {
    qty = 1
  }

  seed := rand.NewSource(time.Now().UnixNano())
  random := rand.New(seed)

  for i := 0; i < qty; i++ {
    num := random.Intn(faces) + modifier + 1
    msg.Respond(fmt.Sprintf("You rolled %v", num))
  }

  return nil
}

