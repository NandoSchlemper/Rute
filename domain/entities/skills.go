package entities

type ISkill interface {
	// criar uma struct para os enemagos kekeke
	Use(caster IPlayer, target int) error
	Description()
	Cost() int
	CanCast() bool
}

type Skill struct {
	Damage
	Name   string
	Origin SkillOrigin
	Mp     int
}

// CanCast implements ISkill.
func (s *Skill) CanCast() bool {
	panic("unimplemented")
}

// Cost implements ISkill.
func (s *Skill) Cost() int {
	panic("unimplemented")
}

// Description implements ISkill.
func (s *Skill) Description() {
	panic("unimplemented")
}

// Use implements ISkill.
func (s *Skill) Use(caster IPlayer, target int) error {
	panic("unimplemented")
}

func NewSkill(
	baseDamage int,
	mpCost int,
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
		Mp:     mpCost,
	}
}
