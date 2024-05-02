package main

import (
	"fmt"
)

func main() {
	var A, B, C, X0, Xi int
	var arr1 []int

	fmt.Println("Введите диапозон значений С:")
	fmt.Print("C = ")
	fmt.Scan(&C)

	fmt.Println("Введите множитель A(0<=A<=C):")
	fmt.Print("A = ")
	fmt.Scan(&A)

	fmt.Println("Введите инкрементируещее значений B(0<=B<=C):")
	fmt.Print("B = ")
	fmt.Scan(&B)

	fmt.Println("Введите начальное значений X0(0<=X0<=C):")
	fmt.Print("X0 = ")
	fmt.Scan(&X0)

	arr1 = append(arr1, X0)
	for i := 1; i < C; i++ {
		Xi = (A*arr1[i-1] + B) % C
		arr1 = append(arr1, Xi)
	}

	for _, i := range arr1 {
		fmt.Print(i, " ")
	}
	fmt.Println()

	for i := 1; i < len(arr1); i++ {
		if arr1[0] == arr1[i] {
			fmt.Println("Индекс начала повторяющейся последовательности =", i+1)
			break
		}
	}
}
