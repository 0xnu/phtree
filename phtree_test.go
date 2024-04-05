package phtree

import (
	"math/rand"
	"sync"
	"testing"
)

func TestPHTree(t *testing.T) {
	tree := New(WithMaxDepth(8))

	// Test insertion
	tree.Insert([]uint64{0, 1, 2}, "value1")
	tree.Insert([]uint64{0, 1, 3}, "value2")
	tree.Insert([]uint64{1, 2, 3}, "value3")
	tree.Insert([]uint64{1, 2, 4}, "value4")

	// Test search
	testCases := []struct {
		key      []uint64
		expected interface{}
	}{
		{[]uint64{0, 1, 2}, "value1"},
		{[]uint64{0, 1, 3}, "value2"},
		{[]uint64{1, 2, 3}, "value3"},
		{[]uint64{1, 2, 4}, "value4"},
		{[]uint64{1, 2, 5}, nil},
	}

	for _, tc := range testCases {
		result := tree.Search(tc.key)
		if result != tc.expected {
			t.Errorf("Search(%v) = %v; want %v", tc.key, result, tc.expected)
		}
	}

	// Test remove
	tree.Remove([]uint64{0, 1, 3})
	if tree.Search([]uint64{0, 1, 3}) != nil {
		t.Error("Expected key to be removed")
	}

	// Test concurrent access
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			key := []uint64{uint64(rand.Intn(10)), uint64(rand.Intn(10)), uint64(rand.Intn(10))}
			value := rand.Intn(100)
			tree.Insert(key, value)
			result := tree.Search(key)
			if result != value {
				t.Errorf("Concurrent access failed. Expected %v, got %v", value, result)
			}
			tree.Remove(key)
		}()
	}
	wg.Wait()
}

func TestWithMaxDepth(t *testing.T) {
	maxDepth := 4
	tree := New(WithMaxDepth(maxDepth))

	// Test insertion with keys exceeding max depth
	tree.Insert([]uint64{0, 1, 2, 3, 4}, "value1")
	tree.Insert([]uint64{0, 1, 2, 3, 5}, "value2")

	// Test search
	if tree.Search([]uint64{0, 1, 2, 3, 4}) != "value1" {
		t.Error("Expected value1")
	}
	if tree.Search([]uint64{0, 1, 2, 3, 5}) != "value2" {
		t.Error("Expected value2")
	}
}

func TestRemove(t *testing.T) {
	tree := New()

	// Insert key-value pairs
	tree.Insert([]uint64{0, 1, 2}, "value1")
	tree.Insert([]uint64{0, 1, 3}, "value2")
	tree.Insert([]uint64{1, 2, 3}, "value3")

	// Remove a key-value pair
	tree.Remove([]uint64{0, 1, 3})

	// Test search after removal
	if tree.Search([]uint64{0, 1, 3}) != nil {
		t.Error("Expected key to be removed")
	}

	// Test search for remaining key-value pairs
	if tree.Search([]uint64{0, 1, 2}) != "value1" {
		t.Error("Expected value1")
	}
	if tree.Search([]uint64{1, 2, 3}) != "value3" {
		t.Error("Expected value3")
	}
}

func TestEqualKeys(t *testing.T) {
	testCases := []struct {
		a, b     []uint64
		expected bool
	}{
		{[]uint64{1, 2, 3}, []uint64{1, 2, 3}, true},
		{[]uint64{1, 2, 3}, []uint64{1, 2, 4}, false},
		{[]uint64{1, 2, 3}, []uint64{1, 2}, false},
		{[]uint64{1, 2}, []uint64{1, 2, 3}, false},
	}

	for _, tc := range testCases {
		result := equalKeys(tc.a, tc.b)
		if result != tc.expected {
			t.Errorf("equalKeys(%v, %v) = %v; want %v", tc.a, tc.b, result, tc.expected)
		}
	}
}

func TestBitCount(t *testing.T) {
	testCases := []struct {
		input    uint64
		expected int
	}{
		{0b0000000000000000, 0},
		{0b0000000000000001, 1},
		{0b0000000000000011, 2},
		{0b0000000000000111, 3},
		{0b1111111111111111, 16},
	}

	for _, tc := range testCases {
		result := bitCount(tc.input)
		if result != tc.expected {
			t.Errorf("bitCount(%b) = %d; want %d", tc.input, result, tc.expected)
		}
	}
}
