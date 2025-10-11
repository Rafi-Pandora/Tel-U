#include <iostream>
#include <random>
#include <string> 

using std::cout;
using std::cin;
using std::endl;

struct single_linked_list_t {
    int data;
    single_linked_list_t* next;
};

class Single_Linked_list {
private:
    single_linked_list_t* head;

public:
    Single_Linked_list() {
        head = nullptr;
    }

    ~Single_Linked_list() {
        single_linked_list_t* current = head;
        while (current != nullptr) {
            single_linked_list_t* nextNode = current->next;
            delete current;
            current = nextNode;
        }
    }

    void insert_first(int data) {
        single_linked_list_t* newNode = new single_linked_list_t();
        newNode->data = data;
        newNode->next = head;
        head = newNode;
    }

    void insert_last(int data) {
        single_linked_list_t* newNode = new single_linked_list_t;
        newNode->data = data;
        newNode->next = nullptr;

        if (head == nullptr) {
            head = newNode;
            return;
        }

        single_linked_list_t* temp = head;
        while (temp->next != nullptr) {
            temp = temp->next;
        }
        temp->next = newNode;
    }

    void delete_first_list() {
        if (head == nullptr) {
            cout << "list kosoong" << endl;
            return;
        }

        single_linked_list_t* temp = head;
        head = head->next;
        cout << "Menghapus data: [" << temp->data << "]" << endl;
        delete temp;
    }

    void tampilkan_list() {
        if (head == nullptr) {
            cout << "List kosong." << endl;
            return;
        }

        single_linked_list_t* temp = head;
        cout << "Isi Linked List: ";
        while (temp != nullptr) {
            cout << "[" << temp->data << "] -> ";
            temp = temp->next;
        }
        cout << "NULL" << endl;
    }
};

int main(int argc, char const *argv[]) {
    Single_Linked_list node;

    if (argc > 1 && std::string(argv[1]) == "auto") {
        std::random_device rd;
        std::mt19937 gen(rd());
        std::uniform_int_distribution<> dist(0, 254);

        for (int i = 0; i < 6; i++) {
            node.insert_first(dist(gen));
        }
        node.tampilkan_list();

        for (int i = 0; i < 6; i++) {
            node.insert_last(dist(gen));
        }
        node.tampilkan_list();

        return;
    }

    return 0;
}
