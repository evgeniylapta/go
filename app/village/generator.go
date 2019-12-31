package village

import (
	"github.com/evgeniylapta/go/utils"
)

func worker(vilChan chan Village, length int) {
	for i := 0; i < length; i++ {
		const min, max = 0, 10000
		vilChan <- New(utils.RandomInt(min, max), utils.RandomInt(min, max))
	}
}

func Generate(length int) (villages []Village) {
	vilChan := make(chan Village, length)

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
