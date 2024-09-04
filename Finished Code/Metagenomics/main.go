package main

import (
	"fmt"
)

func main() {
	fmt.Println("Metagenomics!")

	// step 1: reading input from a single file.

	filename := "Data/Fall_Allegheny_1.txt"
	freqMap := ReadFrequencyMapFromFile(filename)
	fmt.Println("File read successfully! We have", len(freqMap), "total patterns.")

	// we may as well do something with our file. For example, let's print its Simpson's Index.
	fmt.Println("Simpson's Index:", SimpsonsIndex(freqMap))

	//step 2: reading input from a directory

	path := "Data/"
	allMaps := ReadSamplesFromDirectory(path)

	fmt.Println("Samples read successfully! We have", len(allMaps), "total samples.")

	fmt.Println("Let's determine the depth of each sample.")

	depthMap := make(map[string]int)
	for key, freqMap := range allMaps {
		depthMap[key] = SampleTotal(freqMap)
	}

	fmt.Println("Printing the depth of each sample to the console.")

	for key, val := range depthMap {
		fmt.Println(key, "depth:", val)
	}

	// let's take a sequencing depth of 400
	sequencingDepth := 400

	fmt.Println("Downsampling all samples to a threshold of", sequencingDepth)

	downsampledMaps := DownSampleMaps(allMaps, sequencingDepth)

	// step 3: processing the data that we have received.

	// now all of our maps have been processed and we can start working with them.

	// for example, let's compute the evenness of each sample.

	simpson := SimpsonsMap(downsampledMaps)

	// now let's look at beta diversity.

	distMetric := "Bray-Curtis"
	sampleNames, mtx := BetaDiversityMatrix(downsampledMaps, distMetric)

	// we cannot really analyze anything from this printing.

	// It would be better to print to a file.  Hence, we will need to learn writing to a file.

	simpsonFile := "Matrices/SimpsonMatrix.csv"
	WriteSimpsonsMapToFile(simpson, simpsonFile)

	outFilename := "Matrices/BetaDiversityMatrix.csv"
	WriteBetaDiversityMatrixToFile(mtx, sampleNames, outFilename)

	fmt.Println("Success! Now we are ready to do something cool with our data.")
}
