#include <iostream>
#include "Tree.h"

using namespace std;

adrNode newNode_103022300064(infotype x){
    adrNode temp = new Node;
    temp->info = x;
    temp->left = nullptr;
    temp->right = nullptr;
    return temp;
}
adrNode findNode_103022300064(adrNode root, infotype x){
    if(root == nullptr || root->info == x){
        return root;
    }
    if (x < root->info){
        return findNode_103022300064(root->left, x);
    } else {
        return findNode_103022300064(root->right, x);
    }
}
void insertNode_103022300064(adrNode &root, adrNode p){
    if(root == nullptr){
        root = p;
    } else if(root->info < p->info){
        return insertNode_103022300064(root->right, p);
    } else {
        return insertNode_103022300064(root->left, p);
    }
}
void printPreOrder_103022300064(adrNode root){
    if(root != nullptr){
        cout << root->info << " ";
        printPreOrder_103022300064(root->left);
        printPreOrder_103022300064(root->right);
    }
}
void printDescendant_103022300064(adrNode root, infotype x) {
    adrNode node = findNode_103022300064(root, x);
    if (node == nullptr) {
        cout << "Node dengan info " << x << " tidak ditemukan.\n";
        return;
    }

    if (node->left != nullptr) {
        printPreOrder_103022300064(node->left);
    }
    if (node->right != nullptr) {
        printPreOrder_103022300064(node->right);
    }
}
int sumNode_103022300064(adrNode root){
    if(root == nullptr){
        return 0;
    } else {
        return root->info + sumNode_103022300064(root->left)+sumNode_103022300064(root->right);
    }
}
int countLeaves_103022300064(adrNode root){
    if(root == nullptr){
        return 0;
    }
    if(root->left == nullptr && root->right == nullptr){
        return 1;
    }
    return countLeaves_103022300064(root->left) + countLeaves_103022300064(root->right);
}
int heightTree_103022300064(adrNode root){
    if (root == nullptr) {
        return -1;
    }
    int leftHeight = heightTree_103022300064(root->left);
    int rightHeight = heightTree_103022300064(root->right);
    return 1 + max(leftHeight, rightHeight);
}
