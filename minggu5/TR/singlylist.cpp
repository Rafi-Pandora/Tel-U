#include "singlylist.h"

/**
 * @brief Constructor: inisialisasi list kosong.
 */
SinglyList::SinglyList() {
    first = nullptr;
}

/**
 * @brief Destructor: hapus semua elemen list satu per satu.
 */
SinglyList::~SinglyList() {
    Node* current = first;
    while (current != nullptr) {
        Node* temp = current;
        current = current->next;
        dealokasi(temp);
    }
}

/**
 * @brief Tambah elemen di depan list.
 * @param value nilai integer yang mau ditambah.
 */
void SinglyList::insertFirst(int value) {
    Node* newNode = new Node(value);
    newNode->next = first;
    first = newNode;
    cout << value << " ";
}

/**
 * @brief Tampilkan isi list ke console.
 */
void SinglyList::printInfo() const {
    if (first == nullptr) {
        cout << "List kosong." << endl;
        return;
    }
    cout << "Isi List: ";
    Node* current = first;
    while (current != nullptr) {
        cout << current->info << " ";
        current = current->next;
    }
    cout << endl;
}

/**
 * @brief Cek apakah list kosong.
 * @return true kalau kosong, false kalau enggak.
 */
bool SinglyList::isEmpty() const {
    return first == nullptr;
}

/**
 * @brief Cari nilai tertentu di dalam list.
 * 
 * @param target nilai yang mau dicari.
 * @return true kalau ketemu, false kalau enggak.
 */
bool SinglyList::findValue(int target) const {
    Node* current = first;
    while (current != nullptr) {
        if (current->info == target) {
            cout << "Nilai " << target << " ketemu di list!" << endl;
            return true;
        }
        current = current->next;
    }
    cout << "Nilai " << target << " tidak ketemu di list." << endl;
    return false;
}

/**
 * @brief Hitung total dari semua nilai di list.
 * @return jumlah seluruh nilai info di setiap node.
 */
int SinglyList::sumAll() const {
    int total = 0;
    Node* current = first;

    if (current == nullptr) {
        cout << "List kosong, total = 0" << endl;
        return 0;
    }

    while (current != nullptr) {
        total += current->info;
        current = current->next;
    }

    cout << "Total semua nilai di list = " << total << endl;
    return total;
}

/**
 * @brief Dealokasi (hapus) satu node dari memori.
 * @param P node yang mau dihapus.
 */
void SinglyList::dealokasi(Node* P) {
    delete P;
}
