package entropy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInterview(t *testing.T) {
	iter := Commence(200)
	assert.Nil(t, iter.err)
	assert.Nil(t, iter.Next())
}
