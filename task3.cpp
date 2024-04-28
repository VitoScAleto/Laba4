#include <iostream>
#include <vector>
#include <random>
#include <cmath>
#include <locale.h>
using namespace std;

void fillingTheVector(vector<int>& vec) //заполнение вектора случайными числами
{ 
    random_device rd;//random_device, который является источником недетерминированных случайных чисел.
    //Затем мы используем это устройство для заполнения генератора std::minstd_rand. Затем функция генератора() используется для генерации случайного числа
    minstd_rand generator(rd());
    uniform_int_distribution<int> distribution(0, 100);// функция destribition для задания диапозона значений

    for (int& item : vec) 
    {
        item = distribution(generator);
    }
}

int ElementCountInVector(const vector<int>& container, int element) //подсчет количества вхождений заданного элемента в вектор
{ 
    int count = 0;
    for (const int& value : container) 
    {
        if (value == element) count++;
    }
    return count;
}

double mathExpectation(vector<int> container) //вычисление реального математического ожидания
{ 
    float Mx = 0.0;
    for (auto i = 0; i < container.size(); i++)
    {
        Mx += container[i] * 0.01;// диапозон чисел 100. Каждое равновероятно
    }
    return Mx;
}

double mathExp(vector <int> container) //вычисление ожидаемого математического ожидания
{
    float Mx0 = 0.0;
    Mx0 = container.size()/2;//  при равномерном распределении случайной величины в диапазоне от a до b, математическое ожидание вычисляется как среднее арифметическое этих границ
    return Mx0;
}

double mathDispersion(vector<int> container) //функция для вычисления дисперсии
{
    float Mx = mathExpectation(container);// вычисляем дисперсию относительно реального мат ожидания
    float sum = 0.0;
    for (int i = 0; i < container.size(); i++)
    {
        sum += pow(container[i] - Mx, 2);// суммируем разницу между всеми значениями и мат ожиданием в квадрате
    }
    return sum / (container.size()-1);//делим на количество в выборке- 1
}

double tabLaplas(double x) //функции Лапласа (стандартного нормального распределения) для заданного аргумента x
{ 
    return (1.0 / 2.0) * (1.0 + erf(x / sqrt(2.0)));
}

double chiSquare(const vector<int>& container)
{
    double Mx = mathExpectation(container);
    double chi = 0.0;
    for (int i = 0; i < 10; i++) // Перебор 10 интервалов
    {
        int observed = ElementCountInVector(container, i * 10); // Наблюдаемая частота
        double expected = container.size() * 0.1; // Ожидаемая частота

        chi += pow(observed - expected, 2) / expected;
    }
    return chi;
}
double quantile_chi_square(double p, int df)
{
    // Функция для вычисления квантили распределения хи-квадрат
    // Реализация на основе библиотеки math.h
    return sqrt(2.0 * tgamma(df / 2.0) * pow(p, 1.0 / df) / tgamma((df - 1.0) / 2.0));
}
double chiSquareCriticalValue(int df, double alpha) 
{
    // Функция для вычисления критического значения хи-квадрат
    double p = 1.0 - alpha; // Вероятность, соответствующая уровню значимости
    return quantile_chi_square(p, df); // Вычисление критического значения
}



int main() 
{
    setlocale(LC_ALL, "Rus");

    vector<int> vector50(50);
    fillingTheVector(vector50);
   
    vector<int> vector100(100);
    fillingTheVector(vector100);

    vector<int> vector1000(1000);
    fillingTheVector(vector1000);
   
    double criticalValue = chiSquareCriticalValue(7, 0.05); // Критическое значение хи-квадрат для 99 степеней свободы и уровня значимости 0.05
   
   // для 50 элементов
    double chiSquareValue = chiSquare(vector50);
    double meanExpected = mathExpectation(vector50); //реальное математическое ожидание
    double meanObserved = mathExp(vector50);// ожидаемое мат ожидание

    cout << "X^2: " << chiSquareValue << endl;

    if (chiSquareValue <= criticalValue) 
    {
        cout << "Гипотеза о нормальном распределении не отвергается." << endl;
    }
    else 
    {
        cout << "Гипотеза о нормальном распределении отвергается." << endl;
    }

    cout << "Ожидаемое математическое ожидание: " << meanObserved << endl;
    cout << "Реальное математическое ожидание: " << meanExpected << endl;
    cout << endl;

    //для 100 элементов
    chiSquareValue = chiSquare(vector100);
    meanExpected = mathExpectation(vector100);
    meanObserved = mathExp(vector100);

    cout << "X^2: " << chiSquareValue << endl;

    if (chiSquareValue <= criticalValue) 
    {
        cout << "Гипотеза о нормальном распределении не отвергается." << endl;
    }
    else 
    {
        cout << "Гипотеза о нормальном распределении отвергается." << endl;
    }

    cout << "Ожидаемое математическое ожидание: " << meanObserved << endl;
    cout << "Реальное математическое ожидание: " << meanExpected << endl;
    cout << endl;

    //для 1000 элементов
    chiSquareValue = chiSquare(vector1000);
    meanExpected = mathExpectation(vector1000);
    meanObserved = mathExp(vector1000);

    cout << "X^2: " << chiSquareValue << endl;

    if (chiSquareValue <= criticalValue) 
    {
        cout << "Гипотеза о нормальном распределении не отвергается." << endl;
    }
    else 
    {
        cout << "Гипотеза о нормальном распределении отвергается." << endl;
    }

    cout << "Ожидаемое математическое ожидание: " << meanObserved << endl;
    cout << "Реальное математическое ожидание: " << meanExpected << endl;
    cout << endl;

    return 0;
}