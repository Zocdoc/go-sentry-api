package sentry

import (
	"fmt"
	"time"
)

type Repo struct {
	Id             string        `json:"id,omitempty"`
	Status         string        `json:"status,omitempty"`
	Name           string        `json:"name,omitempty"`
	Url            string        `json:"url,omitempty"`
	Provider       *RepoProvider `json:"provider,omitempty"`
	ProviderString string        `json:"provider,omitempty"`
	DateCreated    *time.Time    `json:"dateCreated,omitempty"`
}

type RepoProvider struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func (c *Client) CreateRepo(o Organization, repo Repo) error {
	err := c.do("POST", fmt.Sprintf("organizations/%s/repos", *o.Slug), nil, &repo)
	return err
}

func (c *Client) GetRepos(o Organization) ([]Repo, error) {
	repos := make([]Repo, 0)
	err := c.do("GET", fmt.Sprintf("organizations/%s/repos", *o.Slug), &repos, nil)
	return repos, err
}
