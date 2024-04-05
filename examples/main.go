package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/0xnu/phtree"
)

func main() {
	// Create a new PHTree with a maximum depth of 8
	tree := phtree.New(phtree.WithMaxDepth(8))

	// Insert key-value pairs into the PHTree
	tree.Insert([]uint64{0, 1, 2}, "apple")
	tree.Insert([]uint64{0, 1, 3}, "banana")
	tree.Insert([]uint64{1, 2, 3}, "orange")
	tree.Insert([]uint64{1, 2, 4}, "grape")

	// Search for values in the PHTree
	searchKeys := [][]uint64{
		{0, 1, 2},
		{0, 1, 3},
		{1, 2, 3},
		{1, 2, 4},
		{1, 2, 5},
	}

	for _, key := range searchKeys {
		value := tree.Search(key)
		if value != nil {
			fmt.Printf("Key %v found with value: %v\n", key, value)
		} else {
			fmt.Printf("Key %v not found\n", key)
		}
	}

	// Remove a key-value pair from the PHTree
	tree.Remove([]uint64{0, 1, 3})
	fmt.Println("Removed key [0 1 3]")

	// Search for the removed key
	value := tree.Search([]uint64{0, 1, 3})
	if value != nil {
		fmt.Printf("Key [0 1 3] found with value: %v\n", value)
	} else {
		fmt.Println("Key [0 1 3] not found")
	}

	// Concurrent access to the PHTree
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			key := []uint64{uint64(rand.Intn(5)), uint64(rand.Intn(5)), uint64(rand.Intn(5))}
			value := fmt.Sprintf("value-%d", rand.Intn(100))
			tree.Insert(key, value)
			fmt.Printf("Inserted key %v with value: %v\n", key, value)
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			searchValue := tree.Search(key)
			fmt.Printf("Searched key %v and found value: %v\n", key, searchValue)
			tree.Remove(key)
			fmt.Printf("Removed key %v\n", key)
		}()
	}
	wg.Wait()
}
