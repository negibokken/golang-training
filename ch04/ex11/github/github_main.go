package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
)

func getStringByEditor() (string, error) {
	dir := os.TempDir()
	f, err := ioutil.TempFile(dir, "tmp")
	if f == nil || err != nil {
		return "", fmt.Errorf("File io error")
	}

	defer func() {
		f.Close()
		os.Remove(f.Name())
	}()

	cmd := exec.Command("vim", f.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()

	b, err := ioutil.ReadFile(f.Name())
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// NewClient create
func NewClient(token string) *Client {
	c := &Client{token}
	return c
}

func (c *Client) request(path string, method string, params interface{}, response interface{}) error {
	url := APIURL + path

	var body io.Reader
	if params != nil {
		json, err := json.Marshal(params)
		if err != nil {
			return err
		}
		body = bytes.NewBuffer(json)
	}
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "token "+c.token)
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return err
	}
	return nil
}

// CreateIssue POST /repos/:owner/:repo/issues
func CreateIssue(owner, repository string) (string, error) {

	return "hello", fmt.Errorf("Error dayo")
}

// GetIssue GET
func (c *Client) GetIssue() {

}

// EditIssue POST
func EditIssue() {

}

// CloseIssue POST
func CloseIssue() {

}
