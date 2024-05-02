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
		Mx += float64(i) * 0.01 // диапозон чисел 100. Каждое равновероятно
	}
	return Mx
}

func mathExp(container []int) float64 { // вычисление ожидаемого математического ожидания
	return float64(len(container)) / 2 // при равномерном распределении случайной величины в диапазоне от a до b, математическое ожидание вычисляется как среднее арифметическое этих границ
}

func mathDispersion(container []int) float64 { // функция для вычисления дисперсии
	Mx := mathExpectation(container) // вычисляем дисперсию относительно реального мат ожидания
	var sum float64
	for _, i := range container {
		sum += math.Pow(float64(i)-Mx, 2) // суммируем разницу между всеми значениями и мат ожиданием в квадрате
	}
	return sum / float64(len(container)-1) // делим на количество в выборке- 1
}

func tabLaplas(x float64) float64 { // функции Лапласа (стандартного нормального распределения) для заданного аргумента x
	return 0.5 * (1 + math.Erf(x/math.Sqrt(2)))
}

func chiSquare(container []int) float64 {

	var chi float64
	for i := 0; i < 10; i++ { // Перебор 10 интервалов
		observed := elementCountInVector(container, i*10) // Наблюдаемая частота
		expected := float64(len(container)) * 0.1         // Ожидаемая частота

		chi += math.Pow(float64(observed)-expected, 2) / expected
	}
	return chi
}

func quantileChiSquare(p float64, df int) float64 {
	// Функция для вычисления квантили распределения хи-квадрат
	// Реализация на основе math.Gamma и math.Pow
	return math.Sqrt(2.0 * math.Gamma(float64(df)/2.0) * math.Pow(p, 1.0/float64(df)) / math.Gamma((float64(df)-1.0)/2.0))
}

func chiSquareCriticalValue(df int, alpha float64) float64 {
	// Функция для вычисления критического значения хи-квадрат
	p := 1.0 - alpha                // Вероятность, соответствующая уровню значимости
	return quantileChiSquare(p, df) // Вычисление критического значения
}

func main() {
	rand.Seed(time.Now().UnixNano())

	vector50 := make([]int, 50)
	fillingTheVector(vector50)

	vector100 := make([]int, 100)
	fillingTheVector(vector100)

	vector1000 := make([]int, 1000)
	fillingTheVector(vector1000)

	criticalValue := chiSquareCriticalValue(7, 0.05) // Критическое значение хи-квадрат для 99 степеней свободы и уровня значимости 0.05

	// для 50 элементов
	chiSquareValue := chiSquare(vector50)
	meanExpected := mathExpectation(vector50) // реальное математическое ожидание
	meanObserved := mathExp(vector50)         // ожидаемое мат ожидание

	fmt.Println("X^2:", chiSquareValue)

	if chiSquareValue <= criticalValue {
		fmt.Println("Гипотеза о нормальном распределении не отвергается.")
	} else {
		fmt.Println("Гипотеза о нормальном распределении отвергается.")
	}

	fmt.Println("Ожидаемое математическое ожидание:", meanObserved)
	fmt.Println("Реальное математическое ожидание:", meanExpected)
	fmt.Println()

	// для 100 элементов
	chiSquareValue = chiSquare(vector100)
	meanExpected = mathExpectation(vector100)
	meanObserved = mathExp(vector100)

	fmt.Println("X^2:", chiSquareValue)

	if chiSquareValue <= criticalValue {
		fmt.Println("Гипотеза о нормальном распределении не отвергается.")
	} else {
		fmt.Println("Гипотеза о нормальном распределении отвергается.")
	}

	fmt.Println("Ожидаемое математическое ожидание:", meanObserved)
	fmt.Println("Реальное математическое ожидание:", meanExpected)
	fmt.Println()

	// для 1000 элементов
	chiSquareValue = chiSquare(vector1000)
	meanExpected = mathExpectation(vector1000)
	meanObserved = mathExp(vector1000)

	fmt.Println("X^2:", chiSquareValue)

	if chiSquareValue <= criticalValue {
		fmt.Println("Гипотеза о нормальном распределении не отвергается.")
	} else {
		fmt.Println("Гипотеза о нормальном распределении отвергается.")
	}

	fmt.Println("Ожидаемое математическое ожидание:", meanObserved)
	fmt.Println("Реальное математическое ожидание:", meanExpected)
	fmt.Println()
}
