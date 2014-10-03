package entropy

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"os"
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
	path := "versions.db"
	os.Remove(path)
	v, err := OpenStore(path)
	if err != nil {
		s.T().Fatalf("Could not open store: %s", err)
	}
	s.store = v
}

func (s *VersionTestSuite) TestNewRepository() {

	attributeName := "ts_col"

	parts := make(map[string]PartitionDescriptor)
	parts[attributeName] = &RangePartitionDescriptor{}
	parts["x"] = &SetPartitionDescriptor{Values: []string{"y"}}

	repo, err := s.store.NewRepository("some_src", "some_repo", parts)
	assert.NoError(s.T(), err)
	assert.True(s.T(), repo > 0)

	_, err = s.store.NewRepository("some_src", "some_repo", parts)
	assert.Error(s.T(), err)

	id := "foo"
	version := "v3"

	atts := make(map[string]interface{})
	atts[attributeName] = time.Now()
	atts["x"] = "y"

	ev := NewPartitionedEvent(id, version, atts)
	err = s.store.Accept(repo, ev)

	assert.NoError(s.T(), err)

	digests, err := s.store.Digest(repo)
	assert.NoError(s.T(), err)

	// $ echo -n "v3" | md5
	// 43a03299a3c3fed3d8ce7b820f3aca81

	assert.Equal(s.T(), "43a03299a3c3fed3d8ce7b820f3aca81", digests[id])
}

// func (s *VersionTestSuite) TestVersionStoreSync() {
// 	assert.Equal(s.T(), s.store.SliceThreshold(), 128)
// }
