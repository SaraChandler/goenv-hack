package pkg

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestisVersionInList(t *testing.T) {
  list := []string{"1.16.5", "1.16.7", "1.16.8"}
  lookupVersion := "1.16.8"

  inList, err := isVersionInList(lookupVersion, list)
  assert.Equal(t, inList, true)
  assert.Nil(t, err, "Expect no error for found in list")

  lookupVersion = "1.16.9"

  inList, err = isVersionInList(lookupVersion, list)
  assert.Equal(t, inList, false)
  assert.NotNil(t, err, "expect an error for version not in list")
}

func TestmakeDownloadURL(t *testing.T) {
  version := "1.16.6"
  os := "darwin"
  arch := "amd64"

  url := makeDownloadURL(version, os, arch)
  matchURL := "https://go.dev/dl/go1.16.6.darwin-amd64.tar.gz"

  assert.Equal(t, url, matchURL, "expected url should match output url")
}
