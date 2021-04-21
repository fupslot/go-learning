package main

import (
	"fmt"
	"sort"
	"sync/atomic"
)

type byLength []string

var lenOps uint64
var swapOps uint64
var lessOps uint64

func (s byLength) Len() int {
	atomic.AddUint64(&lenOps, 1)
	return len(s)
}

func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
	atomic.AddUint64(&swapOps, 1)
}

func (s byLength) Less(i, j int) bool {
	atomic.AddUint64(&lessOps, 1)
	return len(s[i]) > len(s[j])
}

func main() {
	strs := []string{"b", "d", "a"}
	sort.Strings(strs)

	fmt.Println("Strings:", strs)
	ints := []int{4, 12, 7, 3, 2}
	sort.Ints(ints)
	fmt.Println("Ints:", ints)

	sorted := sort.IntsAreSorted(ints)
	fmt.Println("Sorted:", sorted)

	fruits := []string{"banana", "kiwi", "apple", "mango", "cucumber", "orange"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)

	lenOpsTotal := atomic.LoadUint64(&lenOps)
	fmt.Println("lenOpsTotal", lenOpsTotal)
	swapOpsTotal := atomic.LoadUint64(&swapOps)
	fmt.Println("swapOpsTotal", swapOpsTotal)
	lessOpsTotal := atomic.LoadUint64(&lessOps)
	fmt.Println("lessOpsTotal", lessOpsTotal)

}

func Clean() {
	fmt.Println("defered function")
}
