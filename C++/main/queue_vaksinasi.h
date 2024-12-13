#ifndef QUEUE_VAKSINASI_H_INCLUDED
#define QUEUE_VAKSINASI_H_INCLUDED

#include <iostream>
#include <string>

using namespace std;

struct Infotype {
    string nama;
    int usia;
    string pekerjaan;
    string prioritas;
    int nomor_antrean;
    int waktu_daftar;
    bool kondisi_darurat;
};

struct ElemQ;
typedef ElemQ* Address;

struct ElemQ {
    Infotype info;
    Address next;
};

struct Queue {
    Address head;
    Address tail;
};

void createQueue_103022300064(Queue &Q);
bool isEmpty_103022300064(const Queue &Q);
Address createElemQueue_103022300064(const string &nama, int usia, const string &pekerjaan, int nomor_antrean);
void enqueue_103022300064(Queue &Q, Address P);
void dequeue_103022300064(Queue &Q, Address &P);
Address front(const Queue &Q);
Address back(const Queue &Q);
int size(const Queue &Q);
void printInfo_103022300064(const Queue &Q);
void serveQueue_103022300064(Queue &Q);
void reassignQueue_103022300064(Queue &Q);
void checkWaitingTime_103022300064(Queue &Q, int waktu_sekarang);
void emergencyHandle_103022300064(Queue &Q, int nomor_antrean);
void updatePriority_103022300064(Queue &Q);
Address findAndRemove_103022300064(Queue &Q, int nomor_antrean);

#endif // QUEUE_VAKSINASI_H_INCLUDED
