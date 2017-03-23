package giphy

import (
    "encoding/json"
    "net/http"
)

type Results struct {
    Data []SearchItem `json:"data"`
}

func (r *Results) Deserialize(t *http.Response) {
    json.NewDecoder(t.Body).Decode(r)
}

type SearchItem struct {
    Type   string `json:"type"`
    URL    string `json:"url"`
    Rating string `json:"rating"`
}
