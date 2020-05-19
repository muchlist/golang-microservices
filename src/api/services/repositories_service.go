package services

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/muchlist/golang-microservices/src/api/domain/github"
	"github.com/muchlist/golang-microservices/src/api/log"
	"github.com/muchlist/golang-microservices/src/api/providers/githubprovider"

	"github.com/muchlist/golang-microservices/src/api/config"
	"github.com/muchlist/golang-microservices/src/api/domain/repositories"
	"github.com/muchlist/golang-microservices/src/api/utils/errors"
)

type repoService struct{}

type repoServiceInterface interface {
	CreateRepo(clientID string, request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError)
	CreateRepos(clientID string, request []repositories.CreateRepoRequest) (repositories.CreateReposResponse, errors.ApiError)
}

var (
	RepositoryService repoServiceInterface
)

func init() {
	RepositoryService = &repoService{}
}

func (s *repoService) CreateRepo(clientID string, input repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	request := github.CreateRepoRequest{
		Name:        input.Name,
		Private:     false,
		Description: input.Description,
	}

	log.Info("about to send request to external api", fmt.Sprintf("client_id:%s", clientID), "status:pending")

	response, err := githubprovider.CreateRepo(config.GetGithubAccessToken(), request)
	if err != nil {
		log.Error("response obtained from external api", err, fmt.Sprintf("client_id:%s", clientID), "status:error")
		apiErr := errors.NewApiError(err.StatusCode, err.Message)
		return nil, apiErr
	}

	log.Info("response obtained from external api", fmt.Sprintf("client_id:%s", clientID), "status:success")
	result := repositories.CreateRepoResponse{
		ID:    response.ID,
		Name:  response.Name,
		Owner: response.Owner.Login,
	}

	return &result, nil
}

func (s *repoService) CreateRepos(clientID string, request []repositories.CreateRepoRequest) (repositories.CreateReposResponse, errors.ApiError) {
	input := make(chan repositories.CreateRepositoriesResult)
	output := make(chan repositories.CreateReposResponse)
	defer close(output)

	var wg sync.WaitGroup
	go s.handleRepoResults(&wg, input, output)
	for _, current := range request {
		wg.Add(1)
		go s.createRepoConcurent(clientID, current, input)
	}

	wg.Wait()
	close(input)
	result := <-output

	succesCreations := 0
	for _, current := range result.Result {
		if current.Response != nil {
			succesCreations++
		}
	}

	if succesCreations == 0 {
		result.StatusCode = result.Result[0].Error.Status()
	} else if succesCreations == len(request) {
		result.StatusCode = http.StatusCreated
	} else {
		result.StatusCode = http.StatusPartialContent
	}

	return result, nil
}

func (s *repoService) handleRepoResults(wg *sync.WaitGroup, input chan repositories.CreateRepositoriesResult, output chan repositories.CreateReposResponse) {
	var results repositories.CreateReposResponse

	for incomingEvent := range input {
		repoResult := repositories.CreateRepositoriesResult{
			Response: incomingEvent.Response,
			Error:    incomingEvent.Error,
		}
		results.Result = append(results.Result, repoResult)
		wg.Done()
	}

	output <- results
}

func (s *repoService) createRepoConcurent(clientID string, input repositories.CreateRepoRequest, output chan repositories.CreateRepositoriesResult) {
	if err := input.Validate(); err != nil {
		output <- repositories.CreateRepositoriesResult{
			Error: err,
		}
		return
	}

	result, err := s.CreateRepo(fmt.Sprintf("client_id:%s", clientID), input)
	if err != nil {
		output <- repositories.CreateRepositoriesResult{
			Error: err,
		}
		return
	}

	output <- repositories.CreateRepositoriesResult{
		Response: result,
	}
}
