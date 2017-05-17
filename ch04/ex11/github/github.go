package github

import "time"

// APIURL GitHub API URL
const APIURL = "https://api.github.com"

// Client GitHub Access Client
type Client struct {
	token string
}

// Issue GitHub Issue structure
type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url,omitempty"`
	Title     string `json:"title,omitempty"`
	State     string `json:"state,omitempty"`
	User      *User
	CreatedAt time.Time `json:"created_at,omitempty"`
	Body      string    `json:"body,omitempty"`
}

// User GitHub User structure
type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}
