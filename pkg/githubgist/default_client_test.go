package githubgist_test

import (
	"errors"
	"proxy/pkg/githubgist"
	"testing"
)

func TestListGists(t *testing.T) {
	username := "sagikazarmark"
	gists, err := githubgist.NewClient().ListGists(username)
	if err != nil {
		t.Errorf("expected no error while listing gists of the user %s, but got one error: %v", username, err)
	}

	_ = gists
	// For Debugging the test
	// t.Logf("the list of gists of %s", username)
	// for _, gist := range gists {
	// 	t.Log(gist)
	// }

	username = "iximiuz"
	gists, err = githubgist.NewClient().ListGists(username)
	if err != nil {
		t.Errorf("expected no error while listing gists of the user %s, but got one error: %v", username, err)
	}

	_ = gists
	// For Debugging the test
	// t.Logf("the list of gists of %s", username)
	// for _, gist := range gists {
	// 	t.Log(gist)
	// }

	username = "non-existent-user-big-username-unimaginable-wow-great"
	// Do one test with a non existent user and check if it gives error - like a 404 error
	gists, err = githubgist.NewClient().ListGists(username)
	if err == nil {
		t.Errorf("expected error while listing gists of the user that is supposed to be non existent - %s, but got no error", username)
	}

	if e, ok := errors.AsType[githubgist.UserNotFoundErr](err); ok {
		_ = e
		// For debugging
		// t.Logf("as expected, we got an error. error is %v", e)
	} else {
		t.Errorf("expected a specific kind of error while listing gists of a user that is supposed to be non existent - user: %s, but got a different error: error type: %T, error: %v", username, err, err)
	}

	// TODO: Check what happens when there are characters like space in the username? As the current implementation does string interpolation of the username inside the URL as is - whatever the user gives as input
}
