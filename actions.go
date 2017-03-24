package main

import (
    "fmt"
    "github.com/Danzabar/calculon/bitbucket"
    "github.com/Danzabar/calculon/slack"
    "math/rand"
)

// Selects a random greeting and sends it back as a message
// @example `hello calculon`
func Greeting(m slack.Message, c *slack.SlackClient) {
    g := [3]string{
        "I was all of history's great acting robots: Acting Unit 0.8, Thespo-mat, David Duchovny!",
        "Noooooooooo",
        "I'm a washed up ham",
    }

    m.Text = g[rand.Intn(len(g))]
    c.PostMessage(m)
}

// Lists action words
// @example `man calculon`
func Man(m slack.Message, c *slack.SlackClient) {
    m.Text = "So you want to know what I can do....\n *The following actions require my attention*\n```"

    for k := range actions {
        m.Text += fmt.Sprintf("%s\n", k)
    }

    m.Text += "```\n *These actions just require a keyword*\n ```"

    for k := range keywords {
        m.Text += fmt.Sprintf("%s\n", k)
    }

    m.Text += "```"

    c.PostMessage(m)
}

// Lists open pull requests from bitbucket repo
// @example `calculon pull requests`
func OpenPullRequests(m slack.Message, c *slack.SlackClient) {
    resp := &bitbucket.PullRequest{}

    err := BB.Execute("GET", `/repositories/`+BB.Owner+`/`+BB.Repo+`/pullrequests/?q=state="OPEN"&pagelen=50`, "", resp)

    if err != nil {
        defaultBBFailResponse(m, c)
        return
    }

    if resp.Size == 0 {
        defaultBBNoPRResponse(m, c)
        return
    }

    m.Text = fmt.Sprintf("There are %d open pull requests!\n", resp.Size)

    for _, v := range resp.Values {
        m.Text += "```"
        m.Text += fmt.Sprintf("https://bitbucket.org/%s/%s/pull-requests/%d\n``` `%s` - It has %d comments and %d tasks\n", BB.Owner, BB.Repo, v.Id, v.Title, v.Comments, v.Tasks)
    }

    c.PostMessage(m)
}

// Returns the user name of the last person who merged
// @example `calculon who broke it?`
func WhoBrokeIt(m slack.Message, c *slack.SlackClient) {
    resp := &bitbucket.PullRequest{}

    err := BB.Execute("GET", fmt.Sprintf(`/repositories/%s/%s/pullrequests?q=state="MERGED"&pagelen=1`, BB.Owner, BB.Repo), "", resp)

    if err != nil {
        defaultBBFailResponse(m, c)
        return
    }

    if resp.Size == 0 {
        defaultBBNoPRResponse(m, c)
        return
    }

    m.Text = "I am but a simple actor, but... I think.. "

    for _, v := range resp.Values {
        m.Text += fmt.Sprintf("%s broke everything! The Shame.", v.Author.Display)
    }

    c.PostMessage(m)
}

// Returns a random gif from giffy, or a message if its R-rated
// @example `calculon random gif`
func RandomGif(m slack.Message, c *slack.SlackClient) {
    r := GIF.Random()

    if r.Data.Rating == "r" {
        m.Text = "I'd love to show you this gif, alas its too filthy for work hours"
    } else {
        m.Text = "As requested.... " + r.Data.URL
    }

    c.PostMessage(m)
}
