package main

import (
	"math/rand"
	"sort"
)

// BetaDiversityMatrix takes a map of frequency maps along with a distance metric
// ("Bray-Curtis" or "Jaccard") as input.
// It returns a slice of strings corresponding to the sorted names of the keys
// in the map, along with a matrix of distances whose (i,j)-th
// element is the distance between the i-th and j-th samples using the input metric.
func BetaDiversityMatrix(allMaps map[string](map[string]int), distanceMetric string) ([]string, [][]float64) {
	numSamples := len(allMaps)
	sampleNames := make([]string, 0, numSamples)

	for sampleName := range allMaps {
		//range over all the sample IDs
		// add current sample ID to sample names slice
		sampleNames = append(sampleNames, sampleName)
	}

	sort.Strings(sampleNames)

	//sampleNames is now sorted

	// matrix should be numSamples x numSamples
	mtx := InitializeSquareMatrix(numSamples)

	// range through the matrix and set its values
	for r := range mtx {
		for c := r + 1; c < len(mtx[r]); c++ {
			// set mtx[r][c]
			sampleName1 := sampleNames[r]
			sampleName2 := sampleNames[c]

			sample1 := allMaps[sampleName1]
			sample2 := allMaps[sampleName2]

			if distanceMetric == "Jaccard" {
				mtx[r][c] = JaccardDistance(sample1, sample2)
			} else if distanceMetric == "Bray-Curtis" {
				mtx[r][c] = BrayCurtisDistance(sample1, sample2)
			} else {
				panic("bad")
			}

			//set the value that is symmetric across main diagonal
			mtx[c][r] = mtx[r][c]
		}
	}

	return sampleNames, mtx
}

// SimpsonsMap takes a map of frequency maps as input. It returns a
// map mapping each sample name to its Simpson's index.
func SimpsonsMap(allMaps map[string](map[string]int)) map[string]float64 {
	s := make(map[string]float64)

	//range over all sample names
	for sampleName := range allMaps {
		//set current sample name's evenness
		//equal to simpson's index of current frequency table
		s[sampleName] = SimpsonsIndex(allMaps[sampleName])
	}

	return s
}

// RichnessMap takes a map of frequency maps as input.  It returns a map
// whose values are the richness of each sample.
func RichnessMap(allMaps map[string](map[string]int)) map[string]int {
	r := make(map[string]int)

	for sampleName := range allMaps {
		// range over all sample names
		// for each sample name, set its richness value by calling Richness()
		currentFrequencyTable := allMaps[sampleName]
		r[sampleName] = Richness(currentFrequencyTable)
	}

	return r
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
