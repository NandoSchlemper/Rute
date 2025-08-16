package entities

type FlagType int

const (
	CharacterFlagType FlagType = iota
	LocalFlagType
	GlobalFlagType
)

func (f FlagType) String() string {
	return [...]string{"Character", "Local", "Global"}[f]
}
