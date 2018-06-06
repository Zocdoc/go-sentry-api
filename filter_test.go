package sentry

import (
	"testing"
)

func TestProjectFilters(t *testing.T) {

	client := newTestClient(t)
	org, err := client.GetOrganization(getDefaultOrg())
	if err != nil {
		t.Fatal(err)
	}

	team, cleanup := createTeamHelper(t)
	defer cleanup()

	project, cleanupproj := createProjectHelper(t, team)
	defer cleanupproj()

	t.Run("Update filters", func(t *testing.T) {
		active := true
		filter := Filter{
			Id:     "browser-extensions",
			Active: &active,
		}

		err := client.UpdateFilter(org, project, filter)
		if err != nil {
			t.Error(err)
		}
	})

	if err := client.DeleteTeam(org, team); err != nil {
		t.Error(err)
	}

}
