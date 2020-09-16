package services

import (
	"strings"

	"github.com/davetweetlive/golang-microservices/src/api/config"
	"github.com/davetweetlive/golang-microservices/src/api/domain/github"
	"github.com/davetweetlive/golang-microservices/src/api/domain/repositories"
	"github.com/davetweetlive/golang-microservices/src/api/providers/github_provider"
	"github.com/davetweetlive/golang-microservices/src/api/utils/errors"
)

type repoService struct {
}

type repoServiceInterface interface {
	CreateRepo(request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.APIError)
}

var (
	RepositoryService repoServiceInterface
)

func init() {
	RepositoryService = &repoService{}
}

func (s *repoService) CreateRepo(input repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.APIError) {
	input.Name = strings.TrimSpace(input.Name)
	if input.Name == "" {
		return nil, errors.NewBadRequestError("Invalid repository name!")
	}

	request := github.CreateRepoRequest{
		Name:        input.Name,
		Private:     false,
		Description: input.Description,
	}

	response, err := github_provider.CreateRepo(config.GetGithubAccessToken(), request)
	if err != nil {
		apiErr := errors.NewApiError(err.StatusCode, err.Message)
		return nil, apiErr
	}

	result := repositories.CreateRepoResponse{
		Id:    response.Id,
		Name:  response.Name,
		Owner: response.Owner.Login,
	}
	return &result, nil
}
