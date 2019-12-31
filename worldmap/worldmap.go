package worldmap

import "github.com/evgeniylapta/go/village"

type WorldMap struct {
	villages []village.Village
	size int
}

type Cord struct {
	x int
	y int
}

func New(size int, villages []village.Village) WorldMap {
	return WorldMap{villages, size}
}

