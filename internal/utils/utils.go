package utils

func SplitSliceInt(input []int, chunkSize int) [][]int {
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

func SplitSliceString(input []string, chunkSize int) [][]string {
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

func ReverseMap(input map[string]int) map[int]string {
	result := make(map[int]string, len(input))
	for k, v := range input {
		result[v] = k
	}
	return result
}

func FilterSliceInt(input []int) []int {
	hardcodedSlice := []int{1, 2, 3}
	result := make([]int, 0, len(input))
	for _, v := range input {
		if !containsInt(hardcodedSlice, v) {
			result = append(result, v)
		}
	}
	return result
}

func FilterSliceString(input []string) []string {
	hardcodedSlice := []string{"a", "b", "c"}
	result := make([]string, 0, len(input))
	for _, v := range input {
		if !containsString(hardcodedSlice, v) {
			result = append(result, v)
		}
	}
	return result
}

func containsInt(input []int, val int) bool {
	for _, v := range input {
		if val == v {
			return true
		}
	}
	return false
}

func containsString(input []string, val string) bool {
	for _, v := range input {
		if val == v {
			return true
		}
	}
	return false
}
