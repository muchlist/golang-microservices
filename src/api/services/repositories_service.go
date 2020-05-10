package services

import (
	"strings"

	"github.com/muchlist/golang-microservices/src/api/domain/github"
	"github.com/muchlist/golang-microservices/src/api/providers/githubprovider"

	"github.com/muchlist/golang-microservices/src/api/config"
	"github.com/muchlist/golang-microservices/src/api/domain/repositories"
	"github.com/muchlist/golang-microservices/src/api/utils/errors"
)

type repoService struct{}

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
		return nil, errors.NewBadRequestError("invalid repository name")
	}

	request := github.CreateRepoRequest{
		Name:        input.Name,
		Private:     false,
		Description: input.Description,
	}

	response, err := githubprovider.CreateRepo(config.GetGithubAccessToken(), request)
	if err != nil {
		apiErr := errors.NewApiError(err.StatusCode, err.Message)
		return nil, apiErr
	}

	result := repositories.CreateRepoResponse{
		ID:    response.ID,
		Name:  response.Name,
		Owner: response.Owner.Login,
	}

	return &result, nil
}
