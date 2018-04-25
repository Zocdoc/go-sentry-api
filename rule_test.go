package sentry

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestRuleResource(t *testing.T) {

	client := newTestClient(t)
	org, err := client.GetOrganization(getDefaultOrg())
	if err != nil {
		t.Fatal(err)
	}

	team, cleanup := createTeamHelper(t)
	defer cleanup()

	project, cleanupproj := createProjectHelper(t, team)
	defer cleanupproj()

	t.Run("Get all rules for a project", func(t *testing.T) {
		rules, err := client.GetRules(org, project)
		if err != nil {
			t.Error(err)
		}
		if len(rules) != 1 {
			t.Fatalf("Expected a single rule got %d rules", len(rules))
		}

		fmt.Printf("%+v", rules)

		ruleStr := `
		{
		  "actionMatch": "all",
		  "actions": [
		    {
		      "id": "sentry.rules.actions.notify_event_service.NotifyEventServiceAction",
		      "service": "mail"
		    }
		  ],
		  "conditions": [
		    {
		      "id": "sentry.rules.conditions.event_frequency.EventFrequencyCondition",
		      "interval": "1h",
		      "value": "10"
		    }
		  ],
		  "name": "Send Email for Error Burst",
		  "frequency": 30
		}
		`

		for _, e := range os.Environ() {
			pair := strings.Split(e, "=")
			fmt.Println(pair)
		}

		rule := Rule{}
		err = json.Unmarshal([]byte(ruleStr), &rule)
		if err != nil {
			t.Error(err)
		}

		fmt.Printf("%+v", rule)

		newRule, err := client.CreateRule(org, project, rule)
		if err != nil {
			t.Error(err)
		}

		if newRule.Actions[0].Id != "sentry.rules.actions.notify_event_service.NotifyEventServiceAction" {
			t.Errorf("Expected action to be NotifyEventServiceAction got %s", newRule.Actions[0].Id)
		}

		condition := newRule.Conditions[0]
		if condition.Id != "sentry.rules.conditions.event_frequency.EventFrequencyCondition" {
			t.Errorf("Expected condition to be EventFrequencyCondition got %s", condition.Id)
		}

		if condition.Interval != "1h" {
			t.Errorf("Expected interval to be 1h got %s", condition.Interval)
		}

		if condition.Value != "10" {
			t.Errorf("Expected value to be 10 got %s", condition.Value)
		}

		rules, err = client.GetRules(org, project)
		if err != nil {
			t.Error(err)
		}

		if len(rules) != 2 {
			t.Fatalf("Expected 2 rules got %d rules", len(rules))
		}

		err = client.DeleteRule(org, project, newRule.Id)
		if err != nil {
			t.Error(err)
		}

		rules, err = client.GetRules(org, project)
		if err != nil {
			t.Error(err)
		}
		if len(rules) != 1 {
			t.Fatalf("Expected 1 rules got %d rules", len(rules))
		}
	})

	if err := client.DeleteTeam(org, team); err != nil {
		t.Error(err)
	}

}
