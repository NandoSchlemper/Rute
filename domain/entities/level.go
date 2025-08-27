package entities

type ILevel interface {
	addXp(xp int) // Verifica se vai upar de level tbm
	up()          // Ajusta os atributos para o novo nível
	level() int
}

type Level struct {
	ActualXP int
	lvl      int
}

// Up implements IProgression.
func (p *Level) up() {
	p.lvl = p.lvl + 1
	p.ActualXP = 0
}

// AddXp implements IProgression.
func (p *Level) addXp(xp int) {
	if p.ActualXP+xp >= p.lvl*20 {
		p.up()
	}
	p.ActualXP = p.ActualXP + xp
}

func (p Level) level() int {
	return p.lvl
}

func NewLevel(lvl int) ILevel {
	return &Level{
		ActualXP: 0,
		lvl:      lvl,
	}
}
