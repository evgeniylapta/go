package utils

import (
	"os"
)

func ArgsWithoutProg() []string {
	return os.Args[1:]
}
