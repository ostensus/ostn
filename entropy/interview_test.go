package entropy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmptyInterview(t *testing.T) {
	iter := NextQuestion(100)
	assert.Nil(t, iter.err)
	assert.Nil(t, iter.Next())
}
