package village

type Village struct {
	x int
	y int
}

func New (x, y int) Village {
	return Village{x, y}
}
