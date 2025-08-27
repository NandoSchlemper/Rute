package factory

import "Rute/domain/entities"

func NewRaceSystem(
	name string,
	skills []entities.ISkillSystem,
	attributes entities.AttributeSystem,
) entities.IRaceSystem {
	return &entities.RaceSystem{
		Name:       name,
		Skills:     skills,
		Attributes: attributes,
	}
}
