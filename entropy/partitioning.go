package entropy

type DataType uint

const (
	DateTime DataType = iota
)

type RangePartitionDescriptor struct {
	DataType DataType
}
