package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateVersion(t *testing.T) {
	validVersion := "1.16.5"
	invalidVersion := "bogus"

	errValid := ValidateVersion(validVersion)
	assert.Nil(t, errValid, "expect no error for valid version")

	errInvalid := ValidateVersion(invalidVersion)
	assert.NotNil(t, errInvalid, "expect an error for invalid version")
}

func TestCompareVersions(t *testing.T) {
	versionA := "1.16.5"
	versionB := "1.16.5"

	compareSame, err := CompareVersions(versionA, versionB)
	expectSame := 0
	assert.Equal(t, compareSame, expectSame, "expect no error for valid version")
	assert.Nil(t, err, "expect no error for compared versions")

}
