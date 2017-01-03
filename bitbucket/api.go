package bitbucket

import (
    "time"
)

// The outer response struct from a pull request
type PullRequest struct {
    Page int `json:"page"`
    PageLength int `json:"pagelen"`
    Size int `json:"size"`
    Values []PullRequestValue `json:"values"`
}

// Represents a single pull request
type PullRequestValue struct {
    Author Author `json:"author"`
    Comments int `json:"comment_count"`
    Tasks int `json:"task_count"`
    Created time.Time `json:"created_on"`
    Updated time.Time `json:"updated_on"`
    Description string `json:"description"`
    Title string `json:"title"`
    Links map[string]map[string]string `json:"links"`
    Id uint `json:"id"`
}

// Respresents an Author
type Author struct {
    Display string `json:"display_name"`
    Name string `json:"username"`
}