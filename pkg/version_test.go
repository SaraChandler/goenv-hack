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
	assert.Equal(t, compareSame, expectSame, "expect return 0 for same version number")
	assert.Nil(t, err, "expect no error for compared versions")

	versionA = "1.16.5"
	versionB = "1.16.8"

	compareLess, err := CompareVersions(versionA, versionB)
	expectLess := -1
	assert.Equal(t, compareLess, expectLess, "expect return -1 for A < B")
	assert.Nil(t, err, "expect no error for compared versions")

	versionA = "1.16.8"
	versionB = "1.16.5"

	compareGreater, err := CompareVersions(versionA, versionB)
	expectGreater := 1
	assert.Equal(t, compareGreater, expectGreater, "expect return 1 for A > B")
	assert.Nil(t, err, "expect no error for compared versions")
}

func TestfuzzyCompare(t *testing.T) {
	versionA := "1.16"
	versionB := "1.16.9"

	wantTrue := fuzzyCompare(versionA, versionB)
	gotTrue := true

	assert.Equal(t, wantTrue, gotTrue, "expect fuzzy match to be true")

	versionA = "1.16.9"
	versionB = "1.15"

	wantFalse := fuzzyCompare(versionA, versionB)
	gotFalse := false

	assert.Equal(t, wantFalse, gotFalse, "expect fuzzy match to be false")

	versionA = "1.14"
	versionB = "1.15.6"

	wantFalse = fuzzyCompare(versionA, versionB)

	assert.Equal(t, wantFalse, gotFalse, "expect fuzzy match to be false")

	versionA = "1.1"
	versionB = "1.15.6"

	wantFalse = fuzzyCompare(versionA, versionB)

	assert.Equal(t, wantFalse, gotFalse, "expect fuzzy match to be false")
}
