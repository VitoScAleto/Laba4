#include <iostream>
#include <random>
#include <vector>

using namespace std;

int GetRandomNumber(int min, int max)
{
    random_device rd;//random_device, который является источником недетерминированных случайных чисел.
    //Затем мы используем это устройство для заполнения генератора std::minstd_rand. Затем функция генератора() используется для генерации случайного числа
    minstd_rand generator(rd());

    uniform_int_distribution<int> distribution(min, max);// функция destribition для задания диапозона значений
    int random_number = distribution(generator);
    return random_number;
}

void task2_2(vector<int>copyArr1)
{
    int max_element = 0, index_max = 0;
    int min_element = 100, index_min = 0;
    for (int i = 0; i < copyArr1.size(); i++)
    {
        if (copyArr1[i] > max_element)
        {
            max_element = copyArr1[i];
            index_max = i;
        }
        if (copyArr1[i] < min_element) min_element = copyArr1[i];
        {
            index_min = i;
        }
    }
    cout << "\nMax element = " << max_element <<"\nMin Element = " << min_element << endl;
    for(;)
}


int main(void)
{
    int sizeArr1 = 20;

    vector <int> arr1(sizeArr1);

    for (auto& i : arr1)
    {
        i = GetRandomNumber(0, 100);
    }
    task2_2(arr1);



    return 0;
}