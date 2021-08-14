package character

type Assassin struct {
	HP        int
	strPoint  int
	agiPoint  int
	dmgPoint  int
	intPoint1 int
	intPoint2 int
}

func NewAssassin(strPoint, agiPoint, dmgPoint, intPoint1, intPoint2 int) *Assassin {
	HP := AssassinStr * strPoint

	return &Assassin{
		HP:        HP,
		strPoint:  strPoint,
		agiPoint:  agiPoint,
		dmgPoint:  dmgPoint,
		intPoint1: intPoint1,
		intPoint2: intPoint2,
	}
}

func (a Assassin) Str() int {
	return AssassinStr
}

func (a Assassin) Agi() int {
	return AssassinAgi
}

func (a Assassin) Int() int {
	return AssassinInt
}

func (a Assassin) HitPoints() int {
	return a.HP
}

func (a *Assassin) Hit(dmg int) {
	if a.HP < dmg {
		a.HP = 0
	} else {
		a.HP -= dmg
	}
}

func (a Assassin) AttackSpeed() int {
	return a.Agi() * a.agiPoint
}

func (a Assassin) CriticalRate() int {
	return a.Int() * a.intPoint1
}

func (a Assassin) EvasionRate() int {
	return a.Int() * a.intPoint2
}

func (a Assassin) Damage() int {
	return a.Agi() * a.dmgPoint
}

func (a Assassin) Invariant() int {
	res := a.strPoint * a.Str()
	res *= a.AttackSpeed()
	res *= 100 + a.CriticalRate()

	return res * a.Damage()
}
