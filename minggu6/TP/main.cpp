#include "doublylist.h"
#include <iostream>
using namespace std;

int main() {
    DoublyLinkedList dll;
    int pilih, nilai, acuan;

    do {
        cout << "\n=== MENU DOUBLY LINKED LIST ===\n";
        cout << "1. Insert First\n";
        cout << "2. Insert Last\n";
        cout << "3. Delete After (berdasarkan nilai acuan)\n";
        cout << "4. Cari Nilai\n";
        cout << "5. Tampilkan List\n";
        cout << "0. Keluar\n";
        cout << "Pilih menu: ";
        cin >> pilih;

        switch (pilih) {
            case 1:
                cout << "Masukkan nilai: ";
                cin >> nilai;
                dll.insertFirst(nilai);
                cout << "Data " << nilai << " ditambahkan di awal.\n";
                break;

            case 2:
                cout << "Masukkan nilai: ";
                cin >> nilai;
                dll.insertLast(nilai);
                cout << "Data " << nilai << " ditambahkan di akhir.\n";
                break;

            case 3:
                cout << "Masukkan nilai acuan: ";
                cin >> acuan;
                if (dll.deleteAfter(acuan))
                    cout << "Node setelah nilai " << acuan << " berhasil dihapus.\n";
                else
                    cout << "Gagal menghapus. Nilai acuan tidak ditemukan atau tidak ada node setelahnya.\n";
                break;

            case 4:
                cout << "Masukkan nilai yang dicari: ";
                cin >> nilai;
                if (dll.findElm(nilai))
                    cout << "Data " << nilai << " ditemukan dalam list.\n";
                else
                    cout << "Data " << nilai << " tidak ditemukan.\n";
                break;

            case 5:
                dll.display();
                break;

            case 0:
                cout << "Program selesai.\n";
                break;

            default:
                cout << "Pilihan tidak valid.\n";
        }
    } while (pilih != 0);

    return 0;
}
