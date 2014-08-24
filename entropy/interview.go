package entropy

type Constraint interface {
	AttributeName() string
}

type Aggregation interface {
	AttributeName() string
	Parent() string
	Bucket(string) string
}

type Question interface {
	Constraints() []Constraint
	Aggregations() []Aggregation
	SliceThreshold() int
}

type Iter struct {
	err error
}

func (iter *Iter) Next() Question {
	return nil
}

func NextQuestion(repository int64) *Iter {
	return &Iter{}
}
