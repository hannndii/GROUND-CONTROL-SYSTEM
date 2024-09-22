// MASIH ERROR
#include <iostream>

using namespace std;

int main() {
    int input;
    bool isPrime = true;

    cout << "Input the number: ";
    cin >> input;
   
    if (input <= 1) {
        isPrime = false; // Bilangan <= 1 bukan bilangan prima
    } else {
        for (int i = 2; i < input; i++) {
            if (input % i == 0) {
                isPrime = false; // Jika ada pembagi selain 1 dan input, bukan prima
                break;
            }
        }
    }

    if (!isPrime) {
        cout << "Not Prime Number!";
    } else {
        cout << "Is Prime";
    }

    return 0;
}
