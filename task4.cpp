#include <iostream>
#include <vector>
#include <random>
#include <locale.h>
using namespace std;

// Алгоритм "Доверчивый"
bool confiding(int round_number, vector<bool> self_choices, vector<bool> enemy_choices) 
{
    // Всегда сотрудничает, независимо от предыдущих ходов
    return true;
}

// Алгоритм "Мстительный"
bool revengeful(int round_number, vector<bool> self_choices, vector<bool> enemy_choices) 
{
    // Сотрудничает, если в предыдущем раунде противник также сотрудничал, иначе предает
    if (round_number == 0 || enemy_choices[round_number - 1]) 
    {
        return true;
    }
    else {
        return false;
    }
}

// Алгоритм "Осторожный"
bool careful(int round_number, vector<bool> self_choices, vector<bool> enemy_choices) 
{
    // Сотрудничает, если в предыдущих 3 раундах противник 2 раза или больше сотрудничал, иначе предает
    int cooperationCount = 0;
    for (int i = max(0, round_number - 3); i < round_number; i++) 
    {
        if (enemy_choices[i]) 
        {
            cooperationCount++;
        }
    }
    return cooperationCount >= 2;
}

int main() 
{
    setlocale(LC_ALL, "Rus");
    // Генерируем случайное количество раундов от 100 до 200
    random_device rd;
    mt19937 gen(rd());
    uniform_int_distribution<> distribution(100, 200);
    int total_rounds = distribution(gen);

    // Массивы для хранения выборов игроков
    vector<bool> player1_choices(total_rounds, false);
    vector<bool> player2_choices(total_rounds, false);

    // Выбор алгоритмов для игроков
    bool (*player1_algorithm)(int, vector<bool>, vector<bool>) = &revengeful;
    bool (*player2_algorithm)(int, vector<bool>, vector<bool>) = &confiding;

    // Игровой цикл
    int player1_score = 0, player2_score = 0;
    for (int round = 0; round < total_rounds; round++) {
        player1_choices[round] = (*player1_algorithm)(round, player1_choices, player2_choices);
        player2_choices[round] = (*player2_algorithm)(round, player2_choices, player1_choices);

        if (player1_choices[round] && player2_choices[round]) 
        {
            player1_score += 24;
            player2_score += 24;
        }
        else if (player1_choices[round] && !player2_choices[round]) 
        {
            player2_score += 20;
        }
        else if (!player1_choices[round] && player2_choices[round]) 
        {
            player1_score += 20;
        }
        else 
        {
            player1_score += 4;
            player2_score += 4;
        }
    }

    // Вывод результатов
    cout << "Количество раундов: " << total_rounds << endl;
    cout << "Счет игрока 1: " << player1_score << endl;
    cout << "Счет игрока 2: " << player2_score << endl;

    return 0;
}