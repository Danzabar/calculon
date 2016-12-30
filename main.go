package main

import (
    "flag"
    "log"
    "github.com/Danzabar/calculon/slack"
)

func main() {
    // Flags
    token := flag.String("token", "", "The token for this bot")
    flag.Parse()

    s := slack.NewClient(*token)

    for {

        m, e := s.GetMessage()

        if e != nil {
            log.Fatal(e)
        }

        log.Printf("%+v", m)
        respond(m, s)
    }
}
