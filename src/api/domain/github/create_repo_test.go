package github

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateRepoRequestAsJson(t *testing.T) {
	request := CreateRepoRequest{
		Name:        "Micro service golang",
		Description: "string",
		Homepage:    "string",
		Private:     true,
		HasIssues:   false,
		HasProject:  true,
		HasWiki:     false,
	}

	//Marshal takes an input interface and attemp to create a valid json string
	bytes, err := json.Marshal(request)
	assert.Nil(t, err)
	assert.NotNil(t, bytes)

	fmt.Println(string(bytes))

	var target CreateRepoRequest

	//Unmarshal takes an input byte array and a *pointer* that we're trying ti fill using this json.
	err = json.Unmarshal(bytes, &target)
	assert.Nil(t, err)

	assert.EqualValues(t, target.Name, "Micro service golang")
}
