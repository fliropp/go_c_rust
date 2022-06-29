#include "funk.c"
#include <stdio.h>

int main() {

  int l[6] = {1,2,5,4,87,6};
  int *pl = l;
  int s = sizeof(l) / sizeof(int);
  struct CData *cd = createCData(cd, pl, s);
  int i;
  for (i = 0; i < s; i++) {
    printf("item: %d", cd->list[i]);
  }
  return 0;
}
