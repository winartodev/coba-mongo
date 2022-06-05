package response_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"winartodev/coba-mongodb/response"
)

func TestHttpResponseSuccess(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rr := httptest.NewRecorder()
	response.HttpResponseSuccess(rr, req, "success")
}

func TestHttpResponseFailed(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rr := httptest.NewRecorder()
	response.HttpResponseFailed(rr, req, http.StatusInternalServerError, errors.New("failed"))
}
