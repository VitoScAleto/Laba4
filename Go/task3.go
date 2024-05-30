package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func fillingTheVector(vec []int) { // заполнение вектора случайными числами
	rand.Seed(time.Now().UnixNano())
	for i := range vec {
		vec[i] = rand.Intn(101)
	}
}

func elementCountInVector(container []int, element int) int { // подсчет количества вхождений заданного элемента в вектор
	count := 0
	for _, value := range container {
		if value == element {
			count++
		}
	}
	return count
}

func mathExpectation(container []int) float64 { // вычисление реального математического ожидания
	var Mx float64
	for _, i := range container {
		Mx += float64(i) * 0.01 // диапазон чисел 100. Каждое равновероятно
	}
	return Mx
}

func mathExp(container []int) float64 { // вычисление ожидаемого математического ожидания
	var Mx0 float64
	Mx0 = float64(len(container)) / 2 // при равномерном распределении случайной величины в диапазоне от a до b, математическое ожидание вычисляется как среднее арифметическое этих границ
	return Mx0
}

func mathDispersion(container []int) float64 { // функция для вычисления дисперсии
	Mx := mathExpectation(container) // вычисляем дисперсию относительно реального мат ожидания
	var sum float64
	for _, i := range container {
		sum += math.Pow(float64(i)-Mx, 2) // суммируем разницу между всеми значениями и мат ожиданием в квадрате
	}
	return sum / float64(len(container)-1) // делим на количество в выборке- 1
}

func chiSquare(container []int) float64 {
	Mx := mathExpectation(container)
	var chi float64
	for i := 0; i < 10; i++ { // Перебор 10 интервалов
		observed := elementCountInVector(container, i*10) // Наблюдаемая частота
		expected := float64(len(container)) * 0.1         // Ожидаемая частота

		chi += math.Pow(float64(observed)-expected, 2) / expected
	}
	return chi
}

func main() {
	vector50 := make([]int, 50)
	fillingTheVector(vector50)

	vector100 := make([]int, 100)
	fillingTheVector(vector100)

	vector1000 := make([]int, 1000)
	fillingTheVector(vector1000)

	var criticalValue float64

	// для 50 элементов
	criticalValue = 66.338 // Критическое значение хи-квадрат для 49 степеней свободы и уровня значимости 0.05
	fmt.Println(criticalValue)
	chiSquareValue := chiSquare(vector50)
	meanExpected := mathExpectation(vector50) // реальное математическое ожидание
	meanObserved := mathExp(vector50)         // ожидаемое мат ожидание

	fmt.Printf("X^2: %f\n", chiSquareValue)

	if chiSquareValue <= criticalValue {
		fmt.Println("Гипотеза о нормальном распределении не отвергается.")
	} else {
		fmt.Println("Гипотеза о нормальном распределении отвергается.")
	}

	fmt.Printf("Ожидаемое математическое ожидание: %f\n", meanObserved)
	fmt.Printf("Реальное математическое ожидание: %f\n", meanExpected)
	fmt.Println()

	// для 100 элементов
	criticalValue = 124.342
	fmt.Println(criticalValue)
	chiSquareValue = chiSquare(vector100)
	meanExpected = mathExpectation(vector100)
	meanObserved = mathExp(vector100)

	fmt.Printf("X^2: %f\n", chiSquareValue)

	if chiSquareValue <= criticalValue {
		fmt.Println("Гипотеза о нормальном распределении не отвергается.")
	} else {
		fmt.Println("Гипотеза о нормальном распределении отвергается.")
	}

	fmt.Printf("Ожидаемое математическое ожидание: %f\n", meanObserved)
	fmt.Printf("Реальное математическое ожидание: %f\n", meanExpected)
	fmt.Println()

	// для 1000 элементов
	criticalValue = 1092.32
	fmt.Println(criticalValue)
	chiSquareValue = chiSquare(vector1000)
	meanExpected = mathExpectation(vector1000)
	meanObserved = mathExp(vector1000)

	fmt.Printf("X^2: %f\n", chiSquareValue)

	if chiSquareValue <= criticalValue {
		fmt.Println("Гипотеза о нормальном распределении не отвергается.")
	} else {
		fmt.Println("Гипотеза о нормальном распределении отвергается.")
	}

	fmt.Printf("Ожидаемое математическое ожидание: %f\n", meanObserved)
	fmt.Printf("Реальное математическое ожидание: %f\n", meanExpected)
	fmt.Println()
}
