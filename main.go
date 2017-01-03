package main

import (
    "flag"
    "log"
    "github.com/Danzabar/calculon/slack"
    "github.com/ktrysmt/go-bitbucket"
    "os"
)

var BB *bitbucket.Client

func main() {
    // Flags
    token := flag.String("token", "", "The token for this bot")
    flag.Parse()

    s := slack.NewClient(*token)
    BB = bitbucket.NewBasicAuth(os.Getenv("BB_User"), os.Getenv("BB_Pass"))

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
