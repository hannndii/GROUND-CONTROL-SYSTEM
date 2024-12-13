#include "stack.h"
#include <iostream>

using namespace std;

int main() {
    stack S;
    createStack_103022300064(S);

    cout << "======MASUKKAN DATA (MAKSIMAL 15)======" << endl;
    for (int i = 0; i < 10; i++) {
        char element;
        cout << "Data ke-" << i+1 << ":";
        cin >> element;

        if (!isFull_103022300064(S)) {
            push_103022300064(S, element);
        } else {
            cout << "======DATA TERISI PENUH=====" << endl;
            break;
        }
    }

    cout << "=======DATA TERBARU======" << endl;
    printInfo_103022300064(S);
    for (int i = 0; i < 10; i++) {
        if(i > 5){
            cout << info(S)[i] << " ";
        }
    }
    cout << endl;

    return 0;
}
