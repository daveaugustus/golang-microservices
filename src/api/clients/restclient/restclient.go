package restclient

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func Post(url string, body interface{}, headers http.Header) (*http.Response, error) {

	jsonBytes, err := json.Marshal(body)
	if err != nil {
		// TODO: Return error
		return nil, err
	}
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonBytes))
	if err != nil {
		//TODO: handle error
	}
	request.Header = headers

	client := http.Client{}
	response, err := client.Do(request)
	return response, err
}
