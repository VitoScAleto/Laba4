package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Алгоритм "Доверчивый"
func confiding(roundNumber int, selfChoices, enemyChoices []bool) bool {
	// Всегда сотрудничает, независимо от предыдущих ходов
	return true
}

// Алгоритм "Мстительный"
func revengeful(roundNumber int, selfChoices, enemyChoices []bool) bool {
	// Сотрудничает, если в предыдущем раунде противник также сотрудничал, иначе предает
	if roundNumber == 0 || enemyChoices[roundNumber-1] {
		return true
	}
	return false
}

// Алгоритм "Осторожный"
func careful(roundNumber int, selfChoices, enemyChoices []bool) bool {
	// Сотрудничает, если в предыдущих 3 раундах противник 2 раза или больше сотрудничал, иначе предает
	cooperationCount := 0
	for i := max(0, roundNumber-3); i < roundNumber; i++ {
		if enemyChoices[i] {
			cooperationCount++
		}
	}
	return cooperationCount >= 2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Генерируем случайное количество раундов от 100 до 200
	rand.Seed(time.Now().UnixNano())
	totalRounds := rand.Intn(101) + 100

	// Массивы для хранения выборов игроков
	player1Choices := make([]bool, totalRounds)
	player2Choices := make([]bool, totalRounds)

	// Выбор алгоритмов для игроков
	player1Algorithm := revengeful
	player2Algorithm := confiding

	// Игровой цикл
	player1Score, player2Score := 0, 0
	for round := 0; round < totalRounds; round++ {
		player1Choices[round] = player1Algorithm(round, player1Choices, player2Choices)
		player2Choices[round] = player2Algorithm(round, player2Choices, player1Choices)

		if player1Choices[round] && player2Choices[round] {
			player1Score += 24
			player2Score += 24
		} else if player1Choices[round] && !player2Choices[round] {
			player2Score += 20
		} else if !player1Choices[round] && player2Choices[round] {
			player1Score += 20
		} else {
			player1Score += 4
			player2Score += 4
		}
	}

	// Вывод результатов
	fmt.Println("Количество раундов:", totalRounds)
	fmt.Println("Счет игрока 1:", player1Score)
	fmt.Println("Счет игрока 2:", player2Score)
}
