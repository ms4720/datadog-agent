package ec2

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsDefaultHostname(t *testing.T) {
	assert.True(t, IsDefaultHostname("IP-FOO"))
	assert.True(t, IsDefaultHostname("domuarigato"))
	assert.False(t, IsDefaultHostname(""))
}

func TestGetInstanceID(t *testing.T) {
	expected := "i-0123456789abcdef0"
	var lastRequest *http.Request
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		io.WriteString(w, expected)
		lastRequest = r
	}))
	defer ts.Close()
	metadataURL = ts.URL

	val, err := GetInstanceID()
	assert.Nil(t, err)
	assert.Equal(t, expected, val)
	assert.Equal(t, lastRequest.URL.Path, "/instance-id")
}