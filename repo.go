package sentry

import (
	"fmt"
	"time"
)

type NewRepo struct {
	Name     string `json:"name,omitempty"`
	Provider string `json:"provider,omitempty"`
}

type Repo struct {
	Id          string        `json:"id,omitempty"`
	Status      string        `json:"status,omitempty"`
	Name        string        `json:"name,omitempty"`
	Url         string        `json:"url,omitempty"`
	Provider    *RepoProvider `json:"provider,omitempty"`
	DateCreated *time.Time    `json:"dateCreated,omitempty"`
}

type RepoProvider struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func (c *Client) CreateRepo(o Organization, repo NewRepo) (Repo, error) {
	var outputRepo Repo
	err := c.do("POST", fmt.Sprintf("organizations/%s/repos", *o.Slug), &outputRepo, &repo)
	return outputRepo, err
}

func (c *Client) GetRepos(o Organization) ([]Repo, error) {
	repos := make([]Repo, 0)
	link, err := c.rawWithPagination("GET", fmt.Sprintf("organizations/%s/repos/", *o.Slug), &repos, nil)
	if err != nil {
		return repos, err
	}

	// Keep fetching while there are more pages
	for link != nil && link.Next.Results {
		var nextRepos []Repo
		link, err = c.GetPage(link.Next, &nextRepos)
		if err != nil {
			return repos, err
		}
		repos = append(repos, nextRepos...)
	}

	return repos, nil
}
