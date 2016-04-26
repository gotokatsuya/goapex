package main

import (
  "encoding/json"
  "strings"

  "github.com/apex/go-apex"

  "github.com/gotokatsuya/goapex/env"
  "github.com/gotokatsuya/goapex/helper/resource/mysql"
)

func init() {
  env.Load()
}

type Message struct {
  Value string `json:"value"`
}

func main() {
  apex.HandleFunc(func(event json.RawMessage, ctx *apex.Context) (interface{}, error) {

    // Ping database
    db := env.GetDB()
    if err := mysql.Ping(db.Name, db.Host, db.Port, db.User, db.Password); err != nil {
      return nil, err
    }

    // Uppercase
    var msg Message
    if err := json.Unmarshal(event, &msg); err != nil {
      return nil, err
    }
    return &Message{strings.ToUpper(msg.Value)}, nil
  })
}
