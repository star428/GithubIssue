package githubissue

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	GITHUB_TOKEN = "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" // change to yourself
	GITHUB_API   = "https://api.github.com"
	GITHUB_OWNER = "star428"     // change to yourself
	GITHUB_REPO  = "GithubIssue" // change to yourself
)

type Issue struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	State string `json:"state"`
}

type Comment struct {
	Body string `json:"body"`
}

func CreateIssue(owner, repo, title, body string) (*Issue, error) {
	url := fmt.Sprintf("%s/repos/%s/%s/issues", GITHUB_API, owner, repo)
	fmt.Println(url)
	issue := &Issue{Title: title, Body: body}

	b, err := json.Marshal(issue)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "token "+GITHUB_TOKEN)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	var responseIssue Issue
	if err := json.NewDecoder(resp.Body).Decode(&responseIssue); err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &responseIssue, err
}

func ReadIssue(owner, repo, issueNumber string) (*Issue, error) {
	url := fmt.Sprintf("%s/repos/%s/%s/issues/%s", GITHUB_API, owner, repo, issueNumber)
	fmt.Println(url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "token "+GITHUB_TOKEN)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	var responseIssue Issue
	if err := json.NewDecoder(resp.Body).Decode(&responseIssue); err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &responseIssue, err
}

func UpdateIssue(owner, repo, issueNumber, title, body string, isclosed bool) (*Issue, error) {
	url := fmt.Sprintf("%s/repos/%s/%s/issues/%s", GITHUB_API, owner, repo, issueNumber)
	var closeOropen string
	if isclosed {
		closeOropen = "closed"
	} else {
		closeOropen = "open"
	}

	issue := &Issue{Title: title, Body: body, State: closeOropen}
	b, err := json.Marshal(issue)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "token "+GITHUB_TOKEN)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	var responseIssue Issue
	if err := json.NewDecoder(resp.Body).Decode(&responseIssue); err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &responseIssue, err
}

func CloseIssue(owner, repo, issueNumber string) (*Issue, error) {
	issue, err := ReadIssue(owner, repo, issueNumber)
	if err != nil {
		return nil, err
	}
	issue, err = UpdateIssue(owner, repo, issueNumber, issue.Title, issue.Body, true)
	if err != nil {
		return nil, err
	}

	return issue, nil
}

func AddComment(owner, repo, issueNumber, comment string) (*Comment, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/issues/%s/comments", owner, repo, issueNumber)
	commentData := &Comment{
		Body: comment,
	}
	b, err := json.Marshal(commentData)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "token "+GITHUB_TOKEN)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("failed to add comment, status code: %d", resp.StatusCode)
	}
	var result Comment
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
