package heap

type Heap struct {
	//heaps are simple but powerful array strucutured in away we can use to implement priority queues or use it for sorting and in graph khawrizims 
	arr []int 
}

func Init() *Heap {
	return &Heap{arr: []int{}}
}
//add/insert new item in the heap 
func (h *Heap) Push(val int) {
	h.arr = append(h.arr, val) 
	//once we added we do heapifyUp
	h.heapifyUp(len(h.arr) - 1)
}

func (h *Heap) heapifyUp(i int) {
	//do heapify up until we see current node is not less than its parent 
	for h.arr[i] < h.arr[parentIndex(i)] {
		swap(h.arr, i, parentIndex(i))
	}
}


func (h *Heap) Len() int {
	return len(h.arr)
}

//aka remove first element in the queue 
func (h *Heap) Pop() int {
	top := h.arr[0]
	//replace last one with the first one 
	lastIndex := len(h.arr) - 1
	h.arr[0] = h.arr[lastIndex]
	//exclude the last one (can be enhanced by maintaining last index position)
	h.arr = h.arr[:lastIndex] 

	h.heapifyDown(0)

	return top
}

//always start from zero index 
func (h *Heap) heapifyDown(i int) {

	left := leftIndex(i)
	right := rightIndex(i)
	if i > len(h.arr) - 1 || left > len(h.arr) -1 || right > len(h.arr) -1 {
		return
	}

	if h.arr[i] < h.arr[left] && h.arr[i] < h.arr[right] {
		return 
	}

	//current elem is greater than children 
	//find the smallest and swap 
	if h.arr[left] < h.arr[right] {
		swap(h.arr, i, left)
		h.heapifyDown(left)
	} else {
		swap(h.arr, i, right)
		h.heapifyDown(right)
	}

}

func (h *Heap) Top() int {
	return h.arr[0]
}

////////////////////////////////////////////////////////////////////////////////////

//internal helpers 
func parentIndex(i int) int {
	return (i - 1) / 2
}

func leftIndex(i int) int {
	return i * 2 + 1
}

func rightIndex(i int) int {
	return i * 2 + 2 
}


func swap(arr []int, i,j int) {
	t := arr[i]
	arr[i]  = arr[j] 
	arr[j] = t 
}