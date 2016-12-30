package main

import (
    "github.com/Danzabar/calculon/slack"
    "math/rand"
)

// Selects a random greeting and sends it back as a message
// @example `hello calculon`
func Greeting(m slack.Message, c *slack.SlackClient) {
    g := [3]string {
        "I was all of history's great acting robots: Acting Unit 0.8, Thespo-mat, David Duchovny!",
        "Noooooooooo",
        "I'm a washed up ham",
    }

    m.Text = g[rand.Intn(len(g))]
    c.PostMessage(m)
}