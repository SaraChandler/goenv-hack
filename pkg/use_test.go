package pkg

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckGoMod(t *testing.T) {
	testmod, err := os.Open("./testdata/gomod.test")
	assert.Nil(t, err, "Expect no error opening pkg/testdata/gomod.test")
	defer testmod.Close()

	version, err := checkGoMod(testmod)
	assert.Nil(t, err, "Expect no error from checkGoMod")
	assert.Equal(t, version, "1.17")
}
