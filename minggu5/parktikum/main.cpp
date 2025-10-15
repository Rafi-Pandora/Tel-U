#include <iostream>
#include <sstream>
#include <random>
#include <vector>
#include <string>
#include "version.h"

using std::cout;
using std::cin;
using std::endl;
using std::string;

struct single_linked_list_t {
    int data;
    single_linked_list_t* next;
};

class single_linked_list {
private:
    single_linked_list_t* head;

public:
    single_linked_list() {
        head = nullptr;
    }

    ~single_linked_list() {
        single_linked_list_t* current = head;
        while (current != nullptr) {
            single_linked_list_t* next_node = current->next;
            delete current;
            current = next_node;
        }
    }

    void insert_first(int data) {
        auto* new_node = new single_linked_list_t();
        new_node->data = data;
        new_node->next = head;
        head = new_node;
    }

    void insert_last(int data) {
        auto* new_node = new single_linked_list_t();
        new_node->data = data;
        new_node->next = nullptr;

        if (head == nullptr) {
            head = new_node;
            return;
        }

        auto* temp = head;
        while (temp->next != nullptr) {
            temp = temp->next;
        }
        temp->next = new_node;
    }

    void delete_first() {
        if (head == nullptr) {
            cout << "List kosong." << endl;
            return;
        }

        auto* temp = head;
        head = head->next;
        cout << "Menghapus data: [" << temp->data << "]" << endl;
        delete temp;
    }

    void search_node(int data) {
        if (head == nullptr) return;

        auto* temp = head;
        while (temp != nullptr)
        {
            if (temp->data == data) {
                cout << "data ditemukan" << endl;
                return;
            }
            temp = temp->next;
        }
        cout << "data tidak ditemukan" << endl;
    }

    void search_node_with_poss(int data) {
        if (head == nullptr) return;

        int poss = 1;
        auto* temp = head;
        while (temp != nullptr)
        {
            if (temp->data == data) {
                cout << "data ditemukan pada node ke - " << poss << endl;
                return;
            }
            temp = temp->next;
            poss++;
        }
        cout << "data tidak ditemukan" << endl;
    }

    void tampilkan() {
        if (head == nullptr) {
            cout << "List kosong." << endl;
            return;
        }

        auto* temp = head;
        cout << "Isi Linked List: ";
        while (temp != nullptr) {
            cout << "[" << temp->data << "] -> ";
            temp = temp->next;
        }
        cout << "NULL" << endl;
    }

    bool update_node(int oldValue, int newValue) {
        if (head == nullptr) return false;

        auto* temp = head;
        while (temp != nullptr)
        {
            if (temp->data == oldValue) {
                temp->data = newValue;
                return true;
            }
            temp = temp->next;
        }
        cout << "data tidak ditemukan" << endl;
        tampilkan();
        return false;
    }
};

class helper {
public:
    static void soal1_manual(single_linked_list& node) {
        int jumlah_data = 0;
        cout << "Berapa banyak data yang ingin dimasukkan? ";
        cin >> jumlah_data;

        for (int i = 0; i < jumlah_data; i++) {
            int data;
            cout << "Data ke-" << (i + 1) << ": ";
            cin >> data;
            node.insert_first(data); 
        }

        node.tampilkan();

        int cari;
        cout << "Masukkan data yang ingin dicari: ";
        cin >> cari;

        
        node.search_node(cari); 
    }

    
    static void soal2_manual(single_linked_list& node) {
        int jumlah_data = 0;
        cout << "Berapa banyak data yang ingin dimasukkan? ";
        cin >> jumlah_data;

        for (int i = 0; i < jumlah_data; i++) {
            int data;
            cout << "Data ke-" << (i + 1) << ": ";
            cin >> data;
            node.insert_first(data); 
        }

        node.tampilkan();

        int cari;
        cout << "Masukkan data yang ingin dicari posisinya: ";
        cin >> cari;

        
        node.search_node_with_poss(cari); 
    }

   
    static void soal3_manual(single_linked_list& node) {
        int jumlah_data = 0;
        cout << "Berapa banyak data yang ingin dimasukkan? ";
        cin >> jumlah_data;

        for (int i = 0; i < jumlah_data; i++) {
            int data;
            cout << "Data ke-" << (i + 1) << ": ";
            cin >> data;
            node.insert_first(data); 
        }

        
        int oldValue, newValue;
        cout << "Masukkan nilai lama (yang ingin diubah): ";
        cin >> oldValue;
        cout << "Masukkan nilai baru: ";
        cin >> newValue;
        
        node.update_node(oldValue, newValue); 
    }

    
    static void soal1_auto(single_linked_list& node) {
        std::random_device rd;
        std::mt19937 gen(rd());
        std::uniform_int_distribution<> dist(3, 7);

        int n_loop = dist(gen);

        for (int i = 0; i < n_loop; i++) node.insert_first(dist(gen));
        node.tampilkan();

        for (int i = 0; i < n_loop; i++) node.insert_last(dist(gen));
        node.tampilkan();
    }

    static void soal2_auto(single_linked_list& node) {
        std::random_device rd;
        std::mt19937 gen(rd());
        std::uniform_int_distribution<> dist(3, 7);

        int n_loop = dist(gen);

        for (int i = 0; i < n_loop; i++) node.insert_last(dist(gen));
        node.tampilkan();

        for (int i = 0; i < n_loop - 2; i++) node.delete_first();
        node.tampilkan();
    }

    static void soal3_auto(single_linked_list& node) {
        std::random_device rd;
        std::mt19937 gen(rd());
        std::uniform_int_distribution<> dist_size(3, 7);  
        std::uniform_int_distribution<> dist_value(10, 30); 
        std::uniform_int_distribution<> dist_update(10, 30); 
        int n_loop = dist_size(gen);

        for (int i = 0; i < n_loop; i++) {
            node.insert_first(dist_value(gen));
        }

        node.tampilkan();

        int oldValue = dist_update(gen);
        int newValue = dist_update(gen);

        cout << "Update semua node yang bernilai " << oldValue << " menjadi " << newValue << endl;

        node.update_node(oldValue, newValue);
    }
};



int main(int argc, char const* argv[]) {
    single_linked_list node;

    if (argc > 2 && string(argv[1]) == "soal1" && string(argv[2]) == "auto") {
        helper::soal1_auto(node);
    }
    else if (argc > 2 && string(argv[1]) == "soal2" && string(argv[2]) == "auto") {
        helper::soal2_auto(node);
    }
    else if (argc > 2 && string(argv[1]) == "soal3" && string(argv[2]) == "auto") {
        helper::soal3_auto(node);
    }
    else if (argc > 1 && string(argv[1]) == "soal1") {
        helper::soal1_manual(node);
    }
    else if (argc > 1 && string(argv[1]) == "soal2") {
        helper::soal2_manual(node);
    }
    else if (argc > 1 && string(argv[1]) == "soal3") {
        helper::soal3_manual(node);
    }
    else if (argc > 1 && (string(argv[1]) == "-v" || string(argv[1]) == "--version")) {
        version::print_version(argv[0]);
    }
    else {
        version::print_help(argv[0]);
    }

    return 0;
}
