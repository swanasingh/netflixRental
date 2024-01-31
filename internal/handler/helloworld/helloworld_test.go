package helloworld

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"netflixRental/internal/models"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	engine := gin.Default()

	engine.GET("/netflix/api/hello", HelloWorld)

	request, err := http.NewRequest(http.MethodGet, "/netflix/api/hello", nil)
	require.NoError(t, err)

	response := httptest.NewRecorder()
	engine.ServeHTTP(response, request)

	var actualResponse models.HelloWorld
	err = json.NewDecoder(response.Body).Decode(&actualResponse)
	require.NoError(t, err)

	expected := "Hello World!"

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, actualResponse.Message, expected)
}

func TestFail(t *testing.T) {
	fmt.Println("still trying sadly")
	t.Fail()
}
