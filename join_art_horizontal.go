package main

func JoinArtHorizontal(left []string, right []string) []string {
	newSlice := make([]string, len(left))

	for i := range left {
		newSlice[i] = left[i] + right[i]
	}

	return newSlice

}
