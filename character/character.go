package character

type Character interface {
	Str() int
	Agi() int
	Int() int
	HitPoints() int
	Hit(int)
	AttackSpeed() int
	CriticalRate() int
	EvasionRate() int
	Damage() int
	Invariant() int
}
