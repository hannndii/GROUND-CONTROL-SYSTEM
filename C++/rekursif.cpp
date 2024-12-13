#include <iostream>
using namespace std;

void printStars(int n) {
    if (n == 0) return; 
    cout << "*";
    printStars(n - 1); 
}

void printPattern(int rows, int current = 1) {
    if (current > rows) return; 
    printStars(current); 
    cout << endl; 
    printPattern(rows, current + 1); 
}

int main()
{
    int a = 5;
    printPattern(a);
    

    return 0;
}


