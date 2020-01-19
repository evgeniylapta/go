package legionary

type Legionary struct {
	power  int
	status bool
}

func (legionary Legionary) GetPower() int {
	return legionary.power
}

func New(power int, status bool) Legionary {
	return Legionary{power, status}
}
