#ifndef SINGLYLIST_H
#define SINGLYLIST_H

#include <iostream>
using namespace std;

/**
 * @class SinglyList
 * @brief Kelas untuk bikin dan ngatur single linked list sederhana.
 * 
 * List ini nyimpen data bertipe `int` dan punya operasi dasar kayak 
 * nambah elemen di awal, ngecek list kosong, sama nampilin isinya.
 */
class SinglyList {
private:
    /**
     * @struct Node
     * @brief Struktur buat nyimpen satu elemen list.
     */
    struct Node {
        int info;   ///< Data utama yang disimpen
        Node* next; ///< Pointer ke node berikutnya

        /**
         * @brief Constructor buat Node.
         * @param val nilai awal yang mau disimpen
         */
        Node(int val) : info(val), next(nullptr) {}
    };
    
    Node* first; ///< Pointer ke node pertama (head list)

public:
    /**
     * @brief Constructor utama. Langsung bikin list kosong.
     */
    SinglyList();

    /**
     * @brief Destructor. Bersihin semua node biar gak ada memory leak.
     */
    ~SinglyList();

    /**
     * @brief Nambah elemen baru di awal list.
     * @param value Nilai yang mau ditambah.
     */
    void insertFirst(int value);

    /**
     * @brief Nampilin semua isi list ke layar.
     */
    void printInfo() const;

    /**
     * @brief Cek apakah list masih kosong atau udah ada isinya.
     * @return true kalau kosong, false kalau ada isi.
     */
    bool isEmpty() const;

private:
    /**
     * @brief Hapus 1 node dari memori.
     * @param P pointer ke node yang mau dihapus.
     */
    void dealokasi(Node* P);
};

#endif
