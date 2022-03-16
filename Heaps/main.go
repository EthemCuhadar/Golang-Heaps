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