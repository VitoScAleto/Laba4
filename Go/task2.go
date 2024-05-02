package main

import (
	"fmt"
	"math/rand"
	"time"
)

func GetRandomNumber(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

func task2_2and2_3(arr1, arr2 []int, sizeArr int) {
	fmt.Println("Task 2_2 and 2_3\n")

	for i := range arr1 {
		arr1[i] = GetRandomNumber(0, 100)
		fmt.Print(arr1[i], " ")
	}

	maxElement, indexMax := 0, 0
	minElement, indexMin := 100, 0
	sum := 0
	counterForArr2 := 0

	for i := range arr1 {
		if arr1[i] > maxElement {
			maxElement = arr1[i]
			indexMax = i
		}
		if arr1[i] < minElement {
			minElement = arr1[i]
			indexMin = i
		}
	}
	fmt.Printf("\nMax element = %d| %d\nMin Element = %d| %d\n", maxElement, indexMax, minElement, indexMin)

	if indexMax < indexMin {
		for i := indexMax + 1; i < indexMin; i++ {
			sum += arr1[i]
			arr2[counterForArr2] = arr1[i]
			counterForArr2++
		}
	} else {
		for i := indexMin + 1; i < indexMax; i++ {
			sum += arr1[i]
			arr2[counterForArr2] = arr1[i]
			counterForArr2++
		}
	}

	fmt.Println("Сумма элементов равна", sum)

	for ; counterForArr2 < sizeArr; counterForArr2++ {
		arr2[counterForArr2] = arr1[GetRandomNumber(0, sizeArr-1)]
	}

	for _, i := range arr2 {
		fmt.Print(i, " ")
	}
	fmt.Println()
}

func task2_4(size int) {
	fmt.Println("\nTask 2_4\n")
	arrChar := make([]rune, size)
	for i := range arrChar {
		arrChar[i] = rune(GetRandomNumber(65, 122))
		fmt.Print(string(arrChar[i]), " ")
	}
	fmt.Println()

	for i := 0; i < len(arrChar); i += 2 {
		minIndex := i
		for j := i + 2; j < len(arrChar); j += 2 {
			if arrChar[j] < arrChar[minIndex] {
				minIndex = j
			}
		}
		arrChar[i], arrChar[minIndex] = arrChar[minIndex], arrChar[i]
	}

	for _, item := range arrChar {
		fmt.Print(string(item), " ")
	}
}

func task2_5(size int) {
	fmt.Println("\n\nTask 2_5\n")

	Array := make([]int, size)
	NewArray := make([]int, size)

	for i := range Array {
		Array[i] = GetRandomNumber(100, 900)
		fmt.Print(Array[i], " ")
	}

	fmt.Println("\nОтсортированный массив")
	sort(Array)

	for _, i := range Array {
		fmt.Print(i, " ")
	}
	fmt.Println()

	for i := range NewArray {
		NewArray[i] = Array[GetRandomNumber(0, size-1)]
	}

	for i := range NewArray {
		if NewArray[i] < Array[i] {
			NewArray[i] = 0
		}
	}

	for _, item := range NewArray {
		fmt.Print(item, " ")
	}
}

func sort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

func main() {
	sizeArr := 20

	arr1 := make([]int, sizeArr)
	arr2 := make([]int, sizeArr)

	task2_2and2_3(arr1, arr2, sizeArr)
	task2_4(sizeArr)
	task2_5(sizeArr)
}
