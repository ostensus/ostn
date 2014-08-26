package entropy

import (
	"crypto/md5"
	"fmt"
	"io"
	"strings"
	"time"
)

type MerkleNode struct {
	Name    string
	Id      string
	Version string
	Child   *MerkleNode
}

func (node *MerkleNode) IsLeaf() bool {
	return node.Child == nil
}

func (node *MerkleNode) DescendencyPath() string {
	if node.IsLeaf() {
		return node.Name
	} else {
		return node.Name + "." + node.Child.DescendencyPath()
	}
}

func BuildHierarchy(path string) *MerkleNode {
	rootNode := &MerkleNode{}

	segments := strings.Split(path, ".")

	if len(segments) > 0 {
		var currentNode *MerkleNode = rootNode

		for _, segment := range segments {
			node := &MerkleNode{Name: segment}
			currentNode.Child = node
			currentNode = node
		}
	}

	return rootNode
}

func BuildEntityNode(id string, version string) *MerkleNode {

	// TODO validate that id and version are not nil

	hash := md5.New()
	io.WriteString(hash, id)
	digest := fmt.Sprintf("%x", hash.Sum(nil))

	leafId := digest[2:3]
	leafNode := MerkleNode{Name: leafId, Id: id, Version: version}

	midId := digest[1:2]
	midNode := MerkleNode{Name: midId, Child: &leafNode}

	rootId := digest[0:1]
	rootNode := MerkleNode{Name: rootId, Child: &midNode}

	return &rootNode
}

func BuildDateTimeNode(id string, version string, d time.Time) *MerkleNode {

	leafName := d.Format(dayFormat)
	leafNode := MerkleNode{Name: leafName, Id: id, Version: version}

	midName := d.Format(monthFormat)
	midNode := MerkleNode{Name: midName, Child: &leafNode}

	rootName := d.Format(yearFormat)
	rootNode := MerkleNode{Name: rootName, Child: &midNode}

	return &rootNode
}
