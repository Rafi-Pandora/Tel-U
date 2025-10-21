#include "doublylist.h"

DoublyNode::DoublyNode(int val) : data(val), prev(nullptr), next(nullptr) {}

DoublyLinkedList::DoublyLinkedList() : head(nullptr), tail(nullptr) {}

DoublyLinkedList::~DoublyLinkedList() {
    DoublyNode* cur = head;
    while (cur) {
        DoublyNode* next = cur->next;
        delete cur;
        cur = next;
    }
}

void DoublyLinkedList::insertFirst(int val) {
    DoublyNode* node = new DoublyNode(val);
    if (!head) {
        head = tail = node;
    } else {
        node->next = head;
        head->prev = node;
        head = node;
    }
}

void DoublyLinkedList::insertLast(int val) {
    DoublyNode* node = new DoublyNode(val);
    if (!tail) {
        head = tail = node;
    } else {
        tail->next = node;
        node->prev = tail;
        tail = node;
    }
}

bool DoublyLinkedList::deleteAfter(int key) {
    if (!head) return false;

    DoublyNode* cur = head;
    while (cur && cur->data != key) {
        cur = cur->next;
    }

    if (!cur || !cur->next) return false;

    DoublyNode* delNode = cur->next;
    cur->next = delNode->next;
    if (delNode->next)
        delNode->next->prev = cur;
    else
        tail = cur;

    delete delNode;
    return true;
}

bool DoublyLinkedList::findElm(int key) const {
    DoublyNode* cur = head;
    while (cur) {
        if (cur->data == key) return true;
        cur = cur->next;
    }
    return false;
}

void DoublyLinkedList::display() const {
    if (!head) {
        std::cout << "List kosong\n";
        return;
    }

    DoublyNode* cur = head;
    std::cout << "Isi list: ";
    while (cur) {
        std::cout << cur->data;
        if (cur->next) std::cout << " ";
        cur = cur->next;
    }
    std::cout << "\n";
}
