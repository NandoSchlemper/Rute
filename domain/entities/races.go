package entities

type IRaceSystem interface {
	setName(name string)
	setSkills(skills []ISkill)          // struct p skills
	setAttributes(atr IAttributeSystem) // atributos da raça
	getName() string
	getSkills() []ISkill // criar uma struct componente pra skills
	getAttributes() IAttributeSystem
}

type RaceSystem struct {
	Name       string
	Skills     []ISkill // criar uma struct componente para skills overall
	Attributes IAttributeSystem
}

func NewRaceSystem(
	name string,
	skills []ISkill,
	atributes IAttributeSystem,
) IRaceSystem {
	return &RaceSystem{
		Name:       name,
		Skills:     skills,
		Attributes: atributes,
	}
}

// getAttributes implements IRaceSystem.
func (r *RaceSystem) getAttributes() IAttributeSystem {
	panic("unimplemented")
}

// getName implements IRaceSystem.
func (r *RaceSystem) getName() string {
	panic("unimplemented")
}

// getSkills implements IRaceSystem.
func (r *RaceSystem) getSkills() []ISkill {
	panic("unimplemented")
}

// setAttributes implements IRaceSystem.
func (r *RaceSystem) setAttributes(atr IAttributeSystem) {
	panic("unimplemented")
}

// setName implements IRaceSystem.
func (r *RaceSystem) setName(name string) {
	panic("unimplemented")
}

// setSkills implements IRaceSystem.
func (r *RaceSystem) setSkills(skills []ISkill) {
	panic("unimplemented")
}
