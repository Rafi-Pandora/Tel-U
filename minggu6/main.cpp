#include <iostream>
#include <string>
using namespace std;

/**
 * @brief Struktur data untuk nyimpen info produk di toko.
 * 
 * Tiap produk punya kode, nama, dan harga.
 * Terhubung dua arah lewat pointer prev dan next (Doubly Linked List).
 */
struct Produk {
    string kodeProduk;  ///< kode unik buat produk
    string namaProduk;  ///< nama dari produk
    int harga;          ///< harga produk
    Produk* prev;       ///< pointer ke node sebelumnya
    Produk* next;       ///< pointer ke node selanjutnya
};

/**
 * @brief Kelas buat ngatur daftar produk (pakai Doubly Linked List).
 * 
 * Bisa nambah produk di awal, akhir, atau setelah produk tertentu,
 * dan juga bisa nampilin semua produk.
 */
class ProdukList {
private:
    Produk* head;  ///< pointer ke produk pertama
    Produk* tail;  ///< pointer ke produk terakhir

public:
    /// Konstruktor: mulaiin list kosong
    ProdukList() {
        head = tail = nullptr;
    }

    /// Cek apakah list kosong
    bool isEmpty() {
        return head == nullptr;
    }

    /**
     * @brief Tambah produk di bagian paling depan list.
     * @param kode kode produk
     * @param nama nama produk
     * @param harga harga produk
     */
    void insertFirst(string kode, string nama, int harga) {
        Produk* baru = new Produk{kode, nama, harga, nullptr, nullptr};
        if (isEmpty()) {
            head = tail = baru;
        } else {
            baru->next = head;
            head->prev = baru;
            head = baru;
        }
    }

    /**
     * @brief Tambah produk di bagian paling akhir list.
     */
    void insertLast(string kode, string nama, int harga) {
        Produk* baru = new Produk{kode, nama, harga, nullptr, nullptr};
        if (isEmpty()) {
            head = tail = baru;
        } else {
            tail->next = baru;
            baru->prev = tail;
            tail = baru;
        }
    }

    /**
     * @brief Tambah produk setelah produk dengan kode tertentu.
     * 
     * Kalau kodenya nggak ketemu, bakal muncul pesan error.
     */
    void insertAfter(string kodeSetelah, string kode, string nama, int harga) {
        if (isEmpty()) {
            cout << "List produk kosong.\n";
            return;
        }

        Produk* bantu = head;
        while (bantu != nullptr && bantu->kodeProduk != kodeSetelah)
            bantu = bantu->next;

        if (bantu == nullptr) {
            cout << "Produk dengan kode " << kodeSetelah << " tidak ditemukan.\n";
        } else {
            Produk* baru = new Produk{kode, nama, harga, bantu, bantu->next};
            if (bantu->next != nullptr)
                bantu->next->prev = baru;
            else
                tail = baru;
            bantu->next = baru;
        }
    }

    /**
     * @brief Nampilin semua produk dari depan ke belakang.
     */
    void tampilProduk() {
        if (isEmpty()) {
            cout << "Tidak ada produk.\n";
            return;
        }

        Produk* bantu = head;
        cout << "\nDaftar Produk:\n";
        while (bantu != nullptr) {
            cout << "- [" << bantu->kodeProduk << "] "
                 << bantu->namaProduk << " | Harga: " << bantu->harga << endl;
            bantu = bantu->next;
        }
    }
};

/* =======================================================
   BAGIAN TRANSAKSI — mirip kayak produk tapi beda data
   ======================================================= */

/**
 * @brief Struktur data buat nyimpen info transaksi.
 * 
 * Transaksi bisa debit atau kredit, dengan jumlah tertentu.
 */
struct Transaksi {
    int idTransaksi;    ///< ID unik buat transaksi
    string jenis;       ///< "Debit" atau "Kredit"
    int jumlah;         ///< jumlah uang di transaksi
    Transaksi* prev;    ///< pointer ke transaksi sebelumnya
    Transaksi* next;    ///< pointer ke transaksi berikutnya
};

/**
 * @brief Kelas buat ngatur daftar transaksi (Doubly Linked List).
 * 
 * Bisa nambah transaksi di akhir, hapus transaksi setelah ID tertentu,
 * dan nampilin semua transaksi.
 */
class TransaksiList {
private:
    Transaksi* head;  ///< transaksi pertama
    Transaksi* tail;  ///< transaksi terakhir

public:
    /// Konstruktor: mulaiin list kosong
    TransaksiList() {
        head = tail = nullptr;
    }

    /// Cek apakah list kosong
    bool isEmpty() {
        return head == nullptr;
    }

