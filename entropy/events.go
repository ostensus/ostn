package entropy

import (
	"time"
)

type ChangeEvent interface {
	Id() string
}

type TombstoneEvent interface {
	ChangeEvent
}

type UpsertEvent interface {
	ChangeEvent
	Version() string
}

type UnpartitionedEvent interface {
	UpsertEvent
	IdHierarchy() *MerkleNode
}

type PartitionedEvent interface {
	UnpartitionedEvent
	Attributes() map[string]string
	AttributeHierarchy() *MerkleNode
}

type baseChangeEvent struct {
	id string
}

type baseUpsertEvent struct {
	baseChangeEvent
	version string
}

type baseUnpartitionedEvent struct {
	baseUpsertEvent
	idHierarchy *MerkleNode
}

type basePartitionedEvent struct {
	baseUnpartitionedEvent
	attributeHierarchy *MerkleNode
	attributes         map[string]string
}

func (d *basePartitionedEvent) Id() string {
	return d.id
}

func (d *basePartitionedEvent) Version() string {
	return d.version
}

func (d *basePartitionedEvent) IdHierarchy() *MerkleNode {
	return d.idHierarchy
}

func (d *basePartitionedEvent) Attributes() map[string]string {
	return d.attributes
}

func (d *basePartitionedEvent) AttributeHierarchy() *MerkleNode {
	return d.attributeHierarchy
}

const (
	dateFormat  = "20060102150405"
	yearFormat  = "2006"
	monthFormat = "01"
	dayFormat   = "02"
)

func NewPartitionedEvent(id, version string, attributes map[string]interface{}) PartitionedEvent {
	atts := make(map[string]string)

	for name, value := range attributes {
		switch v := value.(type) {
		case time.Time:
			atts[name] = v.Format(dateFormat)
		case string:
			atts[name] = v
		}
	}

	ev := &basePartitionedEvent{}
	ev.attributes = atts
	ev.id = id
	ev.version = version
	ev.idHierarchy = BuildEntityNode(id, version)
	//ev.attributeHierarchy = BuildDateTimeNode(id, version, d)

	return ev
}

func NewDatePartitionedEvent(id, version, attributeName string, d time.Time) PartitionedEvent {

	atts := make(map[string]string)
	atts[attributeName] = d.Format(dateFormat)

	ev := &basePartitionedEvent{}
	ev.attributes = atts
	ev.id = id
	ev.version = version
	ev.idHierarchy = BuildEntityNode(id, version)
	ev.attributeHierarchy = BuildDateTimeNode(id, version, d)

	return ev
}
