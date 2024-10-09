package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func WriteNewickToFile(t Tree, fileDest string, fileName string) {

	newickString := ToNewick(t)

	F, err := os.Create(fileDest + "/" + fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = F.WriteString(newickString)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = F.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

}

func ToNewick(tree Tree) string {
	return "(" + subtreeNewickAges(tree[len(tree)-1]) + ");"
}

func subtreeNewickAges(node *Node) string {
	if node.Child1 == nil {
		return node.Label + ":" + fmt.Sprintf("%.2f", node.Age)
	} else {
		return "(" + subtreeNewickAges(node.Child1) + "," + subtreeNewickAges(node.Child2) + "):" + fmt.Sprintf("%.2f", node.Age)
	}
}

// ReadMatrixFromFile takes a file name and reads the information in this file to produce
// a distance matrix and a slice of strings holding the species names.  The first line of the
// file should contain the number of species.  Each other line contains a species name
// and its distance to each other species.
func ReadMatrixFromFile(filename string) ([]string, DistanceMatrix) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	// Initialize species names slice and distance matrix
	var species []string
	distanceMatrix := make(DistanceMatrix, len(records))

	// Populate the species names and distance matrix
	for i, record := range records {
		// First element is the species name
		species = append(species, record[0])

		// Convert the remaining elements to float64
		distanceMatrix[i] = make([]float64, len(record)-1)
		for j, value := range record[1:] {
			distanceMatrix[i][j], err = strconv.ParseFloat(value, 64)
			if err != nil {
				panic(err)
			}
		}
	}

	return species, distanceMatrix
}
