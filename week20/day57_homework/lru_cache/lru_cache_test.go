package lru_cache

import "testing"

func TestLRUCacheCase1(t *testing.T) {
	lru := Constructor(2)
	lru.Put(2, 1)
	lru.Put(1, 1)
	lru.Put(2, 3)
	lru.Put(4, 1)
	got1 := lru.Get(1)
	if got1 != -1 {
		t.Fatalf("tested 1.1 expected %+v but got %+v", -1, got1)
	}
	got2 := lru.Get(2)
	if got2 != 3 {
		t.Fatalf("tested 1.2 expected %+v but got %+v", 3, got1)
	}
}

func TestLRUCacheCase2(t *testing.T) {
	type TestCase struct {
		capacity   int
		operations []struct {
			isPut      bool
			key, value int
		}
	}
	testCase := TestCase{
		capacity: 4,
		operations: []struct {
			isPut bool
			key   int
			value int
		}{
			{isPut: true, key: 5, value: 13}, //1
			{isPut: true, key: 9, value: 6},
			{isPut: true, key: 4, value: 1},
			{isPut: false, key: 4, value: 1},
			{isPut: true, key: 6, value: 1}, //5
			{isPut: true, key: 8, value: 11},
			{isPut: false, key: 13, value: -1},
			{isPut: false, key: 1, value: -1},
			{isPut: true, key: 12, value: 12},
			{isPut: false, key: 10, value: -1}, //10
			{isPut: true, key: 15, value: 14},
			{isPut: true, key: 2, value: 12},
			{isPut: true, key: 7, value: 5},
			{isPut: true, key: 10, value: 3},
			{isPut: false, key: 6, value: -1}, //15
			{isPut: false, key: 10, value: 3},
			{isPut: true, key: 15, value: 14},
			{isPut: true, key: 5, value: 12},
			{isPut: false, key: 5, value: 12},
			{isPut: false, key: 7, value: 5}, //20
			{isPut: false, key: 15, value: 14},
			{isPut: false, key: 5, value: 12},
			{isPut: false, key: 6, value: -1},
			{isPut: false, key: 10, value: 3},
			{isPut: true, key: 7, value: 13}, //25
			{isPut: false, key: 14, value: -1},
			{isPut: true, key: 8, value: 9},
			{isPut: false, key: 4, value: -1},
			{isPut: true, key: 6, value: 11},
			{isPut: false, key: 9, value: -1}, //30
			{isPut: true, key: 6, value: 12},
			{isPut: false, key: 3, value: -1},
		},
	}
	// S 15 14 S 5 12 G 5 G 7 G 15 G 5 G 6 G 10 S 7 13 G 14 S 8 9 G 4 S 6 11 G 9 S 6 12 G 3
	// -1 3 12 5 14 12 -1 3 -1 -1 -1 -1
	lru := Constructor(testCase.capacity)
	for index, operation := range testCase.operations {
		if operation.isPut {
			lru.Put(operation.key, operation.value)
		} else {
			got := lru.Get(operation.key)
			if got != operation.value {
				t.Fatalf("tested 2.%d expected %+v but got %+v", (index + 1), operation.value, got)
			}
		}
	}
}
