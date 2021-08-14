package character

type Warrior struct {
	HP        int
	strPoint  int
	agiPoint  int
	dmgPoint  int
	intPoint1 int
	intPoint2 int
}

func NewWarrior(strPoint, agiPoint, dmgPoint, intPoint1, intPoint2 int) *Warrior {
	HP := WarriorStr * strPoint

	return &Warrior{
		HP:        HP,
		strPoint:  strPoint,
		agiPoint:  agiPoint,
		dmgPoint:  dmgPoint,
		intPoint1: intPoint1,
		intPoint2: intPoint2,
	}
}

func (w Warrior) Str() int {
	return WarriorStr
}

func (w Warrior) Agi() int {
	return WarriorAgi
}

func (w Warrior) Int() int {
	return WarriorInt
}

func (w Warrior) HitPoints() int {
	return w.HP
}

func (w *Warrior) Hit(dmg int) {
	if w.HP < dmg {
		w.HP = 0
	} else {
		w.HP -= dmg
	}
}

func (w Warrior) AttackSpeed() int {
	return w.Agi() * w.agiPoint
}

func (w Warrior) CriticalRate() int {
	return w.Int() * w.intPoint1
}

func (w Warrior) EvasionRate() int {
	return w.Int() * w.intPoint2
}

func (w Warrior) Damage() int {
	return w.Str() * w.dmgPoint
}

func (w Warrior) Invariant() int {
	res := w.strPoint * w.Str()
	res *= w.AttackSpeed()
	res *= 100 + w.CriticalRate()

	return res * w.Damage()
}
