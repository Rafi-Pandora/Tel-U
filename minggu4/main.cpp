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
};

class helper {
    public:

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

    static void soal1_manual(single_linked_list& node) {
        bool loop = true;
        int pilihan = 0;

        while (loop) {
            cout << "\n========== Menu Linked List ==========\n";
            cout << "1. Insert First\n";
            cout << "2. Insert Last\n";
            cout << "3. Keluar\n";
            cout << "Pilih menu : ";
            cin >> pilihan;

            switch (pilihan) {
                case 1:
                    for (int i = 0; i < 3; i++) {
                        int data;
                        cout << "Data " << i + 1 << " : ";
                        cin >> data;
                        node.insert_first(data);
                    }
                    node.tampilkan();
                    break;

                case 2:
                    for (int i = 0; i < 3; i++) {
                        int data;
                        cout << "Data " << i + 1 << " : ";
                        cin >> data;
                        node.insert_last(data);
                    }
                    node.tampilkan();
                    break;

                case 3:
                    loop = false;
                    break;

                default:
                    cout << "Input salah!" << endl;
                    break;
            }
        }
    }

    static void soal2_manual(single_linked_list& node) {
        string input;
        cout << "Masukkan 5 angka (pisahkan dengan koma): ";
        cin.ignore();
        std::getline(cin, input);

        std::stringstream ss(input);
        std::vector<int> numbers;
        string token;

        while (std::getline(ss, token, ',')) {
            numbers.push_back(std::stoi(token));
        }

        for (int n : numbers) {
            node.insert_last(n);
        }

        node.tampilkan();
        node.delete_first();
        node.delete_first();
        node.tampilkan();
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
    else if (argc > 1 && string(argv[1]) == "soal1") {
        helper::soal1_manual(node);
    }
    else if (argc > 1 && string(argv[1]) == "soal2") {
        helper::soal2_manual(node);
    }
    else if (argc > 1 && (string(argv[1]) == "-v" || string(argv[1]) == "--version")) {
        version::print_version(argv[0]);
    }
    else {
        version::print_help(argv[0]);
    }

    return 0;
}
