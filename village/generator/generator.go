package generator

import (
	"github.com/evgeniylapta/go/utils"
	"github.com/evgeniylapta/go/village"
)

func worker(vilChan chan village.Village, length int) {
	for i := 0 ;i < length; i++ {
		const min, max = 0, 10000
		vilChan <- village.New(utils.RandomInt(min, max), utils.RandomInt(min, max))
	}
}

func Generate(length int) (villages []village.Village) {
	vilChan := make(chan village.Village, length)

	go worker(vilChan, length)

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
