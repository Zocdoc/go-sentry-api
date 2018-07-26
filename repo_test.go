package sentry

import (
	"testing"
)

func TestRepos(t *testing.T) {

	client := newTestClient(t)
	org, err := client.GetOrganization(getDefaultOrg())
	if err != nil {
		t.Fatal(err)
	}

	team, cleanup := createTeamHelper(t)
	defer cleanup()

	t.Run("List repos", func(t *testing.T) {

		_, err := client.GetRepos(org)
		if err != nil {
			t.Error(err)
		}
	})

	if err := client.DeleteTeam(org, team); err != nil {
		t.Error(err)
	}

}
