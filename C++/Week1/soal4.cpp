#include <iostream>
using namespace std;

int main() {
    int input, sum = 0;
    cout << "Input value: ";
    cin >> input;

    // Looping hanya sampai setengah dari input
    for(int i = 1; i <= input / 2; i++) {
        if(input % i == 0) {
            sum += i;  // Menambahkan faktor ke sum
        }
    }

    // Mengecek apakah jumlah faktor sama dengan input
    if(sum == input) {
        cout << input << " adalah Bilangan Sempurna" << endl;
    } else {
        cout << input << " bukan Bilangan Sempurna" << endl;
    }

    return 0;
}
