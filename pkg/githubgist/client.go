package githubgist

type Gist struct {
	Url string
}

type Gists []Gist

type Client interface {
	ListGists(username string) (Gists, error)
}
