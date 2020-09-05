package github

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateRepoRequestAsJson(t *testing.T) {
	request := CreateRepoRequest{
		Name:        "golang introduction",
		Description: "Github APIs are being tested, please ignore",
		Homepage:    "https://github.com",
		Private:     true,
		HasIssues:   true,
		HasProjects: true,
		HasWiki:     true,
	}

	// Marshal takes an input interface and  attempts to create a valid json string
	bytes, err := json.Marshal(request)
	assert.Nil(t, err)
	assert.NotNil(t, bytes)
	assert.EqualValues(t, `{"name":"golang introduction","description":"Github APIs are being tested, please ignore","homepage":"https://github.com","private":true,"has_issues":true,"has_projects":true,"has_wiki":true}`, string(bytes))

	// Unmarshal takes an input byte array and a *pointer that we are trying to fill using this json
	var target CreateRepoRequest
	err = json.Unmarshal(bytes, &target)
	assert.Nil(t, err)

	assert.EqualValues(t, target.Name, request.Name)
	assert.EqualValues(t, target.Description, request.Description)
	assert.EqualValues(t, target.Homepage, request.Homepage)
	assert.EqualValues(t, target.Private, request.Private)
	assert.EqualValues(t, target.HasIssues, request.HasIssues)
	assert.EqualValues(t, target.HasProjects, request.HasProjects)
	assert.EqualValues(t, target.HasWiki, request.HasWiki)

}
