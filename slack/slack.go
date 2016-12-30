package slack

import (
    "encoding/json"
    "fmt"
    "golang.org/x/net/websocket"
    "io/ioutil"
    "net/http"
    "sync/atomic"
)

var counter uint64

type responseRtmStart struct {
    Ok    bool         `json:"ok"`
    Error string       `json:"error"`
    Url   string       `json:"url"`
    Self  responseSelf `json:"self"`
}

type responseSelf struct {
    Id string `json:"id"`
}

type Message struct {
    Id      uint64 `json:"id"`
    Type    string `json:"type"`
    Channel string `json:"channel"`
    Text    string `json:"text"`
}

type SlackClient struct {
    Client *websocket.Conn
    Token string
}

// Creates a new Client
func NewClient(token string) *SlackClient {
    s := &SlackClient{ Token: token }
    s.Connect()

    return s
}

// Connects to the web socket using the clients token
func (s *SlackClient) Connect() {
    var r responseRtmStart

    resp, err := http.Get(fmt.Sprintf("https://slack.com/api/rtm.start?token=%s", s.Token))

    if err != nil {
        panic(err)
    } else if resp.StatusCode != 200 {
        panic(fmt.Sprintf("API request failed with code %d", resp.StatusCode))
    }

    body, err := ioutil.ReadAll(resp.Body)
    resp.Body.Close()

    if err != nil {
        panic(err)
    }

    if err = json.Unmarshal(body, &r); err != nil {
        panic(err)
    } else if !r.Ok {
        panic(fmt.Sprintf("Slack error: %s", r.Error))
    }

    w, err := websocket.Dial(r.Url, "", "https://api.slack.com")

    if err != nil {
        panic(err)
    }

    s.Client = w
}

// Retrieves a message using the websocket connection
func (s *SlackClient) GetMessage() (m Message, err error) {
    err = websocket.JSON.Receive(s.Client, &m)
    return
}

// Sends a message using the websocket connection
func (s *SlackClient) PostMessage(m Message) error {
    m.Id = atomic.AddUint64(&counter, 1)
    return websocket.JSON.Send(s.Client, m)
}