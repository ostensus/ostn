package entropy

type Constraint interface {
	AttributeName() string
}

type Question interface {
	Constraints() []Constraint
	Aggregations() []Aggregation
	SliceThreshold() int
}

type Answer struct {
}

type Pruner struct {
}

func (p *Pruner) Prune(a Answer) {
}

func (p *Pruner) Finish() {
}

type Iter struct {
	err error
}

func (iter *Iter) Next() Question {
	return nil
}

func Commence(repository int64) *Iter {
	return &Iter{}
}

func NextQuestion(repository int64, cons []Constraint, aggs []Aggregation, p *Pruner) *Iter {
	return &Iter{}
}
