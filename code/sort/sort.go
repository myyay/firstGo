package main

import "fmt"

var arr = []int{9, 1, 3, 18, 17, 15, 11, 7, 8, 4, 21, 12, 16, 1, 0, 20, 10, 4, 2, 17, 5}

//key = 9
//var arr = []int{5, 1, 3, 18, 17, 15, 11, 7, 8, 4, 21, 12, 16, 1, 0, 20, 10, 4, 2, 17, 5}
//var arr = []int{5, 1, 3, 18, 17, 15, 11, 7, 8, 4, 21, 12, 16, 1, 0, 20, 10, 4, 2, 17, 18}
//var arr = []int{5, 1, 3, 2, 17, 15, 11, 7, 8, 4, 21, 12, 16, 1, 0, 20, 10, 4, 2, 17, 18}
//var arr = []int{5, 1, 3, 2, 17, 15, 11, 7, 8, 4, 21, 12, 16, 1, 0, 20, 10, 4, 17, 17, 18}
//var arr = []int{5, 1, 3, 2, 4, 15, 11, 7, 8, 4, 21, 12, 16, 1, 0, 20, 10, 4, 17, 17, 18}
//var arr = []int{5, 1, 3, 2, 4, 15, 11, 7, 8, 4, 21, 12, 16, 1, 0, 20, 10, 15, 17, 17, 18}

func quickSort(low int, high int, arr []int) {
	if low >= high {
		return
	}

	mid := partition(low, high, arr)
	quickSort(low, mid-1, arr)
	quickSort(mid+1, high, arr)

}

func partition(low int, high int, arr []int) int {
	key := arr[low]

	for low < high {
		for key <= arr[high] && low < high {
			high--
		}
		arr[low] = arr[high]
		for key >= arr[low] && low < high {
			low++
		}
		arr[high] = arr[low]
	}
	arr[high] = key
	return high

}

func binarySearch(v int, arr []int) int {
	low := 0
	high := len(arr)

	for low < high {
		mid := low + (high-low)/2

		if arr[mid] == v {
			return mid
		} else if v > arr[mid] {
			low = mid + 1
		} else if v < arr[mid] {
			high = mid - 1
		}
	}
	return -1

}

func main() {

	quickSort(0, len(arr)-1, arr)

	for _, v := range arr {
		fmt.Printf("%d\t", v)
	}

	fmt.Println()
	for i, _ := range arr {
		fmt.Printf("%d\t", i)
	}
	fmt.Println()

	fmt.Println(binarySearch(21, arr))

}
