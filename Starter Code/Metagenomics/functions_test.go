//write testing code for each function in functions.go
//using the tests in the Tests file

package main

import (
	"bufio"
	"io/fs"
	"math"
	"os"
	"strconv"
	"strings"
	"testing"
)

// RichnessTest is a struct that holds the information for a test that takes a sample and an integer
type RichnessTest struct {
	sample map[string]int
	result int
}

// SimpsonsIndexTest is a struct that holds the information for a test that takes a sample and a float
type SimpsonsIndexTest struct {
	sample map[string]int
	result float64
}

// TestRichness tests the Richness function
func TestRichness(t *testing.T) {
	//read in all tests from the Tests/Richness directory and run them
	tests := ReadRichnessTests("Tests/Richness/")
	for _, test := range tests {
		//run the test
		result := Richness(test.sample)
		//check the result
		if result != test.result {
			t.Errorf("Richness(%v) = %v, want %v", test.sample, result, test.result)
		}
	}
}

// ReadRichnessTests takes as input a directory and returns a slice of RichnessTest objects
func ReadRichnessTests(directory string) []RichnessTest {

	//read in all tests from the directory and run them
	inputFiles := ReadDirectory(directory + "/input")
	numFiles := len(inputFiles)

	tests := make([]RichnessTest, numFiles)
	for i, inputFile := range inputFiles {
		//read in the test's map
		tests[i].sample = ReadSample(directory + "input/" + inputFile.Name())
	}

	//now, read output files
	outputFiles := ReadDirectory(directory + "/output")
	if len(outputFiles) != numFiles {
		panic("Error: number of input and output files do not match!")
	}

	for i, outputFile := range outputFiles {
		//read in the test's result
		tests[i].result = ReadIntegerFromFile(directory + "output/" + outputFile.Name())
	}

	return tests
}

// ReadDirectory reads in a directory and returns a slice of fs.DirEntry objects containing file info for the directory
func ReadDirectory(dir string) []fs.DirEntry {
	//read in all files in the given directory
	files, err := os.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	return files
}

// ReadSample reads in a frequency map from a file
func ReadSample(file string) map[string]int {
	//open the file
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	//create a new scanner
	scanner := bufio.NewScanner(f)

	sample := make(map[string]int)
	//while the scanner still has lines to read,
	//read in the next line
	for scanner.Scan() {
		//read in the line
		line := scanner.Text()
		//split the line into two parts
		parts := strings.Split(line, " ")
		//add the key and value to the map
		pattern := parts[0]
		//convert parts[1] to an int using strconv
		value, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		sample[pattern] = value
	}

	return sample
}

// ReadIntegerFromFile reads in a single integer from a file
func ReadIntegerFromFile(file string) int {
	//open the file
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	//create a new scanner
	scanner := bufio.NewScanner(f)

	//read in the line
	scanner.Scan()
	line := scanner.Text()

	//convert the line to an int using strconv
	value, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	return value
}

// TestSimpsonsIndex tests the SimpsonsIndex function
func TestSimpsonsIndex(t *testing.T) {
	//read in all tests from the Tests/SimpsonsIndex directory and run them
	tests := ReadSimpsonsIndexTests("Tests/SimpsonsIndex/")
	for _, test := range tests {
		//run the test
		result := SimpsonsIndex(test.sample)
		//check the result
		if roundFloat(result, 4) != roundFloat(test.result, 4) {
			t.Errorf("SimpsonsIndex(%v) = %v, want %v", test.sample, result, test.result)
		}
	}
}

// ReadSimpsonsIndexTests takes as input a directory and returns a slice of SimpsonsIndexTest objects
func ReadSimpsonsIndexTests(directory string) []SimpsonsIndexTest {

	//read in all tests from the directory and run them
	inputFiles := ReadDirectory(directory + "/input")
	numFiles := len(inputFiles)

	tests := make([]SimpsonsIndexTest, numFiles)
	for i, inputFile := range inputFiles {
		//read in the test's map
		tests[i].sample = ReadSample(directory + "input/" + inputFile.Name())
	}

	//now, read output files
	outputFiles := ReadDirectory(directory + "/output")
	if len(outputFiles) != numFiles {
		panic("Error: number of input and output files do not match!")
	}

	for i, outputFile := range outputFiles {
		//read in the test's result
		tests[i].result = ReadFloatFromFile(directory + "output/" + outputFile.Name())
	}

	return tests
}

// ReadFloatFromFile reads in a single float from a file
func ReadFloatFromFile(file string) float64 {
	//open the file
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	//create a new scanner
	scanner := bufio.NewScanner(f)

	//read in the line
	scanner.Scan()
	line := scanner.Text()

	//convert the line to an int using strconv
	value, err := strconv.ParseFloat(line, 64)
	if err != nil {
		panic(err)
	}

	return value
}

// roundFloat rounds a float to a given number of decimals precision
func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
