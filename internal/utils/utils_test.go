package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitSliceInt_LastChunk(t *testing.T) {
	assert.Equal(
		t,
		[][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
			{10}},
		SplitSliceInt(
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			3),
		"must be equal")

	assert.Equal(
		t,
		[][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9}},
		SplitSliceInt(
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			3),
		"must be equal")
}

func TestSplitSliceInt_NegativeChunkSize(t *testing.T) {
	assert.Equal(
		t,
		[][]int{{1, 2, 3}},
		SplitSliceInt(
			[]int{1, 2, 3},
			-1),
		"must return entire slice as first element when chunkSize is -1")
}
func TestSplitSliceInt_ZeroChunkSize(t *testing.T) {
	assert.Equal(
		t,
		[][]int{{1, 2, 3}},
		SplitSliceInt(
			[]int{1, 2, 3},
			0),
		"must return entire slice as first element when chunkSize is 0")
}

func TestSplitSliceInt_BigChunkSize(t *testing.T) {
	assert.Equal(
		t,
		[][]int{{1, 2, 3}},
		SplitSliceInt(
			[]int{1, 2, 3},
			10),
		"must return entire slice as first element when chunkSize is bigger than input slice length")
}

func TestSplitSliceString_LastChunk(t *testing.T) {
	assert.Equal(
		t,
		[][]string{
			{"1", "2", "3"},
			{"4", "5", "6"},
			{"7", "8", "9"},
			{"10"}},
		SplitSliceString(
			[]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"},
			3),
		"must be equal")

	assert.Equal(
		t,
		[][]string{
			{"1", "2", "3"},
			{"4", "5", "6"},
			{"7", "8", "9"}},
		SplitSliceString(
			[]string{"1", "2", "3", "4", "5", "6", "7", "8", "9"},
			3),
		"must be equal")
}

func TestSplitSliceString_NegativeChunkSize(t *testing.T) {
	assert.Equal(
		t,
		[][]string{{"1", "2", "3"}},
		SplitSliceString(
			[]string{"1", "2", "3"},
			-1),
		"must return entire slice as first element when chunkSize is -1")
}
func TestSplitSliceString_ZeroChunkSize(t *testing.T) {
	assert.Equal(
		t,
		[][]string{{"1", "2", "3"}},
		SplitSliceString(
			[]string{"1", "2", "3"},
			0),
		"must return entire slice as first element when chunkSize is 0")
}

func TestSplitSliceString_BigChunkSize(t *testing.T) {
	assert.Equal(
		t,
		[][]string{{"1", "2", "3"}},
		SplitSliceString(
			[]string{"1", "2", "3"},
			10),
		"must return entire slice as first element when chunkSize is bigger than input slice length")
}

func TestReverseMap(t *testing.T) {
	assert.Equal(
		t,
		map[int]string{
			1: "a",
			2: "b",
			3: "c",
		},
		ReverseMap(map[string]int{
			"a": 1,
			"b": 2,
			"c": 3,
		}),
		"must be equal",
	)
}

func TestReverseMap_EmptyInput(t *testing.T) {
	assert.Equal(
		t,
		map[int]string{},
		ReverseMap(map[string]int{}),
		"must be equal",
	)
}

func TestFilterSliceInt_NotContains(t *testing.T) {
	assert.NotContains(
		t,
		FilterSliceInt([]int{1, 2, 3, 4, 5, 6, 7, 8}),
		[]int{1, 2, 3},
		"1,2,3 must not be returned",
	)
}

func TestFilterSliceInt_EmptyInput(t *testing.T) {
	assert.Empty(
		t,
		FilterSliceInt([]int{}),
		"empty input must return empty result",
	)
}

func TestFilterSliceInt_DuplicatedValues(t *testing.T) {
	assert.NotContains(
		t,
		FilterSliceInt([]int{1, 1, 2, 3, 3, 3, 3, 4, 5, 6, 7, 8, 1, 1, 1}),
		[]int{1, 2, 3},
		"1,2,3 must not be returned",
	)
}

func TestFilterSliceString_NotContains(t *testing.T) {
	assert.NotContains(
		t,
		FilterSliceString([]string{"a", "b", "c", "d", "e", "f", "g", "h"}),
		[]string{"a", "b", "c"},
		"a,b,c must not be returned",
	)
}

func TestFilterSliceString_EmptyInput(t *testing.T) {
	assert.Empty(
		t,
		FilterSliceString([]string{}),
		"empty input must return empty result",
	)
}

func TestFilterSliceString_DuplicatedValues(t *testing.T) {
	assert.NotContains(
		t,
		FilterSliceString([]string{"a", "a", "b", "a", "c", "a", "d", "a", "e", "a", "f", "a", "g", "h"}),
		[]string{"a", "b", "c"},
		"a,b,c must not be returned",
	)
}

func TestFilterSliceString_DuplicatedValuesInResult(t *testing.T) {
	assert.Equal(
		t,
		[]string{"d", "d", "e", "f", "f", "g", "h"},
		FilterSliceString([]string{"a", "a", "b", "a", "c", "d", "d", "a", "e", "a", "f", "a", "f", "g", "h"}),
		"must be equal",
	)
}
