package bitbucket

import (
    "encoding/json"
    "net/http"
    "strings"
)

const API_BASE_URL = "https://bitbucket.org/api/2.0"

// BB Client struct
type Client struct {
    User string
    Pass string
}

// Creates and returns a new Client pointer
func NewClient(user string, pass string) *Client {
    return &Client {
        User: user,
        Pass: pass,
    }
}

// Executes a REST API request
func (c *Client) Execute(method string, url string, body string, out interface{}) error {
    io := strings.NewReader(body)
    req, _ := http.NewRequest(method, API_BASE_URL + url, io)

    req.Header.Set("Content-Type", "application/json")
    req.SetBasicAuth(c.User, c.Pass)

    resp, err := http.DefaultClient.Do(req)

    if err != nil {
        return  err
    }

    err = json.NewDecoder(resp.Body).Decode(out)

    if err != nil {
        return err
    }

    return nil
}