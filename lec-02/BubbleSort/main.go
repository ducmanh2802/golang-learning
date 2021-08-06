package main

import "fmt"

func bubleSort([] int arr) []int {
	for i := 0; i < len(arr); i++ {
		for j := i +1; j <  len(arr); j++ {
				if arr[i] > arr[j] {
					temp : = arr[j]
					arr[j] = arr[i]
					arr[i] = temp
				}
		}
	}
}
func main() {
	a:=[3] int {3,2,5}
	fmt.Println(bubleSort(a))
}
