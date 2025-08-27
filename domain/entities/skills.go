package entities

type SkillType int

const (
	PASSIVE SkillType = iota
	ACTIVE
)

type ISkillSystem interface {
	getName() string
	getType() SkillType
	setName(name string)
	setType(skill SkillType)
}

type SkillSystem struct {
	Name string
	Type SkillType
}
