package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestisVersionInList(t *testing.T) {
	list := []string{"1.16.5", "1.16.7", "1.16.8"}
	lookupVersion := "1.16.8"

	inList, listed := isVersionInList(lookupVersion, list)
	assert.Equal(t, inList, true)
	assert.Equal(t, listed, "1.16.8")

	lookupVersion = "1.16.9"

	inList, listed = isVersionInList(lookupVersion, list)
	assert.Equal(t, inList, false)
	assert.Equal(t, listed, "")

	lookupVersion = "1.16"

	inList, listed = isVersionInList(lookupVersion, list)
	assert.Equal(t, inList, true)
	assert.Equal(t, listed, "1.16.8")
}

func TestmakeDownloadURL(t *testing.T) {
	version := "1.16.6"
	os := "darwin"
	arch := "amd64"

	url := makeDownloadURL(version, os, arch)
	matchURL := "https://go.dev/dl/go1.16.6.darwin-amd64.tar.gz"

	assert.Equal(t, url, matchURL, "expected url should match output url")
}
