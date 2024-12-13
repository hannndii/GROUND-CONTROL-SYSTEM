x#ifndef GRAPH_H_INCLUDED
#define GRAPH_H_INCLUDED

#ifndef GRAPH_H
#define GRAPH_H
#include <iostream>
using namespace std;


typedef struct Vertex *adrVertex;
typedef struct Edge *adrEdge;

struct Edge {
    char destVertexID;
    int weight;
    Edge* nextEdge;
};

struct Vertex {
    char idVertex;
    Vertex* nextVertex;
    Edge* firstEdge;
};

struct Graph {
    Vertex* firstVertex;
};

void createVertex(char newVertexID, Vertex*& v);
void initGraph(Graph& G);
void addVertex(Graph& G, char newVertexID);
void buildGraph(Graph& G);
void printGraph(const Graph& G);

#endif




#endif // GRAPH_H_INCLUDED
