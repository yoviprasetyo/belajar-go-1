package main

import "fmt"

// Filter to get something in slice.
type Filter func([]int) int

// FilterRecursive to get something in slice.
type FilterRecursive func([]int, int, int) int

func filterMax(slice []int) int {
	max := slice[0]
	for i := 0; i < len(slice); i++ {
		if max <= slice[i] {
			max = slice[i]
		}
	}
	return max
}

func filterMin(slice []int) int {
	min := slice[0]
	for i := 0; i < len(slice); i++ {
		if min >= slice[i] {
			min = slice[i]
		}
	}
	return min
}

func filterMinRecursive(slice []int, value, index int) int {
	if index == 0 {
		return value
	}
	if slice[index] <= value {
		value = slice[index]
	}
	return filterMinRecursive(slice, value, (index - 1))
}

func filterMaxRecursive(slice []int, value, index int) int {
	if index == 0 {
		return value
	}
	if slice[index] >= value {
		value = slice[index]
	}
	return filterMaxRecursive(slice, value, (index - 1))
}

func average(slice []int) int {
	total := sum(slice)
	return total / len(slice)
}

func sum(slice []int) int {
	total := 0
	for _, value := range slice {
		total += value
	}
	return total
}

func slicing(slices []int, amount int) [][]int {
	var chunks [][]int
	length := len(slices)
	divided := length / amount

	for i := 0; i < amount; i++ {
		iteration := i + 1
		finish := iteration * divided
		start := i * divided
		slice := slices[start:finish]
		chunks = append(chunks, slice)
	}
	return chunks
}

func getValue(slices []int, filter Filter) int {
	return filter(slices)
}

func getValueRecursive(slices []int, filter FilterRecursive) int {
	return filter(slices, slices[0], (len(slices) - 1))
}

func getMinTotal(slice []int) (int, int) {
	min := slice[0]
	key := 0
	for i := 0; i < len(slice); i++ {
		if min <= slice[i] {
			min = slice[i]
			key = i
		}
	}
	return key, min
}

func getMaxTotal(slice []int) (int, int) {
	max := slice[0]
	key := 0
	for i := 0; i < len(slice); i++ {
		if max >= slice[i] {
			max = slice[i]
			key = i
		}
	}
	return key, max
}

func main() {
	somethings := []int{23, 45, 67, 54, 66, 19, 56, 78, 89, 44, 11, 22, 33, 44, 55, 66, 77, 88, 99, 23, 34, 32, 23, 12}

	chunks := slicing(somethings, 3)
	totals := []int{
		0, 0, 0,
	}
	mins := [3]int{
		0, 0, 0,
	}
	maxs := [3]int{
		0, 0, 0,
	}
	averages := [3]int{
		0, 0, 0,
	}

	for i := 0; i < len(chunks); i++ {
		iteration := i + 1
		chunk := chunks[i]
		averages[i] = average(chunk)
		totals[i] = sum(chunk)
		mins[i] = getValueRecursive(chunk, filterMinRecursive)
		maxs[i] = getValueRecursive(chunk, filterMaxRecursive)
		fmt.Println("Kumpulan ke-", iteration, chunk, ". Rata-rata:", averages[i], ". Penjumlahan:", totals[i], ". Nilai Minimal:", mins[i], ". Nilai Maksimal", maxs[i])
	}

	keyMin, minTotal := getMinTotal(totals)
	fmt.Println("Total Kumpulan terkecil adalah", minTotal, "oleh kumpulan", chunks[keyMin])

	keyMax, maxTotal := getMaxTotal(totals)
	fmt.Println("Total Kumpulan terbesar adalah", maxTotal, "oleh kumpulan", chunks[keyMax])
}
