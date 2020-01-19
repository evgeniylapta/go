package utils

import "sort"

func StringFound(strings []string, str string) bool {
	return sort.SearchStrings(strings, str) != len(strings)
}
