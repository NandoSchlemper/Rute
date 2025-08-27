package factory

import "Rute/domain/entities"

func NewSkillSystem(
	name string,
	skillType entities.SkillType,
) entities.ISkillSystem {
	return &entities.SkillSystem{
		Name: name,
		Type: skillType,
	}
}
