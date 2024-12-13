#include <iostream>
#include "grap.h";

using namespace std;

void createVertex(char newVertexId, Vertex *&v){
    v = new Vertex;
    idVertex(v) = newVertexId;
    nextVertex(v) = nullptr;
    firstEdge(v) = nullptr;
}
void initGraph(Graph &G){
    firstVertex(G) = nullptr;
}
void addVertex(Graph &G, char newVertexId){
    Vertex *newVertex = nullptr;
    createVertex(newVertexId, newVertex);

    if(firstVertex(G) != nullptr){
        firstVertex(G) = newVertex;
    } else {
        Vertex *temp = firstVertex(G);
        while(next(temp) != nullptr){
            temp = next(temp);
        }
        next(temp) = newVertex;
    }
}
void buildGraph(Graph &G){
    initGraph(G);

    char input;
    int i = 0;
    while(i != 5){
        cout << "Masukkan ID simpul: ";
        cin >> input;
        addVertex(G, input);
        i++;
    }
}
void printGraph(Graph &G){
    Vertex *v = firstVertex(G);
    while(next(v) != nullptr){
        cout << "Vertex " << idVertex(v) << " ";
        Edge *e = firstEdge(e);
        while(e != nullptr){
            cout << destVertexID(e) << ", " << weight(e);
        }
        cout << endl;
        v = nextVertex(v);
    }
}
