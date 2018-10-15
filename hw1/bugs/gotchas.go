package main

import (
	"sort"
	"strconv"
)

type Keys []int

func (this Keys) Len() int {
	return len(this)
}

func (this Keys) Less(i, j int) bool {
	return this[i] < this[j]
}

func (this Keys) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func ReturnInt() int {
	return 1
}

func ReturnFloat() float32 {
	return 1.1
}

func ReturnIntArray() [3]int {
	return [3]int{1, 3, 4}
}

func ReturnIntSlice() []int {
	return []int{1, 2, 3}
}

func IntSliceToString(slice []int) string {
	var res string
	for _, i := range slice {
		res += strconv.Itoa(i)
	}
	return res
}

func MergeSlices(slice1 []float32, slice2 []int32) []int {
	var res []int
	for _, i := range slice1 {
		res = append(res, int(i))
	}
	for _, i := range slice2 {
		res = append(res, int(i))
	}
	return res
}

func GetMapValuesSortedByKey(input map[int]string) []string {
	var res []string
	var keys []int
	for i, _ := range input {
		keys = append(keys, i)
	}
	sort.Sort(Keys(keys))
	for _, i := range keys {
		res = append(res, input[i])
	}
	return res
}
