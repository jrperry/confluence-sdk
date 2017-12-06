package confluence

import (
	"encoding/json"
	"fmt"
	"strings"
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
	View       ContentView `json:"view"`
	StyledView ContentView `json:"styled_view"`
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

type Attachment struct {
	ID         string     `json:"id"`
	Type       string     `json:"type"`
	Status     string     `json:"status"`
	Title      string     `json:"title"`
	Extensions Extensions `json:"extensions"`
	Links      Links      `json:"_links"`
}

func (c *Content) GetBody() (string, error) {
	content := Content{}
	data, err := c.client.Get(fmt.Sprintf("/content/%s?expand=body.styled_view.value", c.ID))
	if err != nil {
		return "", err
	}
	err = json.Unmarshal(data, &content)
	if err != nil {
		return "", err
	}
	contentBody := content.Body.StyledView.Value
	attachments := c.GetAttachments()
	for _, attachment := range attachments {
		attachmentURL := strings.Replace(attachment.Links.Download, "&", "&amp;", -1)
		newURL := fmt.Sprintf("%s%s", c.client.Hostname, attachmentURL)
		contentBody = strings.Replace(contentBody, attachmentURL, newURL, -1)
	}
	return contentBody, nil
}

func (c *Content) GetAttachments() []Attachment {
	attachments := []Attachment{}
	start := 0
	limit := 25
	url := fmt.Sprintf("/content/%s/child/attachment?limit=%d&start=%d", c.ID, limit, start)
	for {
		result := QueryResponse{}
		data, _ := c.client.Get(url)
		json.Unmarshal(data, &result)
		objects := []Attachment{}
		json.Unmarshal(result.Results, &objects)
		attachments = append(attachments, objects...)
		if result.Size < result.Limit {
			return attachments
		}
		start += limit
		url = fmt.Sprintf("/content/%s/child/attachment?limit=%d&start=%d", c.ID, limit, start)
	}
}
