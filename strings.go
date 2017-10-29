package slicer

// Strings slices a given slice into a smaller slices of a given size.
func Strings(slice []string, perSlice int) [][]string {
	if perSlice >= len(slice) {
		return [][]string{slice}
	}

	var result [][]string
	for i := 0; i < len(slice); i += perSlice {
		limit := i + perSlice
		if limit > len(slice) {
			limit = len(slice)
		}
		result = append(result, slice[i:limit])
	}
	return result
}
