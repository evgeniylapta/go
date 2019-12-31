package archer

type Archer struct {
	power int
}

func (archer Archer) GetPower() int {
	return archer.power
}

func New(power int) Archer {
	return Archer{power}
}
