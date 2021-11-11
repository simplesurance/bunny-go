package bunny

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCheckRespWithEmptyUnsuccessfulResp(t *testing.T) {
	req, err := http.NewRequest("get", "http://test.de", nil)
	require.NoError(t, err)

	resp := http.Response{
		StatusCode: 400,
		Body:       io.NopCloser(strings.NewReader("")),
	}

	err = checkResp(req, &resp)
	require.Error(t, err)
	require.IsType(t, &HTTPError{}, err)

	httpErr := err.(*HTTPError)
	assert.Empty(t, httpErr.Errors)
}
