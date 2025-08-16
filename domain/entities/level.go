package entities

type IProgression interface {
	AddXp(xp int) // Verifica se vai upar de level tbm
	Up()          // Ajusta os atributos para o novo nível
}

type Progression struct {
	MaxXP    int
	ActualXP int
	Level    int
}

// Up implements IProgression.
func (p *Progression) Up() {
	p.Level = p.Level + 1
	p.MaxXP = p.Level * 20
	p.ActualXP = 0
}

// AddXp implements IProgression.
func (p *Progression) AddXp(xp int) {
	if p.ActualXP+xp >= p.MaxXP {
		p.Up()
	}
}

func NewLevel(lvl int) IProgression {
	return &Progression{
		MaxXP:    lvl * 20,
		ActualXP: 0,
		Level:    lvl,
	}
}
