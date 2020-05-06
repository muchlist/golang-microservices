package restclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var (
	enableMock = false
	mocks      = make(map[string]*Mock)
)

//Mock is struct yang digunakan untuk keperluan Mocking response
type Mock struct {
	URL        string
	HTTPMethod string
	Response   *http.Response
	Error      error
}

func getMockID(httpMethod string, url string) string {
	return fmt.Sprintf("%s_%s", httpMethod, url)
}

//StartMockUp membuat enableMock menjadi true untuk keperluan testing
func StartMockUp() {
	enableMock = true
}

//FlushMockups is membuat nilai mock pada test sebelumnya menjadi mock kosong
func FlushMockups() {
	mocks = make(map[string]*Mock)
}

//StopMockUp membuat enableMock menjadi false
func StopMockUp() {
	enableMock = false
}

//AddMockUp menambahkan mock file untuk testing
func AddMockUp(mock Mock) {
	mocks[getMockID(mock.HTTPMethod, mock.URL)] = &mock
}

//Post membuat repo github
func Post(url string, body interface{}, headers http.Header) (*http.Response, error) {
	if enableMock {
		mock := mocks[getMockID(http.MethodPost, url)]
		if mock == nil {
			return nil, errors.New("no mockup found for give request")
		}
		return mock.Response, mock.Error
	}

	jsonBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonBytes))
	request.Header = headers

	client := http.Client{}
	return client.Do(request)
}
