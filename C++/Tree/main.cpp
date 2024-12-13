#include <iostream>
#include "Tree.h"

using namespace std;

int main()
{
     int x[9] = {5, 3, 9, 10, 4, 7, 1, 8, 6};
    adrNode root = nullptr;

    // 1. Tambahkan setiap elemen array x ke dalam BST secara berurutan
    for (int i = 0; i < 9; ++i) {
        adrNode newNodePtr = newNode_103022300064(x[i]);
        insertNode_103022300064(root, newNodePtr);
    }

    // 2. Tampilkan node dari BST secara Pre-Order
    cout << "\nPre Order\t: ";
    printPreOrder_103022300064(root);

    // 3. Tampilkan keturunan dari node 9
    cout << "\n\nDescendent of Node 9\t: ";
    printDescendant_103022300064(root, 9);

    // 4. Tampilkan total info semua node pada BST
    cout << "\n\nSum of BST Info\t: ";
    cout << sumNode_103022300064(root) << endl;

    // 5. Tampilkan banyaknya daun dari BST
    cout << "\nNumber of Leaves\t: ";
    cout << countLeaves_103022300064(root) << endl;

    // 6. Tampilkan Tinggi dari Tree
    cout << "\nHeight of Tree\t\t: ";
    cout << heightTree_103022300064(root) << endl;

    return 0;
}
