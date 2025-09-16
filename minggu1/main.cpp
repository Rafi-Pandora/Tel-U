#include <iostream>
#include <stdio.h>
#include <bits/stdc++.h>

using namespace std;

/**
 * @class system
 * @brief ini Kelas dasar yang sya gunakan untuk optimalisasi input/output C++ agar performanya menyamai C.
 *
 * Kelas ini punya fungsi static untuk mempercepat I/O
 * dengan cara mematikan sinkronisasi stdio dan melepas kaitan cin/cout.
 */
class system {
protected:
    /**
     * @brief membantu optimisasi input/output C++.
     *
     * - ios::sync_with_stdio(false) -> fungsi low level C++ untuk mempercepat cin/cout
     * - cin.tie(nullptr) -> fungsi tie dari kelass io mencekal flush otomatis sebelum input
     */
    static void system_cpp_optimize() {
        ios::sync_with_stdio(false);
        cin.tie(nullptr);
    }

public:
    /**
     * @brief mengecek input dari user
     *  - jika input tidak dalam rentang uint8_t(0-255) atau bukan angka maka return false
     * @param input mengecek input uint8_t
     * @return true jika aman dan false jika nilai diluar rentang atau input bukan angka
     */
    static bool input_check(uint8_t input) {
        if (input < 0 || input > 255 || cin.fail()) {
            cout << "input error" << endl;
            return false;
        }
        return true;
    }
};

/**
 * @class soal_1
 * @brief ini Kelas untuk operasi aritmetika (penjumlahan, pengurangan, perkalian, pembagian). sesuai dengan soal di pdf
 *
 * fungsi ini adalah turunan dari class system agar otomatis menggunakan I/O yang lebih cepat.
 */
class soal_1 : public system {
private:
    float angka_1; //Menyimpan angka pertama
    float angka_2; //Menyimpan angka kedua

    //Fungsi private untuk mengambil input dari user
    void input_io() {
        cout << "masukan angka 1" << endl;
        cin >> angka_1;
        cout << "masukan angka 2" << endl;
        cin >> angka_2;
    }

public:
    // Konstruktor (fungsi yang diapnggil saat objek di inisialisai): digunakan untuk mengaktifkan optimisasi dan langsung memanggil input
    soal_1() {
        system::system_cpp_optimize();
        input_io();
    }

    //Menjumlahkan angka_1 + angka_2
    float penjumlahan() {
        return angka_1 + angka_2;
    }

    //Mengurangi angka_1 - angka_2
    float pengurangan() {
        return angka_1 - angka_2;
    }

    //Mengalikan angka_1 * angka_2
    float perkalian() {
        return angka_1 * angka_2;
    }

    /**
     * @brief Membagi angka_1 / angka_2
     * @return Hasil pembagian, atau 0 jika salah satu angka bernilai 0.
     */
    float pembagian() {
        if (angka_1 == 0.0f || angka_2 == 0.0f) {
            cout << "tidak dapat dibagi dengan 0!" << endl;
            return 0;
        }
        return angka_1 / angka_2;
    }
};

/**
 * @class soal_2
 * @brief ini kelas untuk mengkonversi angka rentang (0–100) jadi teks bahasa Indonesia.
 *
 * Misalnya:
 * - 5 -> "lima"
 * - 11 -> "sebelas"
 * - 42 -> "empat puluh dua"
 */
class soal_2 : public system {
private:
    //array string untuk menympan str angka 0–9
    string angka_satuan[10] = {
        "nol", "satu", "dua", "tiga", "empat", "lima",
        "enam", "tujuh", "delapan", "sembilan"
    };

public:
    soal_2() {
        system::system_cpp_optimize();
    }

    /**
     * @brief Mengonversi angka ke teks dalam bahasa Indonesia
     * @param angka Nilai input dari rentang(0–100)
     * @return String representasi angka dalam bahasa Indonesia
     */
    string konversi(uint8_t angka) {
        if (angka > 100) return "angka tidak boleh lebih besar dari seratus";

        if (angka == 100) return "seratus";

        if (angka < 10) return angka_satuan[angka];

        if (angka < 20) {
            if (angka == 10) return "sepuluh";
            if (angka == 11) return "sebelass";
            return angka_satuan[angka % 10] + "belas";
        }

        if (angka < 100) {
            uint8_t puluh = angka / 10;
            uint8_t satuan = angka % 10;
            string hasil = angka_satuan[puluh] + " puluh";
            if (satuan != 0) hasil += " " + angka_satuan[satuan];
            return hasil;
        }

        return "";
    }
};

/**
 * @class soal_3
 * @brief ini kelas untuk mencetak pola angka dari input user.
 *
 * Contoh (n = 5):
 * 5 4 3 2 1 * 1 2 3 4 5
 *   4 3 2 1 * 1 2 3 4
 *     3 2 1 * 1 2 3
 *       2 1 * 1 2
 *         1 * 1
 */
class soal_3 : public system {
private:
    uint8_t input; //Menyimpan input dari user

public:
    soal_3() {
        system::system_cpp_optimize();
    }

    /**
     * @brief Cetak pola angka dengan bintang ditengahnya
     * @param n Ukuran pola bedasarakan input user
     */
    void pattern(int n) {
        for (int i = n; i >= 1; --i) {
            for (int s = 0; s < n - i; ++s)
                cout << "  ";

            for (int j = i; j >= 1; --j)
                cout << j << " ";

            cout << "* ";

            for (int j = 1; j <= i; ++j)
                cout << j << " ";

            cout << endl;
        }
    }
};

int main() {
    soal_1 S1;
    soal_2 S2; //deklarasi objek
    soal_3 S3;

    uint8_t angka; //deklarasi temp var angka dari rentang 0 - 255

    printf("Penjumlahan: %f", S1.penjumlahan());
    printf("Pengurangan: %f", S1.pengurangan());
    printf("Perkalian: %f", S1.perkalian());
    printf("Pembagian: %s", S1.pembagian());

    printf("masukan angka range (0-100)");
    
    cin >> angka;
    if (!system::input_check(angka)) return 0;
    printf("hasil konversi: %s", S2.konversi(angka));

    printf("input");
    cin >> angka;
    if (!system::input_check(angka)) return 0;
    S3.pattern(angka);
}