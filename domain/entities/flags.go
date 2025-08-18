package entities

type IFlag interface {
	SetFlag()
	UnsetFlag()
}

type Flag struct {
	Type        FlagType // Tipo da Flag
	Description string   // debugging
	Activate    bool     // Se está ativa ou não
}

func NewFlag(flagtype FlagType, desc string) IFlag {
	return &Flag{
		Type:        flagtype,
		Description: desc,
	}
}

func (f *Flag) String() string {
	return f.Description
}

func (f *Flag) SetFlag() {
	f.Activate = true
}

func (f *Flag) UnsetFlag() {
	f.Activate = false
}
