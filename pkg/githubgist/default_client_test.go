package githubgist_test

import (
	"errors"
	"proxy/pkg/githubgist"
	"testing"
)

func TestListGists(t *testing.T) {

	t.Run("users who exist", func(t *testing.T) {
		username := "sagikazarmark"
		gists, err := githubgist.NewClient().ListGists(username)
		if err != nil {
			t.Errorf("expected no error while listing gists of the user %s, but got one error: %v", username, err)
		}

		// For Debugging the test
		t.Logf("the list of gists of %s", username)
		for _, gist := range gists {
			t.Log(gist)
		}

		username = "iximiuz"
		gists, err = githubgist.NewClient().ListGists(username)
		if err != nil {
			t.Errorf("expected no error while listing gists of the user %s, but got one error: %v", username, err)
		}

		// For Debugging the test
		t.Logf("the list of gists of %s", username)
		for _, gist := range gists {
			t.Log(gist)
		}
	})

	t.Run("user who does not exist", func(t *testing.T) {
		username := "non-existent-user-big-username-unimaginable-wow-great"
		gists, err := githubgist.NewClient().ListGists(username)
		if err == nil {
			t.Errorf("expected error while listing gists of the user that is supposed to be non existent - %s, but got no error", username)
		}

		_ = gists

		if e, ok := errors.AsType[githubgist.UserNotFoundErr](err); ok {
			_ = e
			// For debugging
			t.Logf("as expected, we got an error. error is %v", e)
		} else {
			t.Errorf("expected a specific kind of error while listing gists of a user that is supposed to be non existent - user: %s, but got a different error: error type: %T, error: %v", username, err, err)
		}
	})

	// What happens when there are characters like space in the username?
	// As the current implementation does string interpolation of the username
	// inside the URL as is - whatever the user gives as input.
	// Looks like github usernames cannot have space in them - so,
	// any username with space in it should give user not found
	t.Run("space in username", func(t *testing.T) {
		username := "something lol"
		gists, err := githubgist.NewClient().ListGists(username)
		if err == nil {
			t.Errorf("expected error while listing gists of the user that is supposed to be non existent - %s, but got no error", username)
		}

		_ = gists

		if e, ok := errors.AsType[githubgist.UserNotFoundErr](err); ok {
			_ = e
			// For debugging
			t.Logf("as expected, we got an error. error is %v", e)
		} else {
			t.Errorf("expected a specific kind of error while listing gists of a user that is supposed to be non existent - user: %s, but got a different error: error type: %T, error: %v", username, err, err)
		}
	})
}
