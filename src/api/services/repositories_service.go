package services

import (
	"strings"

	"github.com/Djudicael/microserviceMVC/src/api/config"
	"github.com/Djudicael/microserviceMVC/src/api/domain/github"
	"github.com/Djudicael/microserviceMVC/src/api/domain/repositories"
	"github.com/Djudicael/microserviceMVC/src/api/providers/github_provider"
	"github.com/Djudicael/microserviceMVC/src/api/utils/errors"
)

type repoService struct {
}

type repoServiceInterface interface {
	CreateRepo(request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError)
}

var (
	RepositoryService repoServiceInterface
)

func init() {
	RepositoryService = &repoService{}
}

func (s *repoService) CreateRepo(input repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError) {
	input.Name = strings.TrimSpace(input.Name)
	if input.Name == "" {
		return nil, errors.NewBadrequestError("Invalid repository name")
	}
	request := github.CreateRepoRequest{
		Name:        input.Name,
		Description: input.Description,
		Private:     true,
	}

	response, err := github_provider.CreateRepo(config.GetgithubAccessToken(), request)

	if err != nil {
		return nil, errors.NewApiError(err.StatusCode, err.Message)
	}

	result := repositories.CreateRepoResponse{
		ID:    response.ID,
		Name:  response.Name,
		Owner: response.Owner.Login,
	}

	return &result, nil
}
