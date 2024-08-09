package main

import (
	"strconv"
	"strings"
)

//GetKeyValues takes a map of strings to strings as input.
//It returns two slices corresponding to the keys and values of the
//map, respectively.
func GetKeyValues(dnaMap map[string]string) ([]string, []string) {
	var keys = make([]string, 0, len(dnaMap))
	var values = make([]string, 0, len(dnaMap))
	for k, v := range dnaMap {
		keys = append(keys, k)
		values = append(values, v)
	}
	return keys, values
}

//GetKeyValuesI takes a map of strings to ints as input.
//It returns two slices corresponding to the keys and values of the
//map, respectively.
func GetKeyValuesI(dnaMap map[string]int) ([]string, []int) {
	var keys = make([]string, 0, len(dnaMap))
	var values = make([]int, 0, len(dnaMap))
	for k, v := range dnaMap {
		keys = append(keys, k)
		values = append(values, v)
	}
	return keys, values
}

//InitializeMatrix takes integers m and n and
//returns a DistanceMatrix of dimensions m x n with default values.
func InitializeMatrix(m int, n int) DistanceMatrix {
	mtx := make([][]float64, m)

	for i := range mtx {
		mtx[i] = make([]float64, n)
	}
	return mtx
}

//CreateFrequencyDNAMap takes a slice of strings. It produces a dictionary where dict[i]
//corresponds to slice[i] in the slice. Essentially provides dummy
//labels for unannotated species.
func CreateFrequencyDNAMap(dnaStrings []string) map[string]string {
	var freqMap = CreateFrequencyMap(dnaStrings)
	var keys, _ = GetKeyValuesI(freqMap)

	var dnaMap = make(map[string]string)
	for i := 0; i < len(keys); i++ {
		dnaMap[strconv.Itoa(i)] = keys[i]
	}
	return dnaMap
}

//CreateFrequencyMap takes a collection of strings and returns
//the frequency table of these strings, mapping a string
//to its number of occurrences.
func CreateFrequencyMap(patterns []string) map[string]int {
	freqMap := make(map[string]int)
	for _, val := range patterns {
		freqMap[val]++
	}
	return freqMap
}

/************************************************
  MISCELLANEOUS
************************************************/

//MinFloat returns the minimum of an arbitrary collection of floats.
func MinFloat(nums ...float64) float64 {
	m := 0.0
	// nums gets converted to an array
	for i, val := range nums {
		if val < m || i == 0 {
			// update m
			m = val
		}
	}
	return m
}

func getCats(labels []string) []string {
	cats := make([]string, 0)
	for _, label := range labels {
		cats = append(cats, strings.Split(label, "|")[0])
	}
	return unique(cats)
}

func unique(strSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range strSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func rearrangeStrings(newLabels []string, oldLabels []string, dnaStringsOld []string) []string {
	dnaStringsNew := make([]string, 0)
	for _, newLabel := range newLabels {
		j := getIndex(oldLabels, newLabel)
		dnaStringsNew = append(dnaStringsNew, dnaStringsOld[j])
	}
	return dnaStringsNew
}

func getIndex(arr []string, target string) int {
	for i, str := range arr {
		if str == target {
			return i
		}
	}
	return 0
}
