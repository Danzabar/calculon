package giphy

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

type Client struct {
    Token string
}

func NewClient() *Client {
    return &Client{
        Token: os.Getenv("GIPHY_TOKEN"),
    }
}

func (g *Client) Random() *Result {
    r := &Result{}
    r.Deserialize(g.Request("gifs/random"))
    return r
}

func (g *Client) Request(uri string) *http.Response {
    req, _ := http.NewRequest("GET", fmt.Sprintf("http://api.giphy.com/v1/%s?api_key=%s", uri, g.Token), nil)
    resp, err := http.DefaultClient.Do(req)

    if err != nil {
        log.Fatal(err)
    }

    return resp
}
