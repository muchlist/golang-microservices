package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"

	"github.com/muchlist/golang-microservices/src/api/domain/repositories"
	"github.com/muchlist/golang-microservices/src/api/services"
	"github.com/muchlist/golang-microservices/src/api/utils/errors"
)

var (
	success map[string]string
	failed  map[string]errors.ApiError
)

type createRepoResult struct {
	Request repositories.CreateRepoRequest
	Result  *repositories.CreateRepoResponse
	Error   errors.ApiError
}

func getRequests() []repositories.CreateRepoRequest {
	result := make([]repositories.CreateRepoRequest, 0)

	file, err := os.Open("Users/admin/Desktop/requests.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		request := repositories.CreateRepoRequest{
			Name: line,
		}
		result = append(result, request)
	}

	return result
}

func main() {
	requests := getRequests()
	fmt.Println(fmt.Sprintf("about to process %d requests", len(requests)))

	input := make(chan createRepoResult)
	buff := make(chan bool, 10)
	var wg sync.WaitGroup

	go handResults(&wg, input)

	for _, request := range requests {
		buff <- true
		wg.Add(1)
		go createRepo(buff, input, request)
	}

	wg.Wait()
	close(input)

	//sekarang baru kita tulis sukses atau failed map ke disk atau pemberitahuan via email
}

func handResults(wg *sync.WaitGroup, input chan createRepoResult) {
	for result := range input {
		if result.Error != nil {
			failed[result.Request.Name] = result.Error
		} else {
			success[result.Request.Name] = result.Result.Name
		}
		wg.Done()
	}
}

func createRepo(buffer chan bool, output chan createRepoResult, request repositories.CreateRepoRequest) {
	result, err := services.RepositoryService.CreateRepo(request)
	output <- createRepoResult{
		Request: request,
		Result:  result,
		Error:   err,
	}

	<-buffer
}
