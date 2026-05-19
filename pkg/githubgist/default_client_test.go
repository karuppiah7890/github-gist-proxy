package githubgist_test

import (
	"proxy/pkg/githubgist"
	"testing"
)

func TestListGists(t *testing.T) {
	gists, err := githubgist.NewClient().ListGists("sagikazarmark")
	if err != nil {
		t.Errorf("expected no error while listing gists of the user sagikazarmark, but got one error: %v", err)
	}

	t.Log("the list of gists of sagikazarmark")

	for _, gist := range gists {
		t.Log(gist)
	}

	gists, err = githubgist.NewClient().ListGists("iximiuz")
	if err != nil {
		t.Errorf("expected no error while listing gists of the user iximiuz, but got one error: %v", err)
	}

	t.Log("the list of gists of iximiuz")

	for _, gist := range gists {
		t.Log(gist)
	}
}
