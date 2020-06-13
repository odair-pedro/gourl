package fetch

import "testing"

func TestNewRequestParams(t *testing.T) {
	const url string = "url"
	const count int = 1
	params := NewRequestParams(url, count)

	if params.Url != url || params.Count != count {
		t.Fail()
	}
}
