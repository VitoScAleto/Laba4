#include <iostream>
#include <random>
#include <vector>
#include <algorithm>

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

void task2_2and2_3(vector <int>& arr1, vector <int>& arr2, int sizeArr)
{
    cout << "Task 2_2 and 2_3\n\n";

    for (auto& i : arr1)
    {
        i = GetRandomNumber(0, 100);
        cout << i << " ";
    }

    int max_element = 0, index_max = 0, sum=0;
    int min_element = 100, index_min = 0;
    int counterForArr2 = 0;

    for (int i = 0; i < arr1.size(); i++)
    {
        if (arr1[i] > max_element)
        {
            max_element = arr1[i];
            index_max = i;
        }
        if (arr1[i] < min_element) 
        {
            min_element = arr1[i];
            index_min = i;
        }
    }
    cout << "\nMax element = " << max_element<<"| " << index_max << "\nMin Element = " << min_element <<"| " << index_min<< endl;
    if (index_max < index_min)
    {
        for (int i = index_max+1; i <index_min; i++, counterForArr2++)
        {
            sum += arr1[i];
            arr2[counterForArr2] = arr1[i];
        }
    }
    else
    {
        for (int i = index_min + 1; i < index_max; i++, counterForArr2++)
        {
            sum += arr1[i];
            arr2[counterForArr2] = arr1[i];
        }

    }
    cout << "Сумма элементов равна " << sum << endl;
    for (; counterForArr2 < sizeArr; counterForArr2++)
    { 
        arr2[counterForArr2] = arr1[GetRandomNumber(0, sizeArr - 1)];
    }
    for (auto i : arr2)
    {
        cout << i << " ";
    }
    cout << endl;
}

void task2_4(int size)// сортировка элементов на четных места по возрастанию
{       
    cout << "\nTask 2_4\n\n";
    vector<char> arrChar(size);
    for (auto& i : arrChar)
    {
        int sumvol = GetRandomNumber(65, 122);
        i = static_cast<char>(sumvol);
        cout << i << " ";
    }
    cout << endl;
    for (size_t i = 0; i < arrChar.size(); i += 2) {
        size_t minIndex = i;
        for (size_t j = i + 2; j < arrChar.size(); j += 2) {
            if (arrChar[j] < arrChar[minIndex]) {
                minIndex = j;
            }
        }
        std::swap(arrChar[i], arrChar[minIndex]);
    }
    for (auto item : arrChar)
    {
        cout << item << " ";
    }





}

void task2_5(int size)
{
    cout << "\n\nTask 2_5\n\n";

    vector<int> Array(size);
    vector<int> NewArray(size);

    for (auto& i : Array)
    {
        i = GetRandomNumber(100, 900);
        cout << i << " ";
    }

    cout << endl;

    sort(Array.begin(), Array.end());

    cout << "\nОтсортированный массив" << endl;

    for (auto& i : Array)// вывод отсортированного массива
    {
        cout << i << " ";
    }
    cout <<"\n" << endl;
   
    for (auto& item : NewArray)//заполняем новый массив случайными элементами из первого
    {
        item = Array[GetRandomNumber(0, size - 1)];
    }
    for (int i = 0; i < size; i++)
    {
        if (NewArray[i] < Array[i]) NewArray[i] = 0;//если в новом массив элемент меньше исходного, то записываем единицу
    }
    for (auto item : NewArray)
    {
        cout << item << " ";
    }
}

int main(void)
{
    setlocale(LC_ALL, "rus");
    int sizeArr1 = 20;

    vector <int> arr1(sizeArr1);
    vector <int> arr2(sizeArr1);

    task2_2and2_3(arr1, arr2,sizeArr1);
    task2_4(sizeArr1);
    task2_5(sizeArr1);

    return 0;

}