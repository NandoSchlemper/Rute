package entities

type IPlayer interface{}

type Player struct {
	name      string
	level     ILevel
	attribute IAttributeSystem
	race      IRaceSystem
	skill     ISkillSystem
}

func NewPlayer(
	playerName string,
	levelInterface ILevel,
	attributeInterface IAttributeSystem,
	raceInterface IRaceSystem,
	skillInterfcae ISkillSystem,
) IPlayer {
	return &Player{
		name:      playerName,
		level:     levelInterface,
		attribute: attributeInterface,
		race:      raceInterface,
		skill:     skillInterfcae,
	}
}
