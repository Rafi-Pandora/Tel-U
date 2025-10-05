#ifndef _PELAJARAN_H_
#define _PELAJARAN_H_

#include <iostream>

using std::string;
using std::cout;
using std::endl;
using std::cin;

struct pelajaran_t {
  string namaMapel;
  string kodeMapel;
};

class Pelajaran {
  public:
    pelajaran_t create_pelajaran(string namaMapel, string kodeMapel);
    void tampil_pelajaran(pelajaran_t pelajaran);
};

#endif // !_PELAJARAN
