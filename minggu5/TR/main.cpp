#include "singlylist.h"

int main() {
    SinglyList list;
    int find = 8;

    list.insertFirst(5);
    list.insertFirst(3);
    list.insertFirst(10);
    list.insertFirst(8);
    list.insertFirst(5);

    if (list.findValue(find)) {
      cout << "Nilai " << find << " ditemukan dalam list" << endl;
    } else {
      cout << "Nilai " << find << " tidak ditemukan dalam list" << endl;
    }

    int total = list.sumAll();
    cout << "Total semua nilai di list = " << total << endl;

    return 0;
}
