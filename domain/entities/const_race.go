package entities

type Race int

const (
	Human Race = iota
)

func (r Race) String() string {
	return [...]string{"Human"}[r]
}
