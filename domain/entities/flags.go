package entities

type Flag struct {
	Type        FlagType // Tipo da Flag
	Description string   // debugging
	Activate    bool     // Se está ativa ou não
}

func (f *Flag) SetFlag() {
	f.Activate = true
}
