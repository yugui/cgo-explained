#include <stdio.h>

#include "_cgo_export.h"

void print_go_version(void) {
  const GoString version = goVersion();
  printf("%.*s\n", (int)version.n, version.p);
}
