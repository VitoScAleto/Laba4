#include <iostream>
#include <cmath>
#include <iomanip>
using namespace std;

void PrintfTablHalfDivisionMethod(double LeftPoint, double RightPoint, int N)
{
    if (N == 1)
    {
        cout << "N";
        cout << "\t     a ";
        cout << "\t\t\t     b ";
        cout << "\t\t\t     b-a ";
        cout << endl;
        cout << "----------------------------------------------------------------------------------";
        cout << endl;
    }
    cout << N<<" |";
    cout <<"\t" << setw(10) << LeftPoint;
    cout << "\t\t" << setw(10) << RightPoint;
    cout << "\t\t" << setw(12)<< RightPoint - LeftPoint;
    cout << endl;
}
double function(double x) {
    return 2*x+cos(x); // Заменить функцией, корни которой мы ищем
}

// a, b - пределы хорды, epsilon — необходимая погрешность
double findSolution(double a, double b, double epsilon) {
    while (abs(b - a) > epsilon) {
        a = a - (b - a) * function(a) / (function(b) - function(a));
        b = b - (a - b) * function(b) / (function(a) - function(b));
    }
    // a, b — (i - 1)-й и i-й члены

    return b;
}
double HalfDivisionMethod(double LeftPoint, double RightPoint,double epsilon)
{
    int iteration = 1;
    double midPoint = 0.0;
    if (function(LeftPoint) * function(RightPoint) < 0)// проверка на разность знаков функции на концах отрезка
    {
        while (abs(RightPoint - LeftPoint) > epsilon)// пока интервал больше погрешности
        {
            midPoint = (RightPoint + LeftPoint) / 2;
            PrintfTablHalfDivisionMethod(LeftPoint, RightPoint, iteration);
            if (function(LeftPoint) * function(midPoint) < 0)  RightPoint = midPoint;// если функция имеет разные знаки, то правая точка середина отрезка
            else LeftPoint = midPoint;
            //midPoint = (RightPoint + LeftPoint) / 2;
            iteration++;
        }
    } 
    else 
    {
        cout << "Интервал выбран неверно. Функция имеет одинаковые знаки на концах отрезка" << endl;
    }
    return midPoint;
}

 

int main(void)
{
    setlocale(LC_ALL, "Rus");
    float left = 15;
    float right = -10;
    float eps = 0.0001;
    cout << "Корень через метод хорд:\n" << findSolution(left, right, eps) << endl;
    cout << "Корень через метод половинного деления:\n" << HalfDivisionMethod(left, right, eps);



    return 0;
}
