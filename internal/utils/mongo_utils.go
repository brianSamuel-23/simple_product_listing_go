package utils

import (
	"strings"
)

func ConvertSort(sortOrder string) int8 {
	if strings.ToLower(sortOrder) == "asc" {
		return 1
	} else {
		return -1
	}
}
