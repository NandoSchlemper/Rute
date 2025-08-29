package entities

type ISkill interface {
	Use()
	Description()
}

type Skill struct {
	Damage
	Name   string
	Origin SkillOrigin
}

// Description implements ISkill.
func (s *Skill) Description() {
	panic("unimplemented")
}

// Use implements ISkill.
func (s *Skill) Use() {
	panic("unimplemented")
}

func NewSkill(
	baseDamage int,
	attribute Attribute,
	multiplier float64,
	damageType DamageType,
	target Target,
	name string,
	origin SkillOrigin,
) ISkill {
	return &Skill{
		Damage: Damage{
			Base:       baseDamage,
			Attribute:  attribute,
			Type:       damageType,
			Multiplier: multiplier,
			Target:     target,
		},
		Name:   name,
		Origin: origin,
	}
}
