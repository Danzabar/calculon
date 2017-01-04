package main

import (
    "flag"
    "log"
    "github.com/Danzabar/calculon/slack"
    "github.com/Danzabar/calculon/bitbucket"
)

var BB *bitbucket.Client

func main() {
    // Flags
    token := flag.String("token", "", "The token for this bot")
    flag.Parse()

    s := slack.NewClient(*token)
    BB = bitbucket.NewClient()
    log.Print("Calculon is active baby!")

    for {

        m, e := s.GetMessage()

        // If there is no error from get message
        // we respond, else we move on, there is no
        // point throwing an error here
        if e == nil {
            respond(m, s)
        }
    }
}
