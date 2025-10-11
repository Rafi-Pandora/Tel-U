#include "pelajaran.h"

/**
 * @class Kelas
 * @brief Nampung dan ngatur nama-nama mahasiswa dalam satu kelas.
 */
class Kelas {
private:
    uint jumlah_mahasiswa;      ///< Jumlah mahasiswa yang ada di kelas.
    char** nama_mahasiswa;      ///< Pointer ke array yang nyimpen nama-nama mahasiswa.

public:
    /**
     * @brief Bikin objek kelas baru dan siapin tempat buat nama-nama mahasiswa.
     * @param jumlah Jumlah mahasiswa yang mau dimasukin.
     */
    Kelas(uint jumlah) {
        jumlah_mahasiswa = jumlah;
        nama_mahasiswa = new char*[jumlah_mahasiswa];

        for (int i = 0; i < static_cast<int>(jumlah_mahasiswa); i++) {
            *(nama_mahasiswa + i) = new char[50];
        }
    }

    /**
     * @brief Bersihin memori pas objeknya udah nggak dipakai.
     */
    ~Kelas() {
        for (int i = 0; i < static_cast<int>(jumlah_mahasiswa); i++) {
            delete[] *(nama_mahasiswa + i);
        }
        delete[] nama_mahasiswa;
    }

    /**
     * @brief Minta user buat masukin nama-nama mahasiswa.
     * 
     * Setiap nama bisa sampai 50 karakter, termasuk spasi.
     */
    void input_nama_mahasiswa() {
        cout << "Masukan Nama Mahasiswa" << endl;
        cin.ignore();
        for (int i = 0; i < static_cast<int>(jumlah_mahasiswa); i++) {
            cout << "Masukan nama mahasiswa ke - " << i + 1 << endl;
            cin.getline(*(nama_mahasiswa + i), 50);
        }
    }

    /**
     * @brief Nampilin semua nama mahasiswa yang udah dimasukin.
     */
    void tampil_mahasiswa() {
        for (int i = 0; i < static_cast<int>(jumlah_mahasiswa); i++) {
            cout << i + 1 << ". " << *(nama_mahasiswa + i) << endl;
        }
    }
};

/**
 * @class Nilai
 * @brief Ngatur nilai-nilai mahasiswa dan ngitung rata-ratanya.
 */
class Nilai {
private:
    uint jumlah_mahasiswa;  ///< Jumlah mahasiswa yang dikasih nilai.
    float* nilai_mahasiswa; ///< Array dinamis buat nyimpen nilai mereka.
    float rata_rata;        ///< Hasil rata-rata nilai mahasiswa.

public:
    /**
     * @brief Bikin objek Nilai dan siapin array buat nyimpen nilai.
     * @param jumlah Jumlah mahasiswa yang bakal diinput nilainya.
     */
    Nilai(uint jumlah) {
        jumlah_mahasiswa = jumlah;
        nilai_mahasiswa = new float[jumlah_mahasiswa];
        rata_rata = 0.0f;
    }

    /**
     * @brief Bersihin memori pas objeknya dihapus.
     */
    ~Nilai() {
        delete[] nilai_mahasiswa;
    }

    /**
     * @brief Minta input nilai dari user buat tiap mahasiswa.
     */
    void input_nilai_mahasiswa() {
        for (int i = 0; i < static_cast<int>(jumlah_mahasiswa); i++) {
            cout << "Masukan nilai mahasiswa ke - " << i + 1 << endl;
            cin >> *(nilai_mahasiswa + i);
        }
    }

    /**
     * @brief Hitung rata-rata dari semua nilai mahasiswa.
     */
    void hitung_rata_rata() {
        float total = 0.0f;
        float* ptr = nilai_mahasiswa;

        for (int i = 0; i < static_cast<int>(jumlah_mahasiswa); i++) {
            total += *(ptr + i);
        }

        rata_rata = total / static_cast<float>(jumlah_mahasiswa);
    }

    /**
     * @brief Tampilkan hasil rata-rata nilai mahasiswa ke layar.
     */
    void tampil_rata_rata() const {
        cout << "Rata-rata nilai: " << rata_rata << endl;
    }
};

/**
 * @brief Fungsi utama dari program ini.
 * 
 * Program bakal nampilin info pelajaran, minta input nilai & nama mahasiswa,
 * hitung rata-rata nilai, terus nampilin semuanya.
 * 
 * @param argc Jumlah argumen dari command line.
 * @param argv Isi argumen command line.
 * @return 0 kalau program jalan dengan sukses.
 */
int main(int argc, char const* argv[]) {
    Pelajaran pelajaran;
    pelajaran_t struct_pelajaran;

    // Bikin pelajaran baru dan tampilkan
    struct_pelajaran = pelajaran.create_pelajaran("Struktur Data", "STD");
    pelajaran.tampil_pelajaran(struct_pelajaran);

    // Input nilai dan hitung rata-rata
    Nilai nilai(5);
    nilai.input_nilai_mahasiswa();
    nilai.hitung_rata_rata();
    nilai.tampil_rata_rata();

    // Input dan tampilkan nama-nama mahasiswa
    Kelas kelas(5);
    kelas.input_nama_mahasiswa();
    kelas.tampil_mahasiswa();

    return 0;
}
