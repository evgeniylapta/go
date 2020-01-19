package village

import "github.com/jinzhu/gorm"

type Village struct {
	gorm.Model
	X int
	Y int
}

func New(x, y int) Village {
	return Village{X: x, Y: y}
}
