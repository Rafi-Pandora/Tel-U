#include "pelajaran.h"

pelajaran_t Pelajaran::create_pelajaran(string namaMapel, string kodeMapel) {
  pelajaran_t pelajaran;
    pelajaran.namaMapel = namaMapel;
    pelajaran.kodeMapel = kodeMapel;

  return pelajaran;
}

void Pelajaran::tampil_pelajaran(pelajaran_t pelajaran) {
  cout << pelajaran.namaMapel << endl;
  cout << pelajaran.kodeMapel << endl;
}
