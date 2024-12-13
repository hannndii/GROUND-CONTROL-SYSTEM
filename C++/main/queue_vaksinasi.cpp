#include "queue_vaksinasi.h"
#include <iostream>

using namespace std;

void createQueue_103022300064(Queue &Q) {
    Q.head = nullptr;
    Q.tail = nullptr;
}

bool isEmpty(const Queue &Q) {
    return Q.head == nullptr;
}

Address createElemQueue_103022300064(const string &nama, int usia, const string &pekerjaan, int nomor_antrean) {
    Address P = new ElemQ;
    P->info.nama = nama;
    P->info.usia = usia;
    P->info.pekerjaan = pekerjaan;
    P->info.prioritas = (usia >= 60 || pekerjaan == "tenaga kesehatan") ? "Ya" : "Tidak";
    P->info.nomor_antrean = nomor_antrean;
    P->info.waktu_daftar = 0;
    P->info.kondisi_darurat = false;
    P->next = nullptr;
    return P;
}

void enqueue_103022300064(Queue &Q, Address P) {
    if (isEmpty_103022300064) {
        Q.head = P;
        Q.tail = P;
    } else if (P->info.prioritas == "Ya") {
        if (Q.head->info.prioritas != "Ya") {
            P->next = Q.head;
            Q.head = P;
        } else {
            Address temp = Q.head;
            Address prev = nullptr;
            while (temp != nullptr && temp->info.prioritas == "Ya" && temp->info.nomor_antrean < P->info.nomor_antrean) {
                prev = temp;
                temp = temp->next;
            }
            if (prev == nullptr) {
                P->next = Q.head;
                Q.head = P;
            } else {
                P->next = temp;
                prev->next = P;
            }
            if (P->next == nullptr) {
                Q.tail = P;
            }
        }
    } else {
        Address temp = Q.head;
        Address prev = nullptr;
        while (temp != nullptr && temp->info.nomor_antrean < P->info.nomor_antrean) {
            prev = temp;
            temp = temp->next;
        }
        if (prev == nullptr) {
            P->next = Q.head;
            Q.head = P;
        } else {
            P->next = temp;
            prev->next = P;
        }
        if (P->next == nullptr) {
            Q.tail = P;
        }
    }
}

void dequeue_103022300064(Queue &Q, Address &P) {
    if (isEmpty_103022300064) {
        P = nullptr;
        cout << "Semua warga telah dilayani." << endl;
    } else {
        P = Q.head;
        Q.head = Q.head->next;
        if (Q.head == nullptr) {
            Q.tail = nullptr;
        }
        P->next = nullptr;
    }
}

Address front(const Queue &Q) {
    return Q.head;
}

Address back(const Queue &Q) {
    return Q.tail;
}

int size(const Queue &Q) {
    int count = 0;
    Address temp = Q.head;
    while (temp != nullptr) {
        count++;
        temp = temp->next;
    }
    return count;
}

void printInfo_103022300064(const Queue &Q) {
    if (isEmpty_103022300064) {
        cout << "Antrean kosong!" << endl;
        return;
    }

    Address temp = Q.head;
    cout << "Daftar antrean:" << endl;

    while (temp != nullptr) {
        cout << "Nama         : " << temp->info.nama << endl;
        cout << "Usia         : " << temp->info.usia << endl;
        cout << "Pekerjaan    : " << temp->info.pekerjaan << endl;
        cout << "Prioritas    : " << temp->info.prioritas << endl;
        cout << "Nomor Antrean: " << temp->info.nomor_antrean << endl;
        cout << "=========================" << endl;
        temp = temp->next;
    }
}

