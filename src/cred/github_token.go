package cred

import (
	"io/ioutil"
	"os"
	"strings"
)

func GetGithubToken() string {
	// Read the following file
	file, err := os.Open("/home/dave/Github/.token")
	if err != nil {
		// TODO: HANDLE ERROR
	}

	// File to bytes
	tokenBytes, err := ioutil.ReadAll(file)
	if err != nil {
		// TODO: HANDLE ERROR
	}
	// Trim "\n" at the end of the string
	token := strings.TrimSuffix(string(tokenBytes), "\n")

	return token
}
