#include <cstdint>
#include <iostream>

using namespace std;

class nilai_mahasiswa {
  private:
    uint8_t mahasiswa;
    uint8_t *nilai;

  public:
    nilai_mahasiswa(int jumlah_mahasiswa) {
      if (jumlah_mahasiswa < 0 || jumlah_mahasiswa > 100) {
        cout << "input error" << endl;
        return;
      }
      
      mahasiswa = (uint8_t)jumlah_mahasiswa;
      nilai = new uint8_t[(uint8_t)jumlah_mahasiswa];
    }

    ~nilai_mahasiswa() {
      delete[] nilai;
    }

    void input_nilai() {
      cout << "Masukan data mahasiswa" << endl;

      for (int i = 0; i < (int)mahasiswa; i++) {
        cout << "masukan nilai untuk mahasiswa - " << (i + 1) << endl;
        cin >> nilai[i];

        if (cin.fail()) return;
      }
    }

    uint8_t get_maks_nilai() {
      uint8_t temp = nilai[0];      

      for (int i = 0; i < (int)mahasiswa; i++) {
        if (temp < nilai[i]) {
          temp = nilai[i];
        }
      }

      return temp;
    }

    uint8_t get_min_nilai() {
      uint8_t temp = nilai[0];

      for (int i = 0; i < (int)mahasiswa; i++ ) {
        if (temp > nilai[i]) {
          temp = nilai[i];
        }
      }

      return temp;
    }

    float get_sum_nilai() {
      float temp;
      for (int i = 0; i < (int)mahasiswa; i++ ) {
        temp += nilai[i];
      }

      return (temp / (int)mahasiswa);
    }
};

class Produk {
  private:
    static const uint8_t jumlah_produk = 5;
    uint8_t produk[(int)jumlah_produk];
  
  public:
    Produk() {
      for (uint8_t i = 0; i < jumlah_produk; i++ ) {
        cout << "masukan produk ke - " << i + 1 << endl;
        cin >> produk[(int)i];
      }
    }

  int total_penjualan() {
    int temp;
    for (uint8_t i = 0; i < jumlah_produk; i++ ) {
      temp += (int)produk[i];
    }
    return temp;
  }

  void maksimum_penjualan() {
    uint8_t temp = produk[0];
    uint8_t j;
    for (uint8_t i = 0; i < jumlah_produk; i++ ) {
      if (temp < produk[i]) {
        temp = produk[i];
        j = i + 1;
      }
    }
    cout << "penjualan tertinggi produk ke - " << j << "(" << temp << ")" << endl;
  }

  void minimum_penjualan() {
    uint8_t temp = produk[0];
    uint8_t j;
    for (uint8_t i = 0; i < jumlah_produk; i++ ) {
      if (temp > produk[i]) {
        temp = produk[i];
        j = i + 1;
      }
    }
    cout << "penjualan terendah produk ke - " << j << "(" << temp << ")" << endl;
  }
};

int main () {
  
  return 0;
} 
