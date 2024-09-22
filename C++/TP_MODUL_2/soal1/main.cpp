#include <iostream>

using namespace std;

int kendaraan(int kapasitas_kendaraan,int jumlah_kendaraan){
    int jumlah;
    jumlah = jumlah_kendaraan / kapasitas_kendaraan;
    if(jumlah_kendaraan % kapasitas_kendaraan > 0){
        jumlah++;
    }
    return jumlah;
}

int main()
{
    int kap_kendaraan, jum_kendaraan, banyak_kendaraan;

    cout << "Masukkkan kapasitas kendaraan: ";
    cin >> kap_kendaraan;
    cout << "Masukkkan jumlah kendaraan: ";
    cin >> jum_kendaraan;
    banyak_kendaraan = kendaraan(kap_kendaraan, jum_kendaraan);
    cout << "Banyak kendaraan yang disewa: " << banyak_kendaraan << endl;
    return 0;
}
