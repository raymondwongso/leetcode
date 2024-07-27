package main

/*
Heap Sort
Comparison-based sorting technique based on Binary Heap data structure.
There are max heap sort and min heap sort.
Max heap sort stores the largest element at root

	for each node i, the children of that node has value smaller than node i

Min heap sort stores the smallest element at root

	for each node i, the children of that node has value larger than node i

The algorithm:
1. Heapify means correcting the tree to conform with heap property.
2. Sort ascending use max heap sort, sort decending use min heap sort.
3. Heapify from last non-leaf node to root. last non-leaf node is n/2 - 1
  - For each swap happened, heapify again to the affected sub-tree

4. Starting from right most leaf, swap with root and heapify the reduced heap (heapify root while excluding the just swapped leaf)
  - repeat until root

Time complexity: O(n log n)
Space complexity: O(1)
*/
func heapSort(arr []int) []int {
	n := len(arr)

	// heapify from non-leaf node to root.
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	// after tree is heapified correctly, start extracting element
	// swap the root with the leaf, so that leaf now contains the largest value
	// heapify the root reduced tree,
	for i := n - 1; i >= 0; i-- {
		arr[i], arr[0] = arr[0], arr[i]
		// we use i here because i contains the index of last leaf processed.
		heapify(arr, i, 0)
	}

	return arr
}

func heapify(arr []int, n, currentIndex int) {
	largest := currentIndex
	left := largest*2 + 1
	right := largest*2 + 2

	// check if left child exists and whether the value is larger than current node
	if left < n && arr[left] > arr[largest] {
		largest = left
	}

	// check if right child exists and whether the value is larger than current node
	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	// if there is larger child node, swap with current index and heapify again the affected sub-tree
	// largest contains the original index of the largest child, heapify it
	if largest != currentIndex {
		arr[largest], arr[currentIndex] = arr[currentIndex], arr[largest]
		heapify(arr, n, largest)
	}
}
