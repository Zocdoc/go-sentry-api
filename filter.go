package sentry

import (
	"fmt"
)

type Filter struct {
	Id     string `json:"id"`
	Active *bool  `json:"active,omitempty"`
}

type SubFilter struct {
	Id         string   `json:"id"`
	SubFilters []string `json:"subfilters,omitempty"`
}

func (c *Client) UpdateFilter(o Organization, p Project, filter Filter) error {

	err := c.do("PUT", fmt.Sprintf("projects/%s/%s/filters/%s", *o.Slug, *p.Slug, filter.Id), nil, &filter)
	return err
}

func (c *Client) UpdateSubFilter(o Organization, p Project, filter SubFilter) error {

	err := c.do("PUT", fmt.Sprintf("projects/%s/%s/filters/%s", *o.Slug, *p.Slug, filter.Id), nil, &filter)
	return err
}
