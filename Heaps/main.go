package main

import "fmt"

// Heap - Struct which has decsending order from up to bottom.
type Heap struct{
	array 	[]int
}

// Insert - Adds an element in the heap structure.
func (h *Heap) Insert(key int){
	h.array = append(h.array, key)
	h.heapifyUp(len(h.array)-1)
}

// Extract - Eliminates the element which has the largest key in the heap.
func (h *Heap) Extract() int{
	extracted := h.array[0]
	l := len(h.array)-1
	// Error case when there is no item in heap.
	if len(h.array) == 0{
		fmt.Println("Error: No item to extract!")
		return -1
	}
	// Take the last index of heap and assign it as root of heap.
	h.array[0] = h.array[l]
	h.array = h.array[:l]
	
	h.heapifyDown(0)
	
	return extracted
}


// maxHeapifyUp - Compare and change indices of child
// and parent elements from bottom to top. 
func (h *Heap) heapifyUp(index int){
	for h.array[parent(index)] < h.array[index]{
		h.swap(parent(index), index)
		index = parent(index)
	}
}

// heapifyDown - Compare and change indices of parent
// and child elements from top to bottom.
func (h *Heap) heapifyDown(index int){

	lastIndex := len(h.array)-1
	l, r := left(index), right(index)
	childToCompare := 0
	
	// loop while index has at least one child.
	for l <= lastIndex{
		if l == lastIndex{		// when left child is the only child.
			childToCompare = l
		}else if h.array[l] > h.array[r]{	// when left child is larger.
			childToCompare = l
		}else{					// when right child is larger.
			childToCompare = r
		}
		// compare array value of current index to larger child and swap if smaller
		if h.array[index] < h.array[childToCompare]{
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		}else{
			return
		}
	}
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

func main(){
	H := &Heap{}
	fmt.Println(H)
	
	buildHeap := []int{100, 200, 300, 150, 222, 15, 369, 555, 136, 458, 364, 562}
	for _, v := range buildHeap{
		H.Insert(v)
		fmt.Println(H)
	}
	for i:=0; i<5; i++{
		H.Extract()
		fmt.Println(H)
	}
}