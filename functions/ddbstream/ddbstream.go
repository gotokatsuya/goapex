package main

import (
  "fmt"
  "os"

  "github.com/apex/go-apex"
  "github.com/apex/go-apex/dynamo"
)

func main() {
  dynamo.HandleFunc(func(event *dynamo.Event, ctx *apex.Context) error {
    for _, record := range event.Records {
      fmt.Fprintf(os.Stderr, "record: %v\n", record)
    }
    return nil
  })
}
