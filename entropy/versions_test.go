package entropy

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type VersionTestSuite struct {
	suite.Suite
	store *VersionStore
}

func TestVersionTestSuite(t *testing.T) {
	suite.Run(t, new(VersionTestSuite))
}

func (s *VersionTestSuite) SetupTest() {
	v, err := OpenStore()
	if err != nil {
		s.T().Fatalf("Could not open store: %s", err)
	}

	s.store = v

	id := "foo"
	version := "v3"
	attributeName := "z"
	d := time.Now()

	ev := NewDatePartitionedEvent(id, version, attributeName, d)
	err = v.Accept(ev)

	assert.Nil(s.T(), err)
}

func (s *VersionTestSuite) TestVersionStoreSync() {
	assert.Equal(s.T(), s.store.SliceThreshold(), 128)
}
