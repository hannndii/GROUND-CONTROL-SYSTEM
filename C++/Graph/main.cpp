#include "graph.h"

int main() {
    Graph G;

    buildGraph(G);

    cout << "Graf yang telah dibangun: " << endl;
    printGraph(G);

    return 0;
}
