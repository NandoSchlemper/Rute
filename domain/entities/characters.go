package entities

type Character struct {
	Name        string
	Race        Race
	Level       int
	Age         int
	Relatioship int
}

func NewCharacter() *Character {
	return &Character{}
}
