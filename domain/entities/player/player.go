package player

import "Rute/domain/entities/attributes"

type IPlayer interface {
}

type Player struct {
	level     ILevel
	attribute attributes.IAttributeSystem
}

func NewPlayer(
	levelInterface ILevel,
	attributeInterface attributes.IAttributeSystem,
) IPlayer {
	return &Player{
		level:     levelInterface,
		attribute: attributeInterface,
	}
}
