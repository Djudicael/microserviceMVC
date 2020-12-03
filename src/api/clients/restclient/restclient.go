package restclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var (
	enableMocks = false
	mocks       = make(map[string]*Mock)
)

type Mock struct {
	Url        string
	HttpMethod string
	Response   *http.Response
	Err        error
}

func getMockeId(httpMethod string, url string) string {
	return fmt.Sprintf("%s_%s", httpMethod, url)
}

func StartMockups() {
	enableMocks = true

}

func FlushMockups() {
	mocks = make(map[string]*Mock)
}

func StopMockups() {
	enableMocks = false

}

func Addmockup(mock Mock) {
	mocks[getMockeId(mock.HttpMethod, mock.Url)] = &mock

}

func Post(url string, body interface{}, headers http.Header) (*http.Response, error) {
	if enableMocks {
		mock := mocks[getMockeId(http.MethodPost, url)]

		if mock == nil {
			return nil, errors.New("no mockup found fot give request")
		}
		//return local mock without calling any external ressource!
		return mock.Response, mock.Err
	}

	jsonByte, err := json.Marshal(body)

	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonByte))
	request.Header = headers
	client := http.Client{}
	return client.Do(request)
}
