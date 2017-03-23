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

func (g *Client) Search(q string) *Results {
    r := &Results{}
    resp, err := g.Request(fmt.Sprintf("gifs/search?q=%s", q))

    if err != nil {
        log.Fatal(err)
    }

    r.Deserialize(resp)
    return r
}

func (g *Client) Request(uri string) (*http.Response, error) {
    req, _ := http.NewRequest("GET", fmt.Sprintf("http://api.giphy.com/v1/%s&api_key=%s", uri, g.Token), nil)
    return http.DefaultClient.Do(req)
}
