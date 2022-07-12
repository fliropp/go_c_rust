#include "../go_rust/dist/lib.h"
#include <stdio.h>

int main() {
  int l[6] = {1,2,5,4,87,6};
  int *pl = l;
  int s = sizeof(l) / sizeof(int);
  struct RData rd = new_rd_data(pl, s);
  int i;
  printf("Before Bubble Sort:\n");
  for (i = 0; i < s; i++) {
    printf("%d ", rd.list[i]);
  }
  rBubbleSort(&rd, s, s);
  printf("\nAfter Bubble Sort:\n");
  for (i = 0; i < s; i++) {
    printf("%d ", rd.list[i]);
  }
  return 0;
}
