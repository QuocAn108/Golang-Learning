package main

import "fmt"

func mapInt(arr []int, f func(int) int) []int {
	result := make([]int, len(arr))
	for i, v := range arr {
		result[i] = f(v)
	}
	return result
}
func mapAny[K any, V any](arr []K, f func(K) V) []V {
	result := make([]V, len(arr))
	for i, v := range arr {
		result[i] = f(v)
	}
	return result
}
func main() {
	arr := []int{1, 2, 3, 4, 5}
	result := mapInt(arr, func(v int) int {
		return v * 2
	})
	fmt.Println(result)

	arr2 := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	rs := mapAny(arr2, func(v float64) float64 {
		return v * 2
	})
	fmt.Println(rs)

	arr3 := []string{"an", "binh", "phuoc", "tinh", "yeu"}
	rs1 := mapAny(arr3, func(v string) string {
		return v + "|"
	})
	fmt.Println(rs1)

}
