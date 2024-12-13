#ifndef TREE_H_INCLUDED
#define TREE_H_INCLUDED
#include <iostream>

typedef struct Node *adrNode;
typedef int infotype;

struct Node{
    infotype info;
    adrNode left;
    adrNode right;
};

adrNode newNode_103022300064(infotype x);
adrNode findNode_103022300064(adrNode root, infotype x);
void insertNode_103022300064(adrNode &root, adrNode p);
void printPreOrder_103022300064(adrNode root);
void printDescendant_103022300064(adrNode root, infotype x);
int sumNode_103022300064(adrNode root);
int countLeaves_103022300064(adrNode root);
int heightTree_103022300064(adrNode root);


#endif // TREE_H_INCLUDED
