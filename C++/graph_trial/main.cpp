#include <iostream>

using namespace std;

int main()
{
    Graph G;

    buildGraph(G);

    cout << "Graf:" << endl;
    printGraph(G);
    return 0;
}
