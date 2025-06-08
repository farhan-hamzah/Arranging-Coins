package main

import (
	"fmt"
)

func arrangeCoins(n int) int {
	left, right := 0, n
	for left <= right {
		mid := (left + right) / 2
		sum := mid * (mid + 1) / 2
		if sum == n {
			return mid
		} else if sum < n {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return right
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah koin: ")
	fmt.Scan(&n)

	result := arrangeCoins(n)
	fmt.Printf("Jumlah baris penuh yang bisa dibentuk: %d\n", result)
}
