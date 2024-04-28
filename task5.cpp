#include <iostream>
#include <vector>
#include <locale.h>

using namespace std;

int main()
{
	setlocale(LC_ALL, "Rus");
	int A = 0, B = 0, C = 0, X0 = 0, Xi = 0;;
	vector<int> arr1;
	cout << "Введите диапозон значений С:\nC = ";
	cin >> C;
	cout << "Введите множитель A(0<=A<=C):\nA = ";
	cin >> A;
	cout << "Введите инкрементируещее значений B(0<=B<=C):\nB = ";
	cin >> B;
	cout << "Введите начальное значений X0(0<=X0<=C):\nX0 = ";
	cin >> X0;
	arr1.emplace(arr1.begin(), X0);
	for (int i = 1; i < C; i++)
	{
		Xi = (A * arr1[i - 1] + B) % C;
		arr1.emplace(arr1.begin() + i, Xi);
	}
	for (auto i : arr1)
	{
		cout << i << " ";
	}
	for (size_t i = 1; i<arr1.size(); i++)
	{
		if (arr1[0] == arr1[i])
		{
			cout << "\nИндекс начала повторяющейся последовательности = " << i+1 << endl;
			break;
		}
	}

	return 0;
}