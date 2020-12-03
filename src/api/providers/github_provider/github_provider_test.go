package github_provider

import (
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/Djudicael/microserviceMVC/src/api/clients/restclient"
	"github.com/Djudicael/microserviceMVC/src/api/domain/github"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	restclient.StartMockups()
	os.Exit(m.Run())

}

func TestConstants(t *testing.T) {
	assert.EqualValues(t, "Authorization", headerAuthozitation)

	assert.EqualValues(t, "token %s", headerAuthozitationFormat)
}

func TestGetAuthorizationHeader(t *testing.T) {
	header := GetAuthozitationHeader("abc123")
	assert.EqualValues(t, "token abc123", header)
}

func TestCreateRepoErrorRestclient(t *testing.T) {
	restclient.FlushMockups()
	restclient.Addmockup(restclient.Mock{
		Url:        "https://api.github.com/user/repo",
		HttpMethod: http.MethodPost,
		Err:        errors.New("Invalide rest client"),
	})
	response, err := CreateRepo("", github.CreateRepoRequest{})
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "Invalide rest client", err.Message)

}
func TestCreateRepoInvalidresponseBody(t *testing.T) {
	restclient.FlushMockups()
	invalidCloser, _ := os.Open("-asf3")
	restclient.Addmockup(restclient.Mock{
		Url:        "https://api.github.com/user/repo",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       invalidCloser,
		},
	})
	response, err := CreateRepo("", github.CreateRepoRequest{})
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "invalid response body", err.Message)

}
func TestCreateRepoInvalidErrorInterface(t *testing.T) {
	restclient.FlushMockups()
	restclient.Addmockup(restclient.Mock{
		Url:        "https://api.github.com/user/repo",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       ioutil.NopCloser(strings.NewReader(`{"message":1}`)),
		},
	})
	response, err := CreateRepo("", github.CreateRepoRequest{})
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "invalid json  response body", err.Message)

}
func TestCreateRepoSucessInvalidResponse(t *testing.T) {
	restclient.FlushMockups()
	restclient.Addmockup(restclient.Mock{
		Url:        "https://api.github.com/user/repo",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       ioutil.NopCloser(strings.NewReader(`{"message":"Requires authentication","documentation_url":"fffgg"}`)),
		},
	})
	response, err := CreateRepo("", github.CreateRepoRequest{})
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusUnauthorized, err.StatusCode)
	assert.EqualValues(t, "Requires authentication", err.Message)

}
func TestCreateRepoInvalidUnauthorize(t *testing.T) {
	restclient.FlushMockups()
	restclient.Addmockup(restclient.Mock{
		Url:        "https://api.github.com/user/repo",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id":"123"}`)),
		},
	})
	response, err := CreateRepo("", github.CreateRepoRequest{})
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "error unmarchalling github create repo response", err.Message)

}
func TestCreateRepoNoError(t *testing.T) {
	restclient.FlushMockups()
	restclient.Addmockup(restclient.Mock{
		Url:        "https://api.github.com/user/repo",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id":123,"name":"tata","full_name":"tatayoyo"}`)),
		},
	})
	response, err := CreateRepo("", github.CreateRepoRequest{})
	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.EqualValues(t, 123, response.ID)
	assert.EqualValues(t, "tata", response.Name)
	assert.EqualValues(t, "tatayoyo", response.Fullname)
}
