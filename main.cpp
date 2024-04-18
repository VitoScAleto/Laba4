#include<iostream>
#include<cmath>

using namespace std;

#include <iostream>
#include <math.h>

double f(double x) {
    return 2*x+cos(x); // Заменить функцией, корни которой мы ищем
}

// a, b - пределы хорды, epsilon — необходимая погрешность
double findRoot(double a, double b, double epsilon) {
    while (fabs(b - a) > epsilon) {
        a = a - (b - a) * f(a) / (f(b) - f(a));
        b = b - (a - b) * f(b) / (f(a) - f(b));
    }
    // a, b — (i - 1)-й и i-й члены

    return b;
}

int main(void)
{
    float xyrav = 0.0;
    float left = 1;
    float rigth = 2;
    float eps = 0.0001;
    xyrav = findRoot(left, rigth, eps);
    std::cout << xyrav;





    return 0;
}
