package firsttry_test

import (
	"firstTry/go-try-1/firsttry"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTry(t *testing.T) {
	output := firsttry.Try()
	assert.Equal(t, output, "Si")
}
