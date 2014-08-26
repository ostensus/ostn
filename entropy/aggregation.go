package entropy

type Aggregation interface {
	AttributeName() string
	Parent() string
	Bucket(string) string
}

type DateGranularity uint

const (
	Yearly DateGranularity = iota
	Monthly
	Daily
)

type DateAggregation struct{}

func (d *DateAggregation) AttributeName() string {
	return ""
}

func (d *DateAggregation) Parent() string {
	return ""
}

func (d *DateAggregation) Bucket(value string) string {
	return value
}

func (d *DateAggregation) Granularity() DateGranularity {
	return Yearly
}
