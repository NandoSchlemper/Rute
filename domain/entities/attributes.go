package entities

type Attribute int

const (
	STRENGTH Attribute = iota
	DEXTERITY
	CONSTITUTION
	INTELLIGENCE
	WISDOM
	CHARISMA
)

type IAttributeSystem interface {
	update(upd AttributeSystem) error
}

type AttributeSystem map[Attribute]int

// update implements IAttributeSystem.
func (a *AttributeSystem) update(upd AttributeSystem) error {
	for idx := range *a {
		(*a)[idx] = (*a)[idx] + upd[idx]
	}
	return nil
}

func NewAttributeSystem(
	str int,
	dex int,
	con int,
	ine int,
	wis int,
	cha int,
) IAttributeSystem {
	return &AttributeSystem{
		STRENGTH:     str,
		DEXTERITY:    dex,
		CONSTITUTION: con,
		INTELLIGENCE: ine,
		WISDOM:       wis,
		CHARISMA:     cha,
	}
}
