package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

func makeRandomSlice(numItems, max int) []int {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	slice := make([]int, numItems)

	for i := 0; i < numItems; i++ {
		slice[i] = random.Intn(max)
	}
	return slice
}

func printSlice(slice []int, numItems int) {
	if len(slice) <= numItems {
		fmt.Println(slice)
	} else {
		fmt.Println(slice[:numItems])
	}
}

// Perform binary search.
// Return the target's location in the slice and the number of tests.
// If the item is not found, return -1 and the number tests.
func binarySearch(slice []int, target int) (index, numTests int) {
	var L int
	R := len(slice)
	// first, check boundary cases
	if R == 0 || target > slice[R-1] || target < slice[0] {
		return -1, numTests
	}

	for L <= R {
		numTests++
		index = int(math.Floor(float64((L + R) / 2)))
		if slice[index] < target {
			L = index + 1
		} else if slice[index] > target {
			R = index - 1
		} else {
			return
		}
	}
	return -1, numTests
}

func quicksort(slice []int) {
	hi := len(slice)

	if hi <= 1 {
		return
	}

	p := partition(slice, 0, len(slice)-1)
	quicksort(slice[:p])
	quicksort(slice[p+1:])
}

func partition(slice []int, lo, hi int) int {
	pivot := slice[hi]

	i := lo - 1

	for j := lo; j < hi; j++ {
		if slice[j] <= pivot {
			i = i + 1
			slice[i], slice[j] = slice[j], slice[i]
		}
	}

	i = i + 1
	slice[i], slice[hi] = slice[hi], slice[i]
	return i
}

func main() {
	// Get the number of items and maximum item value.
	var numItems, max int
	var target string
	fmt.Printf("# Items: ")
	fmt.Scanln(&numItems)
	fmt.Printf("Max: ")
	fmt.Scanln(&max)

	// Make and display the unsorted slice.
	slice := makeRandomSlice(numItems, max)
	quicksort(slice)

	printSlice(slice, 40)
	fmt.Println()

	for {
		fmt.Printf("Target: ")
		_, err := fmt.Scanln(&target)
		if err != nil {
			break
		}

		if len(target) == 0 {
			break
		}

		ntarget, err := strconv.Atoi(target)
		if err != nil {
			fmt.Println(err)
			continue
		}

		i, n := binarySearch(slice, ntarget)
		if i == -1 {
			fmt.Printf("Target %d not found, %d tests\n", ntarget, n)
		} else {
			fmt.Printf("values[%d] = %d, %d tests\n", i, ntarget, n)
		}

	}
}
