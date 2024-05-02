package main

import (
	"fmt"
	"math"
)

func InputTable(leftPoint, rightPoint float64, N int) {
	if N == 1 {
		fmt.Println("N\t     a \t\t\t     b \t\t\t     b-a")
		fmt.Println("----------------------------------------------------------------------------------")
	}
	fmt.Printf("%d |\t%.10f\t\t%.10f\t\t%.12f\n", N, leftPoint, rightPoint, rightPoint-leftPoint)
}

func InputNewtonMethod(leftPoint, rightPoint float64, N int) {
	if N == 1 {
		fmt.Println("N\t     Xn \t\t\t     Xn+1 \t\t\t     Xn+1-Xn")
		fmt.Println("----------------------------------------------------------------------------------")
	}
	fmt.Printf("%d |\t%.10f\t\t%.10f\t\t%.12f\n", N, leftPoint, rightPoint, rightPoint-leftPoint)
}

func function(x float64) float64 {
	return 2*x + math.Cos(x) // Заменить функцией, корни которой мы ищем
}

func df(x float64) float64 {
	return 2 - math.Sin(x)
}

func chordMethod(a, b, epsilon float64) float64 {
	counter := 1
	for math.Abs(b-a) > epsilon { // пока отрезок больше епсилон
		InputTable(a, b, counter)
		a = a - (b-a)*function(a)/(function(b)-function(a))
		b = b - (a-b)*function(b)/(function(a)-function(b))
		counter++
	}
	// a, b — (i - 1)-й и i-й члены
	fmt.Println("\nКорень через метод хорд =", b)
	return b
}

func HalfDivisionMethod(leftPoint, rightPoint, epsilon float64) float64 {
	iteration := 1
	var midPoint float64
	if function(leftPoint)*function(rightPoint) < 0 { // проверка на разность знаков функции на концах отрезка
		for math.Abs(rightPoint-leftPoint) > epsilon { // пока интервал больше погрешности
			midPoint = (rightPoint + leftPoint) / 2
			InputTable(leftPoint, rightPoint, iteration)
			if function(leftPoint)*function(midPoint) < 0 {
				rightPoint = midPoint // если функция имеет разные знаки, то правая точка середина отрезка
			} else {
				leftPoint = midPoint
			}
			iteration++
		}
	} else {
		fmt.Println("Интервал выбран неверно. Функция имеет одинаковые знаки на концах отрезка")
	}
	fmt.Println("\nКорень через метод половинного деления =", midPoint)
	return midPoint
}

func findGraficalSolution(left, right *float64) {
	for x := -1.0; x < 5; x += 0.01 {
		if math.Ceil(function(x)) == 0 {
			*left = x - 1.0
			*right = x + 1.0
		}
	}
}

func NewtownMethod(x0, epsilon float64) float64 {
	var x float64
	for i := 1; math.Abs(function(x0)) >= epsilon && i < 10; i++ {
		x = x0 - function(x0)/df(x0)
		InputNewtonMethod(x, x0, i)
		x0 = x
	}
	fmt.Println("\nКорень через метод Ньютона =", x)
	return x
}

func main() {
	var left, right float64
	x0 := 10.0
	eps := 0.0001

	findGraficalSolution(&left, &right) // отделяем корни

	fmt.Println("Метод Ньютона")
	NewtownMethod(x0, eps)
	fmt.Println("\nМетод половинного деления")
	HalfDivisionMethod(left, right, eps)
	fmt.Println("\nМетод хорд")
	chordMethod(left, right, eps)
}
