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

		subFilter := SubFilter{
			Id:         "legacy-browsers",
			SubFilters: []string{"ie_pre_9", "android_pre_4", "ie10", "opera_pre_15", "safari_pre_6", "ie9"},
		}

		err = client.UpdateSubFilter(org, project, subFilter)
		if err != nil {
			t.Error(err)
		}
	})

	if err := client.DeleteTeam(org, team); err != nil {
		t.Error(err)
	}

}
