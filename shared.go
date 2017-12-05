package confluence

import (
	"encoding/json"
)

type Links struct {
	Base    string `json:"base"`
	Context string `json:"context"`
	Self    string `json:"self"`
	Next    string `json:"next"`
	WebUI   string `json:"webui"`
	Edit    string `json:"edit"`
}

type Expandable struct {
	Metadata    string `json:"metadata"`
	Icon        string `json:"icon"`
	Description string `json:"description"`
	Homepage    string `json:"homepage"`
	Container   string `json:"container"`
	Operations  string `json:"operations"`
	Children    string `json:"children"`
	Ancestors   string `json:"ancestors"`
	Descendants string `json:"descendants"`
	History     string `json:"history"`
	Body        string `json:"body"`
	Version     string `json:"version"`
	Space       string `json:"space"`
}

type QueryResponse struct {
	Results json.RawMessage `json:"results"`
	Start   int             `json:"start"`
	Limit   int             `json:"limit"`
	Size    int             `json:"size"`
	Links   Links           `json:"_links"`
}
