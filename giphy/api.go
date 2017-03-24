package giphy

import (
    "encoding/json"
    "net/http"
)

type Result struct {
    Data Giphy `json:"data"`
}

func (r *Result) Deserialize(t *http.Response) error {
    return json.NewDecoder(t.Body).Decode(r)
}

type Giphy struct {
    Type   string `json:"type"`
    URL    string `json:"url"`
    Rating string `json:"rating"`
}
