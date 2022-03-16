package main

import "fmt"

// Heap - Struct which has decsending order from up to bottom.
type Heap struct{
	array 	[]int
}

// Insert - Adds an element in the heap structure.
func (h *Heap) Insert(key int){
	h.array = append(h.array, key)
}

// Extract - Eliminates the element which has the largest key in the heap.


func (h *Heap) maxHeapifyUp(index int){
	
}

// parent - Returns parent index of an element.
func parent(i int) int{
	return (i-1)/2
}

// left - Returns left child index of an element.
func left(i int) int{
	return 2*i + 1
}

// right - Return right child index of an element.
func right(i int) int{
	return 2*i + 2
}

// swap - Simply swaps two elements in array.
func (h *Heap) swap(i1, i2 int){
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}