package entities

type IDamage interface{}

type Damage struct {
	Base       int
	Attribute  Attribute
	Multiplier float64
	Type       DamageType
	Target     Target
}
