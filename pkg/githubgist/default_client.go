package githubgist

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type gistJsonResponse struct {
	HtmlUrl string `json:"html_url"`
}

type gistListJsonResponse []gistJsonResponse

type defaultClient struct {
	*http.Client
}

var c = &defaultClient{
	// TODO: Add timeout for the HTTP Client
	http.DefaultClient,
}

func NewClient() Client {
	return c
}

type UserNotFoundErr string

func (err UserNotFoundErr) Error() string {
	return string(err)
}

func (c *defaultClient) ListGists(username string, page int) (Gists, error) {
	// TODO: Add `Accept` and `X-GitHub-Api-Version` headers
	// TODO: Add timeouts using Context API
	// TODO: Send pagination inputs - to paginate results and not just show
	// first page always
	// TODO: Check and Sanitize `username` parameter, either here or in
	// the caller - so as to avoid any injection through user input
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.github.com/users/%s/gists?page=%d", username, page), nil)

	if err != nil {
		return nil, fmt.Errorf("error forming a request: %v", err)
	}

	resp, err := c.Do(req)

	if err != nil {
		return nil, fmt.Errorf("error sending a request: %v", err)
	}

	defer resp.Body.Close()

	// Read the body into a byte slice
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	statusCode := resp.StatusCode

	if statusCode == 404 {
		return nil, UserNotFoundErr(fmt.Sprintf("response code is %v which means http client error. Response body is: %v", statusCode, string(body)))
	}

	if statusCode >= 400 && statusCode <= 499 {
		return nil, fmt.Errorf("response code is %v which means http client error: %v. Response body is: %v", statusCode, err, string(body))
	}

	if statusCode >= 500 && statusCode <= 599 {
		return nil, fmt.Errorf("response code is %v which means http server error: %v. Response body is: %v", statusCode, err, string(body))
	}

	if statusCode != 200 {
		return nil, fmt.Errorf("response code is %v which is an unexpected response code. Response body is: %v", statusCode, string(body))
	}

	var gistList gistListJsonResponse

	err = json.Unmarshal(body, &gistList)
	if err != nil {
		return nil, fmt.Errorf("error parsing response body json: %v. Response code is: %v. Response body is: %v", err, statusCode, string(body))
	}

	gists := make(Gists, 0, len(gistList))

	for _, gist := range gistList {
		gists = append(gists, Gist{Url: gist.HtmlUrl})
	}

	return gists, nil
}
