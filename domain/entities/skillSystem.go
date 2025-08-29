package entities

type ISkillSystem interface {
	Learn(skills []Skill) error
}

type SkillSystem struct {
	Learned  []Skill
	Equipped []Skill // actives + passives
}

// Learn implements ISkillSystem.
func (s *SkillSystem) Learn(skills []Skill) error {
	s.Learned = append(s.Learned, skills...)
	return nil
}

func NewSkillSystem() ISkillSystem {
	return &SkillSystem{
		Learned:  make([]Skill, 0),
		Equipped: make([]Skill, 0, 5),
	}
}
