package lib

type BuffTypeEnum uint

//noinspection ALL
const (
	HealHP BuffTypeEnum = iota + 1
	HealMP
	Defense
	Damage
	Magic
	Confused
	Blindness
	Mute
	Burned
	Freezing
	Tired
	ElementChance
	Poisoned
	Electrocuted
	Slow
	Speed
	ElementDamage
	Critical
	Dodge
	MagicDefense
	StunInmune
	CancelPoisoned
	CriticalDamage
	Hide
	Nature
	SwitchMagicDefense
	SwitchPhysicalDefense
	ReduceDamage
	ReduceMagic
	MaxHP
	MaxMP
	Count
)