void serveQueue_103022300064(Queue &Q) {
    Address P = nullptr;
    int serve_count = 0;

    while (!isEmpty_103022300064 && serve_count < 100) {
        dequeue_103022300064(Q, P);
        cout << "Melayani warga:" << endl;
        cout << "Nama         : " << P->info.nama << endl;
        cout << "Usia         : " << P->info.usia << endl;
        cout << "Pekerjaan    : " << P->info.pekerjaan << endl;
        cout << "Prioritas    : " << P->info.prioritas << endl;
        cout << "Nomor Antrean: " << P->info.nomor_antrean << endl;
        cout << "Vaksinasi berhasil" << endl;
        cout << "=========================" << endl;
        serve_count++;
    }

    if (serve_count == 100) {
        cout << "Kapasitas harian penuh" << endl;
    }

    if (!isEmpty_103022300064) {
        cout << "Warga yang belum terlayani diminta kembali besok." << endl;
    }
}

void reassignQueue_103022300064(Queue &Q) {
    Queue tempQPriority;
    Queue tempQNormal;
    createQueue_103022300064(tempQPriority);
    createQueue_103022300064(tempQNormal);

    Address P = nullptr;
    while (!isEmpty_103022300064) {
        dequeue_103022300064(Q, P);
        if (P->info.prioritas == "Ya") {
            enqueue_103022300064(tempQPriority, P);
        } else {
            enqueue_103022300064(tempQNormal, P);
        }
    }

    while (!isEmpty(tempQPriority)) {
        dequeue_103022300064(tempQPriority, P);
        enqueue_103022300064(Q, P);
    }

    while (!isEmpty(tempQNormal)) {
        dequeue_103022300064(tempQNormal, P);
        enqueue_103022300064(Q, P);
    }
}

void checkWaitingTime_103022300064(Queue &Q, int waktu_sekarang) {
    Queue tempQ;
    createQueue_103022300064(tempQ);

    Address P;
    while (!isEmpty_103022300064) {
        dequeue_103022300064(Q, P);
        if ((waktu_sekarang - P->info.waktu_daftar) > 120 && P->info.prioritas != "Ya") {
            P->info.prioritas = "Ya";
        }
        enqueue_103022300064(tempQ, P);
    }

    while (!isEmpty(tempQ)) {
        dequeue_103022300064(tempQ, P);
        enqueue_103022300064(Q, P);
    }
}

void emergencyHandle_103022300064(Queue &Q, int nomor_antrean) {
    Address emergencyElem = findAndRemove_103022300064(Q, nomor_antrean);
    if (emergencyElem == nullptr) {
        cout << "Warga dengan nomor antrean " << nomor_antrean << " tidak ditemukan." << endl;
        return;
    }

    emergencyElem->info.kondisi_darurat = true;
    emergencyElem->info.prioritas = "Ya";
    emergencyElem->next = Q.head;
    Q.head = emergencyElem;
    cout << "Warga dengan nomor antrean " << nomor_antrean << " diprioritaskan karena kondisi darurat." << endl;
}

void updatePriority_103022300064(Queue &Q) {
    checkWaitingTime_103022300064(Q, 120);

    Queue tempQ;
    createQueue_103022300064(tempQ);

    Address P;
    while (!isEmpty_103022300064) {
        dequeue_103022300064(Q, P);
        if (P->info.kondisi_darurat || P->info.prioritas == "Ya") {
            enqueue_103022300064(tempQ, P);
        } else {
            enqueue_103022300064(Q, P);
        }
    }

    while (!isEmpty(tempQ)) {
        dequeue_103022300064(tempQ, P);
        enqueue_103022300064(Q, P);
    }
}

Address findAndRemove_103022300064(Queue &Q, int nomor_antrean) {
    if (isEmpty_103022300064) return nullptr;

    Address P = Q.head;
    Address prev = nullptr;

    while (P != nullptr && P->info.nomor_antrean != nomor_antrean) {
        prev = P;
        P = P->next;
    }

    if (P == nullptr) return nullptr;

    if (prev == nullptr) {
        Q.head = P->next;
    } else {
        prev->next = P->next;
    }

    if (P == Q.tail) {
        Q.tail = prev;
    }

    P->next = nullptr;
    return P;
}
