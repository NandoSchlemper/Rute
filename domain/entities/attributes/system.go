package attributes

type IAttributeSystem interface{}

type AttributeSystem map[Attribute]int

func NewAttributSystem() IAttributeSystem {
	return &AttributeSystem{
		FORÇA:        0,
		VELOCIDADE:   0,
		CONSTITUIÇÃO: 0,
		INTELIGENCIA: 0,
		SABEDORIA:    0,
		CARISMA:      0,
	}
}
