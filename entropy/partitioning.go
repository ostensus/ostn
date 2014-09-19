package entropy

type DataType uint

const (
	DateTime DataType = iota
)

type PartitionDescriptor interface {
	DoesAggregate() bool
}

type RangePartitionDescriptor struct {
	DataType DataType
}

func (d *RangePartitionDescriptor) DoesAggregate() bool {
	return true
}

type SetPartitionDescriptor struct {
	Values []string
}

func (d *SetPartitionDescriptor) DoesAggregate() bool {
	return true
}
