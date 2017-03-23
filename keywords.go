package main

import (
    "github.com/Danzabar/calculon/slack"
    "math/rand"
)

// Reacts to the word `Strings`
// @examples `Wow hese are crazy strings`
func Strings(m slack.Message, c *slack.SlackClient) {
    r := [3]string {
        "Oh really? Go talk to Ultron, I digress.",
        "I used to be king here. Now I'm nothing but a mere peasant. Or at best, a viscount.",
        "Go talk to Ultron",
    }

    m.Text = r[rand.Intn(len(r))]
    c.PostMessage(m)
}

// Reacts to the word `Fifty-six`
// @example `Fifty-six`
func FiftySix(m slack.Message, c *slack.SlackClient) {
    m.Text = "Fifty-six? FIFTY-SIX?! NOW THAT'S ALL I CAN THINK ABOUT. I'M GONNA KILL YOU, YOU NO GOOD FIFTY-SIXER!"
    c.PostMessage(m)
}