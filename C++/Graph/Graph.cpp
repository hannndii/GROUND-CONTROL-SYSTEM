#include "graph.h"

void createVertex(char newVertexID, Vertex*& v) {
    v = new Vertex;
    v->idVertex = newVertexID;
    v->nextVertex = nullptr;
    v->firstEdge = nullptr;
}

void initGraph(Graph& G) {
    G.firstVertex = nullptr;
}

void addVertex(Graph& G, char newVertexID) {
    Vertex* newVertex = nullptr;
    createVertex(newVertexID, newVertex);

    if (G.firstVertex == nullptr) {
        G.firstVertex = newVertex;
    } else {
        Vertex* temp = G.firstVertex;
        while (temp->nextVertex != nullptr) {
            temp = temp->nextVertex;
        }
        temp->nextVertex = newVertex;
    }
}

void buildGraph(Graph& G) {
    initGraph(G);

    char input;
    cout << "Masukkan ID simpul (A-Z): ";
    while (cin >> input) {
        if (!isalpha(input) || !isupper(input)) break;
        addVertex(G, input);
        cout << "Masukkan ID simpul (A-Z): ";
    }
}

void printGraph(const Graph& G) {
    Vertex* v = G.firstVertex;
    while (v != nullptr) {
        cout << "Vertex " << v->idVertex << " -> ";
        Edge* e = v->firstEdge;
        while (e != nullptr) {
            cout << "(" << e->destVertexID << ", " << e->weight << ") ";
            e = e->nextEdge;
        }
        cout << endl;
        v = v->nextVertex;
    }
}
