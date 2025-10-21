#ifndef DOUBLYLIST_H
#define DOUBLYLIST_H

#include <iostream>

class DoublyNode {
public:
    int data;
    DoublyNode* prev;
    DoublyNode* next;

    DoublyNode(int val);
};

class DoublyLinkedList {
private:
    DoublyNode* head;
    DoublyNode* tail;

public:
    DoublyLinkedList();
    ~DoublyLinkedList();

    void insertFirst(int val);
    void insertLast(int val);
    bool deleteAfter(int key);
    bool findElm(int key) const;
    void display() const;
};

#endif
