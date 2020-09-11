package restclient

import (
	"bytes"
	"encoding/json"
	"net/http"
)


var (
	enabledMocks = false
	mocks        = make(map[string]*response)
)

type response struct{}

func StartMockups() {
	enabledMocks = true
}
func StopMockups() {
	enabledMocks = false
}

func Post(url string, body interface{}, headers http.Header) (*http.Response, error) {
	if enabledMocks {
		// TODO: Retrun local mock without calling any external resource
	}

	jsonBytes, err := json.Marshal(body)
	if err != nil {
		// TODO: Return error
		return nil, err
	}
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonBytes))
	if err != nil {
		//TODO: handle error
		return nil, err

	}
	request.Header = headers

	client := http.Client{}
	response, err := client.Do(request)
	return response, err
}
