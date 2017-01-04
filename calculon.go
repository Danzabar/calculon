package main

import (
    "github.com/Danzabar/calculon/slack"
    "strings"
)

// Constant for calculons name
const name = "calculon"

var actions map[string]func(m slack.Message, c *slack.SlackClient)

// On init (this is like a construct)
func init() {
    // Create the actions map
    actions = make(map[string]func(m slack.Message, c *slack.SlackClient))

    // Add actions
    actions["hello"] = Greeting
    actions["pull requests"] = OpenPullRequests
}

// Responds to a message if we deem it nessecary
func respond(m slack.Message, c *slack.SlackClient) {

    // We only want to support messages
    if m.Type != "message" {
        return
    }

    m.Text = strings.ToLower(m.Text)

    // We only want to respond if calculon is mentioned
    if !mentioned(m.Text) {
        return
    }

    // At this point calculon was mentioned, so check the keywords
    for k, v := range(actions) {
        if strings.Contains(m.Text, k) {
            v(m, c)
        }
    }
}

// Checks if calculon is mentioned and returns a bool
func mentioned(text string) bool {
    return strings.Contains(text, name) || strings.Contains(text, "<@u3krkndc5>")
}