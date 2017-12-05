package confluence

import (
	"encoding/json"
	"fmt"
)

type Space struct {
	client     *Client
	ID         int        `json:"id"`
	Key        string     `json:"key"`
	Name       string     `json:"name"`
	Type       string     `json:"type"`
	Links      Links      `json:"_links"`
	Expandable Expandable `json:"_expandable"`
}

func (s *Space) GetContent() []Content {
	content := []Content{}
	start := 0
	limit := 25
	url := fmt.Sprintf("/space/%s/content/page?limit=%d&start=%d", s.Key, limit, start)
	for {
		result := QueryResponse{}
		data, _ := s.client.Get(url)
		json.Unmarshal(data, &result)
		objects := []Content{}
		json.Unmarshal(result.Results, &objects)
		for _, obj := range objects {
			obj.client = s.client
			content = append(content, obj)
		}
		if result.Size < result.Limit {
			return content
		}
		start += limit
		url = fmt.Sprintf("/space/%s/content/page?limit=%d&start=%d", s.Key, limit, start)
	}
}
