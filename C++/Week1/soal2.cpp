#include <iostream>

using namespace std;

int main(){
    int input, maxEven = 0, sumEven;
    cout << "Masukkan Nilai: ";
    cin >> input;
    while(input != 0){
        if(input % 2 == 0) {
            sumEven++;
            if(input > maxEven){
                maxEven = input ;
            }
        }
        cin >> input;
    } 
    cout << "Sum of Even Numbers: " << sumEven << "\n";
    cout << "The Highest Even Number is " << maxEven;
    return 0;
}