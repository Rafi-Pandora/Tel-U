#include <iostream>
#include <bits/stdc++.h>

using namespace std;

class soal_1 {
    private:
        float angka_1;
        float angka_2;

        void input_io() {
            cout << "masukan angka 1" << endl;
            cin >> angka_1;
            cout << "masukan angka 2" << endl;
            cin >> angka_2;
        }
    
    public:
    soal_1() {
        ios::sync_with_stdio(false);
        cin.tie(nullptr);
        input_io();
    }

    float penjumlahan() {
        return angka_1 + angka_2;
    }

    float pengurangan() {
        return angka_1 - angka_2;
    }

    float perkalian() {
        return angka_1 * angka_2;
    }

    float pembagian() {
        if (angka_1 == 0.0f || angka_2 == 0.0f)
        {
            cout << "tidak dapat dibagi dengan 0!" << endl;
            return 0;
        }
        return angka_1 / angka_2;
    }
};

class soal_2 {
    private:
        string angka_satuan[] = {
            "nol", "satu", "dua", "tiga", "empat", "lima",
            "enam", "tujuh", "delapan", "sembilan"
        };
};

int main() {

    return 0;
}