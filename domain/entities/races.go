package entities

type IRaceSystem interface {
	setName(name string)
	setSkills(skills []ISkillSystem)   // struct p skills
	setAttributes(atr AttributeSystem) // atributos da raça
	getName() string
	getSkills() []ISkillSystem // criar uma struct componente pra skills
	getAttributes() AttributeSystem
}

type RaceSystem struct {
	Name       string
	Skills     []ISkillSystem // criar uma struct componente para skills overall
	Attributes AttributeSystem
}

// getAttributes implements IRaceSystem.
func (r *RaceSystem) getAttributes() AttributeSystem {
	panic("unimplemented")
}

// getName implements IRaceSystem.
func (r *RaceSystem) getName() string {
	panic("unimplemented")
}

// getSkills implements IRaceSystem.
func (r *RaceSystem) getSkills() []ISkillSystem {
	panic("unimplemented")
}

// setAttributes implements IRaceSystem.
func (r *RaceSystem) setAttributes(atr AttributeSystem) {
	panic("unimplemented")
}

// setName implements IRaceSystem.
func (r *RaceSystem) setName(name string) {
	panic("unimplemented")
}

// setSkills implements IRaceSystem.
func (r *RaceSystem) setSkills(skills []ISkillSystem) {
	panic("unimplemented")
}
