package village

import (
	"github.com/evgeniylapta/go/utils"
)

func isUniqueVillage(villages []Village, vil Village) bool {
	for _, curVil := range villages {
		if curVil.Y == vil.Y && curVil.X == vil.X {
			return false
		}
	}

	return true
}

func worker(vilChan chan Village, length int, minCord, maxCord int) {
	vils := []Village{}

	for i := 0; i < length; i++ {
		var isUnique bool

		for !isUnique {
			vil := New(utils.RandomInt(minCord, maxCord), utils.RandomInt(minCord, maxCord))

			if isUnique = isUniqueVillage(vils, vil); isUnique {
				vilChan <- vil
				vils = append(vils, vil)
			}
		}
	}
}

func GenerateVillages(length int, minCord, maxCord int) (villages []Village) {
	vilChan := make(chan Village, length)

	go worker(vilChan, length, minCord, maxCord)

	vilCount := 0
	for {
		select {
		case vil := <-vilChan:
			villages = append(villages, vil)
			vilCount++
		}

		if vilCount >= length {
			close(vilChan)
			return
		}
	}
}
