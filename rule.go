package sentry

import (
	"fmt"
	"time"
)

type Rule struct {
	ActionMatch string      `json:"actionMatch"`
	Actions     []Action    `json:"actions"`
	Conditions  []Condition `json:"conditions"`
	Name        string      `json:"name"`
	Frequency   int         `json:"frequency"`
	Id          string      `json:"id,omitempty"`
	DateCreated time.Time   `json:"dateCreated,omitempty"`
}

type Action struct {
	Id      string `json:"id"`
	Service string `json:"service,omitempty"`
}

type Condition struct {
	Id       string `json:"id"`
	Interval string `json:"interval,omitempty"`
	Value    string `json:"value,omitempty"`
}

// GetRules fetchs all rules for a given project
func (c *Client) GetRules(o Organization, p Project) ([]Rule, error) {
	var rules []Rule

	err := c.do("GET", fmt.Sprintf("projects/%s/%s/rules", *o.Slug, *p.Slug), &rules, nil)
	return rules, err
}

func (c *Client) CreateRule(o Organization, p Project, rule Rule) (Rule, error) {
	var resRule Rule

	err := c.do("POST", fmt.Sprintf("projects/%s/%s/rules", *o.Slug, *p.Slug), &resRule, &rule)
	return resRule, err
}

func (c *Client) DeleteRule(o Organization, p Project, ruleId string) error {
	err := c.do("DELETE", fmt.Sprintf("projects/%s/%s/rules/%s", *o.Slug, *p.Slug, ruleId), nil, nil)
	return err
}
