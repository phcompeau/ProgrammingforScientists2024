package main

import (
	"math/rand"
	"sort"
)

// SimpsonsMap takes an array of frequency maps as input. It returns a
// map mapping each sample name to its Simpson's index.
func SimpsonsMap(allMaps map[string](map[string]int)) map[string]float64 {

	s := make(map[string]float64)

	for sampleName, freqMap := range allMaps {
		s[sampleName] = SimpsonsIndex(freqMap)
	}
	return s
}

// RichnessMap takes a map of frequency maps as input.  It returns a map
// whose values are the richness of each sample.
func RichnessMap(allMaps map[string](map[string]int)) map[string]int {

	r := make(map[string]int)

	for sampleName, freqMap := range allMaps {
		r[sampleName] = Richness(freqMap)
	}

	return r
}

// BetaDiversityMatrix takes a map of frequency maps along with a distance metric
// ("Bray-Curtis" or "Jaccard") as input.
// It returns a slice of strings corresponding to the sorted names of the keys
// in the map, along with a matrix of distances whose (i,j)-th
// element is the distance between the i-th and j-th samples using the input metric.
func BetaDiversityMatrix(allMaps map[string](map[string]int), distMetric string) ([]string, [][]float64) {

	//first grab all strings
	numSamples := len(allMaps)
	sampleNames := make([]string, 0)
	for name := range allMaps {
		sampleNames = append(sampleNames, name)
	}

	// now sort sample names to make matrix ordered
	sort.Strings(sampleNames)

	// now form the distance matrix

	mtx := InitializeSquareMatrix(numSamples)

	for i := 0; i < numSamples; i++ {
		for j := i; j < numSamples; j++ {
			if distMetric == "Bray-Curtis" {
				d := BrayCurtisDistance(allMaps[sampleNames[i]], allMaps[sampleNames[j]])
				mtx[i][j] = d
				mtx[j][i] = d
			} else if distMetric == "Jaccard" {
				d := JaccardDistance(allMaps[sampleNames[i]], allMaps[sampleNames[j]])
				mtx[i][j] = d
				mtx[j][i] = d
			} else {
				panic("Error: Invalid distance metric name given to BetaDiversityMatrix().")
			}
		}
	}
	return sampleNames, mtx
}

// DownSampleMaps takes a map of frequency maps and a threshold, and returns a map of frequency maps with the same keys, but with each frequency map randomly downsampled to the threshold.
func DownSampleMaps(allMaps map[string]map[string]int, threshold int) map[string]map[string]int {
	newMaps := make(map[string]map[string]int)

	for key, freqMap := range allMaps {
		newMap := DownSample(freqMap, threshold)
		newMaps[key] = newMap
	}

	return newMaps
}

// DownSample takes as input a frequency map and a threshold, and returns a new frequency map with the same keys, but with each value randomly downsampled to the threshold.
func DownSample(freqMap map[string]int, threshold int) map[string]int {
	total := SampleTotal(freqMap)
	if total < threshold {
		panic("DownSample() called on a frequency map with a total value less than the threshold!")
	}

	// Create a slice to store all the keys, repeated according to their frequency
	allKeys := make([]string, 0, total)
	for key, count := range freqMap {
		for i := 0; i < count; i++ {
			allKeys = append(allKeys, key)
		}
	}

	// Get a random permutation of indices
	perm := rand.Perm(total)

	// Create a new map to store the downsampled results
	newMap := make(map[string]int)

	// Take the first 'threshold' keys from the permuted list and count them
	for i := 0; i < threshold; i++ {
		key := allKeys[perm[i]]
		newMap[key]++
	}

	return newMap
}
