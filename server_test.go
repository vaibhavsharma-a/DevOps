package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func AllTest(t *testing.T, fnc http.HandlerFunc, code int, message string) {
	testServer := httptest.NewServer(http.HandlerFunc(fnc))
	defer testServer.Close()

	fmt.Println(testServer.URL)

	testClient := testServer.Client()
	respose, err := testClient.Get(testServer.URL)
	if err != nil {
		t.Error(err)
	}

	body, err := io.ReadAll(respose.Body)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, message, string(body))
	assert.Equal(t, code, respose.StatusCode)

}
func TestMainPage(t *testing.T) {
	AllTest(t, handleMainPage, 200, "This is the main page")
}

func TestHealth(t *testing.T) {
	AllTest(t, handleHealth, 200, "All Good!")

}

func TestNewEndpoint(t *testing.T) {
	AllTest(t, handleNewEndpoint, 200, "This is the new endpoint")
}
