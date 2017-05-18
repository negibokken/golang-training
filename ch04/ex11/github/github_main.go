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

type GetIssueResponse struct {
	ID      int    `json:"url"`
	URL     string `json:"url"`
	RepoURL string `json:"repository_url"`
	State   string `json:"state"`
	Title   string `json:"title"`
}

func (c *Client) request(method string, path string, params interface{}, response interface{}) error {
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
func (c *Client) CreateIssue(owner, repository string) error {
	path := fmt.Sprintf("/repos/%s/%s/issues", owner, repository)
	str, err := getStringByEditor()
	if err != nil {
		fmt.Println(err)
		return err
	}

	var response interface{}
	var params interface{}
	json.Unmarshal([]byte(str), &params)

	err = c.request("POST", path, params, &response)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Issue is created")
	fmt.Println(response)
	return nil
}

// GetIssue GET
func (c *Client) GetIssue(owner, repository, issue string) error {
	path := fmt.Sprintf("/repos/%s/%s/issues/%s", owner, repository, issue)
	fmt.Println(path)
	var response interface{}

	err := c.request("GET", path, nil, &response)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Got Issue")
	json.Marshal(response)
	fmt.Printf("%#v", response)
	return nil
}

// EditIssue POST
func (c *Client) EditIssue(owner, repository, issue string) error {
	path := fmt.Sprintf("/repos/%s/%s/issues/%s", owner, repository, issue)

	str, err := getStringByEditor()
	if err != nil {
		fmt.Println(err)
		return err
	}

	var response interface{}
	var params interface{}
	json.Unmarshal([]byte(str), &params)

	err = c.request("PATCH", path, params, &response)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Edited Issue")
	return nil
}

// CloseIssue POST
func (c *Client) CloseIssue(owner, repository, issue string) error {
	path := fmt.Sprintf("/repos/%s/%s/issues/%s", owner, repository, issue)

	str := `{"state": "closed"} `
	var response interface{}
	var params interface{}
	json.Unmarshal([]byte(str), &params)
	err := c.request("PATCH", path, params, &response)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Closed Issue")
	return nil
}
