package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/url"
)

type ConfigureResponse struct {
	Success bool   `json:"success"`
	Err     string `json:"error"`
}

type Cmd struct {
	Operation string   `json:"op"`
	Path      []string `json:"path"`
}

func (c *Client) Configure(cmds []Cmd) error {
	data, err := json.Marshal(cmds)
	if err != nil {
		return err
	}

	res, err := c.httpClient.PostForm(fmt.Sprintf("%s/configure", c.Options.Host), url.Values{
		"data": {string(data)},
		"key":  {c.Options.ApiKey},
	})
	if err != nil {
		return err
	}

	defer func() {
		if tempErr := res.Body.Close(); tempErr != nil {
			err = tempErr
		}
	}()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	r := ConfigureResponse{}
	err = json.Unmarshal(body, &r)
	if err != nil {
		return err
	}
	if r.Success {
		fmt.Println("✅ Changes applied successfully.")
	} else {
		fmt.Println("❌ Error applying changes.")
		return errors.New(r.Err)
	}

	return nil
}
