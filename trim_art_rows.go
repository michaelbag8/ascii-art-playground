package main

import "strings"

func TrimArtRows(rows []string) []string {
	lines := make([]string, len(rows))
	for i, row := range rows {
		lines[i] = strings.TrimRight(row, " ")
	}
	return lines
}
