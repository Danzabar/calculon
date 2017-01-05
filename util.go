package main

import (
    "github.com/Danzabar/calculon/slack"
)

// Provides a default message for when bitbucket is unreachable
func defaultBBFailResponse(m slack.Message, c *slack.SlackClient) {
    m.Text = "Amateurs! I was unable to obtain a usable response from your bucket of bits"
    c.PostMessage(m)
}

// Provides a default response when no pull requests are found
func defaultBBNoPRResponse(m slack.Message, c *slack.SlackClient) {
    m.Text = "There are no pull requests, you peasant"
    c.PostMessage(m)
}