    /**
     * @brief Nambah transaksi di akhir list.
     */
    void insertLast(int id, string jenis, int jumlah) {
        Transaksi* baru = new Transaksi{id, jenis, jumlah, nullptr, nullptr};
        if (isEmpty()) {
            head = tail = baru;
        } else {
            tail->next = baru;
            baru->prev = tail;
            tail = baru;
        }
    }

    /**
     * @brief Hapus transaksi yang posisinya setelah ID tertentu.
     * 
     * Kalau ID nggak ketemu atau nggak ada transaksi setelahnya, muncul pesan.
     */
    void deleteAfter(int id) {
        if (isEmpty()) {
            cout << "List transaksi kosong.\n";
            return;
        }

        Transaksi* bantu = head;
        while (bantu != nullptr && bantu->idTransaksi != id)
            bantu = bantu->next;

        if (bantu == nullptr) {
            cout << "Transaksi dengan ID " << id << " tidak ditemukan.\n";
        } else if (bantu->next == nullptr) {
            cout << "Tidak ada transaksi setelah ID " << id << ".\n";
        } else {
            Transaksi* hapus = bantu->next;
            bantu->next = hapus->next;
            if (hapus->next != nullptr)
                hapus->next->prev = bantu;
            else
                tail = bantu;
            delete hapus;
            cout << "Transaksi setelah ID " << id << " berhasil dihapus.\n";
        }
    }

    /**
     * @brief Nampilin semua transaksi dari awal sampai akhir.
     */
    void viewAll() {
        if (isEmpty()) {
            cout << "Tidak ada transaksi.\n";
            return;
        }

        Transaksi* bantu = head;
        cout << "\nRiwayat Transaksi:\n";
        while (bantu != nullptr) {
            cout << "- ID: " << bantu->idTransaksi
                 << " | Jenis: " << bantu->jenis
                 << " | Jumlah: " << bantu->jumlah << endl;
            bantu = bantu->next;
        }
    }
};

/**
 * @brief Fungsi utama buat jalanin program.
 * 
 * Nyediain menu interaktif buat ngatur produk dan transaksi.
 */
int main() {
    ProdukList produkList;
    TransaksiList transaksiList;
    int pilihan;

    do {
        cout << "\n========== MENU ==========\n";
        cout << "1. Tambah Produk di Awal\n";
        cout << "2. Tambah Produk di Akhir\n";
        cout << "3. Tambah Produk Setelah Kode Tertentu\n";
        cout << "4. Lihat Semua Produk\n";
        cout << "--------------------------\n";
        cout << "5. Tambah Transaksi di Akhir\n";
        cout << "6. Hapus Transaksi Setelah ID Tertentu\n";
        cout << "7. Lihat Semua Transaksi\n";
        cout << "0. Keluar\n";
        cout << "Pilih menu: ";
        cin >> pilihan;
        cin.ignore();

        if (pilihan == 1) {
            string kode, nama;
            int harga;
            cout << "Kode Produk: "; getline(cin, kode);
            cout << "Nama Produk: "; getline(cin, nama);
            cout << "Harga: "; cin >> harga;
            produkList.insertFirst(kode, nama, harga);
        } 
        else if (pilihan == 2) {
            string kode, nama;
            int harga;
            cout << "Kode Produk: "; getline(cin, kode);
            cout << "Nama Produk: "; getline(cin, nama);
            cout << "Harga: "; cin >> harga;
            produkList.insertLast(kode, nama, harga);
        } 
        else if (pilihan == 3) {
            string kodeSetelah, kode, nama;
            int harga;
            cout << "Masukkan kode produk yang ingin disisipi setelahnya: ";
            getline(cin, kodeSetelah);
            cout << "Kode Produk Baru: "; getline(cin, kode);
            cout << "Nama Produk Baru: "; getline(cin, nama);
            cout << "Harga: "; cin >> harga;
            produkList.insertAfter(kodeSetelah, kode, nama, harga);
        } 
        else if (pilihan == 4) {
            produkList.tampilProduk();
        } 
        else if (pilihan == 5) {
            int id, jumlah;
            string jenis;
            cout << "ID Transaksi: "; cin >> id;
            cout << "Jenis (Debit/Kredit): "; cin >> jenis;
            cout << "Jumlah: "; cin >> jumlah;
            transaksiList.insertLast(id, jenis, jumlah);
        } 
        else if (pilihan == 6) {
            int id;
            cout << "Masukkan ID transaksi yang ingin dihapus sesudahnya: ";
            cin >> id;
            transaksiList.deleteAfter(id);
        } 
        else if (pilihan == 7) {
            transaksiList.viewAll();
        }

    } while (pilihan != 0);

    cout << "Program selesai.\n";
    return 0;
}
