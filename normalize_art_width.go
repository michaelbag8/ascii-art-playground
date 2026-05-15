package main

import (
	"strings"
)

func NormalizeArtWidth(rows []string) []string {
	if len(rows) == 0 {
		return []string{}
	}

	result := make([]string, len(rows))

	maxL := len(rows[0])

	for _, value := range rows {
		if len(value) > maxL {
			maxL = len(value)
		}
	}

	for index, value := range rows {
		padding := maxL - len(value)
		result[index] = value + strings.Repeat(" ", padding)
	}

	return result
}
