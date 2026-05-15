package main

import "strings"

func PadArtRows(rows []string, width int) []string {
	lines := make([]string, len(rows))

	for index, values := range rows {
		padding := width - len(values)
		if padding > 0 {
			lines[index] = values + strings.Repeat(" ", padding)
		} else {
			lines[index] = values
		}

	}
	return lines
}
