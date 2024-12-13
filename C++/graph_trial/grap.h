#ifndef GRAP_H_INCLUDED
#define GRAP_H_INCLUDED
#define firstVertex(G) G.firstVertex
#define nextVertex(v) v->nextVertex
#define nextEdge(G) G->nextEdge
#define firstEdge(G) G.firstEdge
#define idVertex(v) v->idVertex
#define destVertexID(v) v->destVertexID
#include <iostream>

typedef struct Vertex *adrVertex;
typedef struct Edge *adrEdge;

struct Edge{
    char destVertexID;
    int weight;
    Edge *nextEdge;
};

struct Vertex{
    char idVertex;
    Vertex *nextVertex;
    Edge *firstEdge;
};

struct Graph{
    Vertex* firstVertex;
};

void createVertex(char newVertexId, Vertex *&v);
void initGraph(Graph &G);
void addVertex(Graph &G, char newVertexId);
void buildGraph(Graph &G);
void printGraph(const Graph &G);


#endif // GRAP_H_INCLUDED
