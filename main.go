package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"proxy/pkg/githubgist"
	"strconv"
)

func main() {
	fmt.Println("Running GitHub Gist Proxy")

	mux := http.NewServeMux()

	mux.HandleFunc("GET /livez", livenessHandler)
	mux.HandleFunc("GET /{username}", usernameHandler)

	server := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()

	if err != nil {
		log.Fatalf("error occurred while trying to run server: %v", err)
	}
}

func livenessHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func usernameHandler(w http.ResponseWriter, r *http.Request) {
	username := r.PathValue("username")
	page := 1
	if r.URL.Query().Has("page") {
		pageInput := r.URL.Query().Get("page")
		pageNumber, err := strconv.Atoi(pageInput)
		if err != nil {
			fmt.Printf("an error occurred while converting %s page input to number: %v", pageInput, err)
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Add("Content-Type", "application/json")
			w.Write([]byte(fmt.Sprintf("{\"message\": \"page '%s' is not a number\"}", pageInput)))
			return
		}
		page = pageNumber
	}

	gists, err := githubgist.NewClient().ListGists(username, page)
	if err != nil {
		if e, ok := errors.AsType[githubgist.UserNotFoundErr](err); ok {
			fmt.Printf("an error occurred while listing gists of the user %s: %v. it seems like a user not found error: %v", username, err, e)
			w.WriteHeader(http.StatusNotFound)
			w.Header().Add("Content-Type", "application/json")
			w.Write([]byte(fmt.Sprintf("{\"message\": \"username '%s' does not exist\"}", username)))
		} else {
			reportUnexpectedError(err, w, username)
		}
		return
	}

	gistsJson, err := json.Marshal(gists)
	if err != nil {
		reportUnexpectedError(err, w, username)
		return
	}
	w.Write(gistsJson)
}

func reportUnexpectedError(err error, w http.ResponseWriter, username string) {
	fmt.Printf("an unexpected error occurred while listing gists of the user %s: %v", username, err)
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte(fmt.Sprintf("{\"message\": \"an error occurred while getting the gists of the username '%s'. please contact support or the api admin to resolve the issue\"}", username)))
}
