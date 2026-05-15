package main


func copyBanners(banner []string) []string{
	copied := make([]string, len(banner))
	copy(copied,banner)

	return copied
}

func MergeBanners(base map[rune][]string, priority map[rune][]string) map[rune][]string{
	merged := make(map[rune][]string)

	for key, values := range base{
		merged[key] = copyBanners(values)
	}

	for key, values := range priority{
		merged[key] = copyBanners(values)
	}

	return merged
}