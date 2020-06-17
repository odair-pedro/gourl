package fetch

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func TestNewRequestParams_Url_ShouldBeEqual(t *testing.T) {
	const url string = "url"
	const count int = 1
	params := NewRequestParams(url, count)
	assert.Equal(t, url, params.Url)
}

func TestNewRequestParams_Count_ShouldBeEqual(t *testing.T) {
	const url string = "url"
	const count int = 1
	params := NewRequestParams(url, count)
	assert.Equal(t, count, params.Count)
}

func TestCloseHttpResponse_ShouldCloseBeTrue(t *testing.T) {
	bodyMock := new(BodyMock)
	response := &http.Response{Body: bodyMock}
	closeHttpResponse(response)
	assert.True(t, response.Close)
}

func TestCloseHttpResponse_WithNilBody_ShouldNotPanic(t *testing.T) {
	require.NotPanics(t, func() {
		closeHttpResponse(new(http.Response))
	})
}

type BodyMock struct {
	mock.Mock
}

func (body BodyMock) Close() error {
	return nil
}

func (body BodyMock) Read(p []byte) (n int, err error) {
	return 0, nil
}
