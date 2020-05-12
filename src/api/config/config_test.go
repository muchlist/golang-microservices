package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstant(t *testing.T) {
	assert.EqualValues(t, "SECRET_GITHUB_ACCESS_TOKEN", apiGithubAccessToken)
}

func TestGetGithubAccessToken(t *testing.T) {
	token := GetGithubAccessToken()
	assert.EqualValues(t, "", token) //string kosong karena environment tidak di setting
}
