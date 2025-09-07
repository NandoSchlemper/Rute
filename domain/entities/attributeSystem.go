package entities

type IAttributeSystem interface {
	Get(atr Attribute) int
}

type AttributeSystem map[Attribute]int

// Get implements IAttributeSystem.
func (a AttributeSystem) Get(atr Attribute) int {
	return a[atr]
}

func NewAttributeSystem() IAttributeSystem {
	return &AttributeSystem{
		STRENGTH:     0,
		DEXTERITY:    0,
		CONSTITUTION: 0,
		INTELLIGENCE: 0,
		WISDOM:       0,
		CHARISMA:     0,
	}
}
