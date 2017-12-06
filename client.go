package confluence

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func NewClient(hostname, username, password string) (*Client, error) {
	client := Client{
		Hostname: hostname,
		username: username,
		password: password,
	}
	var err error
	client.User, err = client.GetCurrentUser()
	if err != nil {
		return &client, err
	}
	return &client, nil
}

type Client struct {
	Hostname string
	username string
	password string
	User     User
}

func (c *Client) Get(relPath string) ([]byte, error) {
	client := http.Client{}
	url := fmt.Sprintf("%s/rest/api%s", c.Hostname, relPath)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []byte{}, err
	}
	req.SetBasicAuth(c.username, c.password)
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	data, _ := ioutil.ReadAll(resp.Body)
	return data, nil
}

func (c *Client) GetCurrentUser() (User, error) {
	user := User{}
	data, err := c.Get("/user/current")
	if err != nil {
		return user, err
	}
	err = json.Unmarshal(data, &user)
	return user, err
}

func (c *Client) GetUserByUsername(username string) (User, error) {
	user := User{}
	data, err := c.Get(fmt.Sprintf("/user?username=%s", username))
	if err != nil {
		return user, err
	}
	err = json.Unmarshal(data, &user)
	return user, err
}

func (c *Client) GetUserByKey(key string) (User, error) {
	user := User{}
	data, err := c.Get(fmt.Sprintf("/user?key=%s", key))
	if err != nil {
		return user, err
	}
	err = json.Unmarshal(data, &user)
	return user, err

}

func (c *Client) GetSpaces() []Space {
	spaces := []Space{}
	start := 0
	limit := 25
	url := fmt.Sprintf("/space?limit=%d&start=%d", limit, start)
	for {
		result := QueryResponse{}
		data, _ := c.Get(url)
		json.Unmarshal(data, &result)
		objects := []Space{}
		json.Unmarshal(result.Results, &objects)
		spaces = append(spaces, objects...)
		if result.Size < result.Limit {
			return spaces
		}
		start += limit
		url = fmt.Sprintf("/space?limit=%d&start=%d", limit, start)
	}
}

func (c *Client) GetSpace(key string) (Space, error) {
	space := Space{}
	data, err := c.Get(fmt.Sprintf("/space/%s", key))
	if err != nil {
		return space, err
	}
	err = json.Unmarshal(data, &space)
	space.client = c
	return space, err
}

func (c *Client) GetContent(id string) (Content, error) {
	content := Content{}
	data, err := c.Get(fmt.Sprintf("/content/%s", id))
	if err != nil {
		return content, err
	}
	err = json.Unmarshal(data, &content)
	content.client = c
	return content, err
}
