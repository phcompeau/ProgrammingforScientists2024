package main

// BrayCurtisDistance takes two frequency maps and returns the Bray-Curtis
// distance between them.
func BrayCurtisDistance(map1 map[string]int, map2 map[string]int) float64 {
	c := SumOfMinima(map1, map2)
	s1 := SampleTotal(map1)
	s2 := SampleTotal(map2)

	//don't allow the situation in which we have zero richness.
	if s1 == 0 || s2 == 0 {
		panic("Error: sample given to BrayCurtisDistance() has no positive values.")
	}

	av := Average(float64(s1), float64(s2))
	return 1 - (float64(c) / av)
}

// JaccardDistance takes two frequency maps and returns the Jaccard
// distance between them.
func JaccardDistance(map1 map[string]int, map2 map[string]int) float64 {
	c := SumOfMinima(map1, map2)
	t := SumOfMaxima(map1, map2)
	j := 1 - (float64(c) / float64(t))
	return j
}

// SumOfMaxima takes two frequency maps as input.
// It identifies the keys that are shared by the two frequency maps
// and returns the sum of the corresponding values. (a.k.a. "union")
// SumOfMaxima takes two frequency maps as input.
// It identifies the keys that are shared by the two frequency maps
// and returns the sum of the corresponding values. (a.k.a. "union")
func SumOfMaxima(map1 map[string]int, map2 map[string]int) int {
	c := 0

	for key := range map2 {
		_, exists := map1[key]
		if exists {
			c += Max2(map1[key], map2[key])
		} else {
			c += map2[key]
		}
	}
	for key := range map1 {
		_, exists := map2[key]
		if !exists {
			c += map1[key]
		}
	}

	// panic if c is equal to zero since we don't want 0/0
	if c == 0 {
		panic("Error: no species common to two maps given to SumOfMaxima")
	}

	return c
}

// Max2 takes two integers and returns their maximum.
func Max2(n1, n2 int) int {
	if n1 < n2 {
		return n2
	}
	return n1
}

// Richness takes a frequency map. It returns the richness of the frequency map
// (i.e., the number of keys in the map corresponding to nonzero values.)
func Richness(sample map[string]int) int {
	c := 0

	for _, val := range sample {
		if val < 0 {
			panic("Error: negative value in frequency map given to Richness()")
		}
		c++
	}

	return c
}

// SimpsonsIndex takes a frequency map and returns a decimal corresponding to Simpson's index:
// the sum of (n/N)^2 where n is the number of individuals found for a given string/species
// and N is the total number of individuals. The sum is over all species present.
func SimpsonsIndex(sample map[string]int) float64 {
	total := SampleTotal(sample)
	simpson := 0.0

	if total == 0 {
		panic("Error: Empty frequency map given to SimpsonsIndex()!")
	}

	for _, val := range sample {
		x := float64(val) / float64(total)
		simpson += x * x
	}
	return simpson
}

// InitializeSquareMatrix takes an integer n and returns an nxn slice of
// floats initialized to 0.0.
func InitializeSquareMatrix(n int) [][]float64 {
	mtx := make([][]float64, n)

	for i := range mtx {
		mtx[i] = make([]float64, n)
	}
	return mtx
}

// FrequencyMap forms the frequency map of a collection of input patterns.
// Input: one collection of strings patterns
// Output: a frequency map of strings to their # of counts in patterns
func FrequencyMap(patterns []string) map[string]int {
	freqMap := make(map[string]int)
	for _, val := range patterns {
		freqMap[val]++
	}
	return freqMap
}

// Average takes two floats and returns their average.
func Average(x, y float64) float64 {
	return (x + y) / 2.0
}

// SampleTotal takes a frequency map as input.
// It returns the sum of the counts for each string in a sample.
func SampleTotal(freqMap map[string]int) int {
	total := 0
	for _, val := range freqMap {
		if val > 0 {
			total += val
		}
	}
	return total
}

// SumOfMinima takes two frequency maps as input.
// It identifies the keys that are shared by the two frequency maps
// and returns the sum of the corresponding values.
func SumOfMinima(map1 map[string]int, map2 map[string]int) int {
	c := 0

	for key := range map1 {
		_, exists := map2[key]
		if exists {
			c += Min2(map1[key], map2[key])
		}
	}

	return c
}

// Min2 takes two integers and returns their minimum.
func Min2(n1, n2 int) int {
	if n1 < n2 {
		return n1
	}
	return n2
}
