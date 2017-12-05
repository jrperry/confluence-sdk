package confluence

import (
	"time"
)

type Content struct {
	client     *Client
	ID         string         `json:"id"`
	Type       string         `json:"type"`
	Status     string         `json:"current"`
	Title      string         `json:"title"`
	Space      Space          `json:"space"`
	Body       ContentBody    `json:"body"`
	History    ContentHistory `json:"history"`
	Version    ContentVersion `json:"version"`
	Links      Links          `json:"_links"`
	Expandable Expandable     `json:"_expandable"`
}

type ContentBody struct {
	View ContentView `json:"view"`
}

type ContentView struct {
	Value string `json:"value"`
}

type ContentHistory struct {
	Latest      bool      `json:"latest"`
	CreatedBy   User      `json:"user"`
	CreatedDate time.Time `json:"created_date"`
}

type ContentVersion struct {
	Number    int       `json:"number"`
	MinorEdit bool      `json:"minorEdit"`
	Hidden    bool      `json:"hidden"`
	By        User      `json:"by"`
	When      time.Time `json:"when"`
	Message   string    `json:"message"`
}
