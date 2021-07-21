package utils

// SplitSliceInt converts input slice of integers into slice of slices of integers.
// Size of chunk is defined by chunkSize. Size of last chunk is less or equal chunkSize.
func SplitSliceInt(input []int, chunkSize int) [][]int {
	if chunkSize <= 0 {
		return [][]int{input}
	}

	// calculating number of chunks to allocate memory
	// for resulting slice (len and cap)
	chunksCount := len(input) / chunkSize
	if len(input)%chunkSize > 0 {
		chunksCount = chunksCount + 1
	}
	result := make([][]int, chunksCount)

	for i := 0; i < chunksCount; i++ {
		lowBound := i * chunkSize
		highBound := lowBound + chunkSize
		if highBound > len(input) {
			highBound = len(input)
		}
		result[i] = input[lowBound:highBound]
	}
	return result
}

// SplitSliceString converts input slice of strings into slice of slices of strings.
// Size of chunk is defined by chunkSize. Size of last chunk is less or equal chunkSize.
func SplitSliceString(input []string, chunkSize int) [][]string {
	if chunkSize <= 0 {
		return [][]string{input}
	}
	// calculating number of chunks to allocate memory
	// for resulting slice (len and cap)
	chunksCount := len(input) / chunkSize
	if len(input)%chunkSize > 0 {
		chunksCount = chunksCount + 1
	}
	result := make([][]string, chunksCount)

	for i := 0; i < chunksCount; i++ {
		lowBound := i * chunkSize
		highBound := lowBound + chunkSize
		if highBound > len(input) {
			highBound = len(input)
		}
		result[i] = input[lowBound:highBound]
	}
	return result
}

// ReverseMap transforms key-value input map into value-key resulting map.
func ReverseMap(input map[string]int) map[int]string {
	result := make(map[int]string, len(input))
	for k, v := range input {
		result[v] = k
	}
	return result
}

// FilterSliceInt returns a slice with integer values of given input slice of integers,
// which are not in the harcoded list.
func FilterSliceInt(input []int) []int {
	hardcodedSlice := map[int]struct{}{1: {}, 2: {}, 3: {}}
	result := make([]int, 0, len(input))
	for _, v := range input {
		if _, ok := hardcodedSlice[v]; !ok {
			result = append(result, v)
		}
	}
	return result
}

// FilterSliceString returns a slice with string values of given input slice of strings,
// which are not in the harcoded list.
func FilterSliceString(input []string) []string {
	hardcodedSlice := map[string]struct{}{"a": {}, "b": {}, "c": {}}
	result := make([]string, 0, len(input))
	for _, v := range input {
		if _, ok := hardcodedSlice[v]; !ok {
			result = append(result, v)
		}
	}
	return result
}
