#include "pelajaran.h"

class Kelas {
private:
    uint jumlah_mahasiswa;
    char** nama_mahasiswa;

public:
    Kelas(uint jumlah) {
        jumlah_mahasiswa = jumlah;
        nama_mahasiswa = new char*[jumlah_mahasiswa];

        for (int i = 0; i < static_cast<int>(jumlah_mahasiswa); i++)
        {
            *(nama_mahasiswa + i) = new char[50];
        }
    }

    ~Kelas() {
        for (int i = 0; i < static_cast<int>(jumlah_mahasiswa); i++)
        {
            delete[] *(nama_mahasiswa + i);
        }
        delete[] nama_mahasiswa;
    }

    void input_nama_mahasiswa() {
        cout << "Masukan Nama Mahasiswa" << endl;
        cin.ignore();
        for (int i = 0; i < static_cast<int>(jumlah_mahasiswa); i++)
        {
            cout << "masukan nama mahasiswa ke - " << i + 1 << endl;
            cin.getline(*(nama_mahasiswa + i), 50);
        }
    }

    void tampil_mahasiswa() {
        for (int i = 0; i < static_cast<int>(jumlah_mahasiswa); i++)
        {
            cout << i + 1 << ". " << *(nama_mahasiswa + i) << endl;
        }
    }
};

int main(int argc, char const *argv[])
{
    Pelajaran pelajaran;
    pelajaran_t struct_pelajaran;

    struct_pelajaran = pelajaran.create_pelajaran("Struktur Data","STD");
    pelajaran.tampil_pelajaran(struct_pelajaran);
    return 0;
}
