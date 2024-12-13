#include "queue_vaksinasi.h"
#include <iostream>

int main() {
    Queue Q;
    createQueue_103022300064(Q);

    // Menambahkan beberapa elemen ke dalam antrean
    Address P1 = createElemQueue_103022300064("John Doe", 65, "lansia", 1);
    Address P2 = createElemQueue_103022300064("Alice", 30, "tenaga kesehatan", 2);
    Address P3 = createElemQueue_103022300064("Bob", 25, "pekerja", 3);
    Address P4 = createElemQueue_103022300064("Charlie", 70, "pensiunan", 4);
    Address P5 = createElemQueue_103022300064("David", 28, "pekerja", 5);

    enqueue_103022300064(Q, P1);
    enqueue_103022300064(Q, P2);
    enqueue_103022300064(Q, P3);
    enqueue_103022300064(Q, P4);
    enqueue_103022300064(Q, P5);

    // Menampilkan isi antrean awal
    std::cout << "Isi antrean awal:" << std::endl;
    printInfo_103022300064(Q);

    // Melayani antrean
    std::cout << "\nMelakukan pelayanan pada antrean:" << std::endl;
    serveQueue_103022300064(Q);

    // Menampilkan isi antrean setelah pelayanan
    std::cout << "\nIsi antrean setelah pelayanan:" << std::endl;
    printInfo_103022300064(Q);

    // Menambahkan elemen baru untuk simulasi pengaturan ulang prioritas
    ElemQ* P6 = createElemQueue_103022300064("Edward", 22, "pekerja", 6);
    enqueue_103022300064(Q, P6);

    // Simulasi reassignQueue untuk mengatur ulang antrean berdasarkan prioritas
    std::cout << "\nMengatur ulang antrean berdasarkan prioritas:" << std::endl;
    reassignQueue_103022300064(Q);
    printInfo_103022300064(Q);

    // Simulasi kondisi warga yang menunggu lebih dari 2 jam
    std::cout << "\nMemeriksa waktu tunggu dan mengubah prioritas jika lebih dari 2 jam:" << std::endl;
    checkWaitingTime_103022300064(Q, 130); // Misalnya waktu sekarang 130 menit dari waktu daftar pertama
    printInfo_103022300064(Q);

    // Menangani kondisi darurat untuk seorang warga
    std::cout << "\nMenangani kondisi darurat untuk warga dengan nomor antrean 5:" << std::endl;
    emergencyHandle_103022300064(Q, 5);
    printInfo_103022300064(Q);

    // Update prioritas warga setiap jam
    std::cout << "\nMengupdate prioritas antrean setiap jam:" << std::endl;
    updatePriority_103022300064(Q);
    printInfo_103022300064(Q);

    // Mencari dan menghapus warga dengan nomor antrean tertentu
    std::cout << "\nMenghapus warga dengan nomor antrean 3:" << std::endl;
    Address removedElem = findAndRemove_103022300064(Q, 3);

    if (removedElem) {
        std::cout << "Warga yang dihapus: " << removedElem->info.nama << std::endl;
        delete removedElem;
    } else {
        std::cout << "Warga dengan nomor antrean 3 tidak ditemukan." << std::endl;
    }
    printInfo_103022300064(Q);

    // Mengecek ukuran antrean saat ini
    std::cout << "\nUkuran antrean saat ini: " << size(Q) << std::endl;

    return 0;
}
