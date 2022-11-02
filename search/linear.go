package search

// Linear Simple linear search algorithm that iterates over all elements of an array in the worst case scenario
func Linear(array []int, query int, Size int) (bool, error) {
	for i := 0; i < Size; i++ {
		if array[i] == query {
			return true, nil // true we found that query
		}
	}
	return false, ErrNotFound //we didn't find it
}
