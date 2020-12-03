package github_provider

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Djudicael/microserviceMVC/src/api/clients/restclient"
	"github.com/Djudicael/microserviceMVC/src/api/domain/github"
)

const (
	headerAuthozitation       = "Authorization"
	headerAuthozitationFormat = "token %s"
	urlCreateRepo             = "https://api.github.com/user/repo"
)

func GetAuthozitationHeader(accessToken string) string {
	return fmt.Sprintf(headerAuthozitationFormat, accessToken)
}

func CreateRepo(accessToken string, request github.CreateRepoRequest) (*github.CreateRepoResponse, *github.GithubErrorResponse) {

	headers := http.Header{}
	headers.Set(headerAuthozitation, GetAuthozitationHeader(accessToken))

	response, err := restclient.Post(urlCreateRepo, request, headers)
	if err != nil {
		log.Println(fmt.Sprintf("error when trying to create new repo in github: %s", err.Error()))
		return nil, &github.GithubErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	bytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, &github.GithubErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "invalid response body",
		}
	}

	defer response.Body.Close()

	if response.StatusCode > 299 {
		var errResponse github.GithubErrorResponse
		if err := json.Unmarshal(bytes, &errResponse); err != nil {

			return nil, &github.GithubErrorResponse{
				StatusCode: http.StatusInternalServerError,
				Message:    "invalid json  response body",
			}
		}

		errResponse.StatusCode = response.StatusCode
		return nil, &errResponse

	}

	var result github.CreateRepoResponse

	if err := json.Unmarshal(bytes, &result); err != nil {
		log.Println(fmt.Sprintf("error when trying to unmarshal create repo successful response: %s", err.Error()))

		return nil, &github.GithubErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "error unmarchalling github create repo response",
		}
	}

	return &result, nil
}
