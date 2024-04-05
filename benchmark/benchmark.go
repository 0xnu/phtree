package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/0xnu/phtree"
)

const (
	numDimensions = 4
	numPoints     = 1000000
	numQueries    = 10000
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// Generate random data points
	data := make([][]uint64, numPoints)
	for i := 0; i < numPoints; i++ {
		point := make([]uint64, numDimensions)
		for j := 0; j < numDimensions; j++ {
			point[j] = rand.Uint64()
		}
		data[i] = point
	}

	// Create a new PH-Tree
	tree := phtree.New(phtree.WithMaxDepth(16))

	// Benchmark insertion
	start := time.Now()
	for _, point := range data {
		tree.Insert(point, nil)
	}
	elapsed := time.Since(start)
	fmt.Printf("Insertion of %d points took %s\n", numPoints, elapsed)

	// Generate random query points
	queries := make([][]uint64, numQueries)
	for i := 0; i < numQueries; i++ {
		query := make([]uint64, numDimensions)
		for j := 0; j < numDimensions; j++ {
			query[j] = rand.Uint64()
		}
		queries[i] = query
	}

	// Benchmark search
	start = time.Now()
	for _, query := range queries {
		tree.Search(query)
	}
	elapsed = time.Since(start)
	fmt.Printf("Search of %d queries took %s\n", numQueries, elapsed)

	// Benchmark removal
	start = time.Now()
	for _, point := range data {
		tree.Remove(point)
	}
	elapsed = time.Since(start)
	fmt.Printf("Removal of %d points took %s\n", numPoints, elapsed)
}
