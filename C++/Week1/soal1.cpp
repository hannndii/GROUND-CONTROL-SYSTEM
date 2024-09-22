#include <iostream>

using namespace std;

int main(){
    int input, positive = 0, negative = 0;
    cout << "Masukkan Nilai:";
    cin >> input;
    while(input != 0){
        if(input > 0){
            positive++;
        } else {
            negative++;
        }
        cin >> input;
    }
    string status;
    if(positive > negative){
        status = "P";
    } else if (positive == negative){
        status = "Imbang";
    } else{
        status = "N";
    }

    cout << positive << " " << negative << " " << status << endl;
    return 0;
}