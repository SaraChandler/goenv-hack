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
