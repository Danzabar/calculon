package main

import (
    "github.com/Danzabar/calculon/slack"
)

func Greeting(m slack.Message, c *slack.SlackClient) {
    m.Text = "I was all of history's great acting robots: Acting Unit 0.8, Thespo-mat, David Duchovny!"
    c.PostMessage(m)
}