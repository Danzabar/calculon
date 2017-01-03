package main

import (
    "github.com/Danzabar/calculon/slack"
    "github.com/Danzabar/calculon/bitbucket"
    "math/rand"
    "fmt"
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

// Lists open pull requests from bitbucket repo
// @example `@calculon pull requests`
func OpenPullRequests(m slack.Message, c *slack.SlackClient) {
    resp := &bitbucket.PullRequest{}

    err := BB.Execute("GET", `/repositories/`+ BB.Owner +`/`+ BB.Repo +`/pullrequests/?q=state="OPEN"&pagelen=50`, "", resp)

    if err != nil {
        m.Text = "Amateurs! I was unable to obtain a usable response from your bucket of bits"
        c.PostMessage(m)
        return
    }

    if resp.Size == 0 {
        m.Text = "There are no open pull requests, you peasant"
        c.PostMessage(m)
        return 
    }

    m.Text = fmt.Sprintf("There are %d open pull requests!\n", resp.Size)

    for _, v := range(resp.Values) {
        m.Text += "```"
        m.Text += fmt.Sprintf("https://bitbucket.org/%s/%s/pull-requests/%d\n``` `%s` - It has %d comments and %d tasks\n", BB.Owner, BB.Repo, v.Id, v.Title, v.Comments, v.Tasks)
    }

    c.PostMessage(m)
}