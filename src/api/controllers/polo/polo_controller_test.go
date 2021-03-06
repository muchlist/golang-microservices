package polo

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/muchlist/golang-microservices/src/api/utils/test_utils"
	"github.com/stretchr/testify/assert"
)

func TestConstants(t *testing.T) {
	assert.EqualValues(t, "polo", polo)
}

func TestMarco(t *testing.T) {
	response := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/marco", nil)
	c := test_utils.GetMockedContext(request, response)

	Marco(c)

	assert.EqualValues(t, http.StatusOK, response.Code)
	assert.EqualValues(t, "\"polo\"", response.Body.String())
}
