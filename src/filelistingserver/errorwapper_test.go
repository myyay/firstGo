package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func errPanic(writer http.ResponseWriter, request *http.Request) error {
	panic(123)
}

type testingUserError string

func (e testingUserError) Error() string {
	return e.Message()
}

func (e testingUserError) Message() string {
	return string(e)
}

func errUserError(writer http.ResponseWriter, request *http.Request) error {
	return testingUserError("user error")
}

func errNotFound(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrNotExist
}
func errNoPermission(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrPermission
}
func errUnknown(writer http.ResponseWriter, request *http.Request) error {
	return errors.New("unknown error")
}

func errNoError(writer http.ResponseWriter, request *http.Request) error {
	b, _ := fmt.Fprintln(writer, "no error")
	fmt.Println(b)
	return nil
}

var tests = []struct {
	h       appHandler
	code    int
	message string
}{
	{errPanic, 500, "Internal Server Error"},
	{errUserError, 400, "user error"},
	{errNotFound, 404, "Not Found"},
	{errNoPermission, 403, "Forbidden"},
	{errUnknown, 500, "Internal Server Error"},
	{errNoError, 200, "no error"},
}

func TestErrorWrapper(t *testing.T) {

	for _, tt := range tests {
		f := errWrapper(tt.h)
		response := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "http://www.baidu.com", nil)

		f(response, request)

		verifyResponse(response.Result(), tt.code, tt.message, t)

	}
}

func TestErrorWrapperInServer(t *testing.T) {

	for _, tt := range tests {
		f := errWrapper(tt.h)
		server := httptest.NewServer(http.HandlerFunc(f))

		fmt.Println(server.URL)
		resp, _ := http.Get(server.URL)

		verifyResponse(resp, tt.code, tt.message, t)

	}
}

func verifyResponse(resp *http.Response, expectedCode int, expectedMsg string, t *testing.T) {
	b, _ := ioutil.ReadAll(resp.Body)
	body := strings.Trim(string(b), "\n")
	if resp.StatusCode != expectedCode || body != expectedMsg {
		t.Errorf("got (%d, %s); expected (%d, %s);", resp.StatusCode, body, expectedCode, expectedMsg)
	}
}
