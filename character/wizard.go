package character

type Wizard struct {
	HP        int
	strPoint  int
	agiPoint  int
	dmgPoint  int
	intPoint1 int
	intPoint2 int
}

func NewWizard(strPoint, agiPoint, dmgPoint, intPoint1, intPoint2 int) *Wizard {
	HP := WizardStr * strPoint

	return &Wizard{
		HP:        HP,
		strPoint:  strPoint,
		agiPoint:  agiPoint,
		dmgPoint:  dmgPoint,
		intPoint1: intPoint1,
		intPoint2: intPoint2,
	}
}

func (w Wizard) Str() int {
	return WizardStr
}

func (w Wizard) Agi() int {
	return WizardAgi
}

func (w Wizard) Int() int {
	return WizardInt
}

func (w Wizard) HitPoints() int {
	return w.HP
}

func (w *Wizard) Hit(dmg int) {
	if w.HP < dmg {
		w.HP = 0
	} else {
		w.HP -= dmg
	}
}

func (w Wizard) AttackSpeed() int {
	return w.Agi() * w.agiPoint
}

func (w Wizard) CriticalRate() int {
	return w.Int() * w.intPoint1
}

func (w Wizard) EvasionRate() int {
	return w.Int() * w.intPoint2
}

func (w Wizard) Damage() int {
	return w.Int() * w.dmgPoint
}

func (w Wizard) Invariant() int {
	res := w.strPoint * w.Str()
	res *= w.AttackSpeed()
	res *= 100 + w.CriticalRate()

	return res * w.Damage()
}