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
 * @brief Dealokasi (hapus) satu node dari memori.
 * @param P node yang mau dihapus.
 */
void SinglyList::dealokasi(Node* P) {
    delete P;
}
