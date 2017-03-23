package main

import (
    "flag"
    "github.com/Danzabar/calculon/bitbucket"
    "github.com/Danzabar/calculon/giphy"
    "github.com/Danzabar/calculon/slack"
    "log"
)

var (
    BB  *bitbucket.Client
    GIF *giphy.Client
)

func main() {
    // Flags
    token := flag.String("token", "", "The token for this bot")
    flag.Parse()

    s := slack.NewClient(*token)
    BB = bitbucket.NewClient()
    GIF = giphy.NewClient()

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
