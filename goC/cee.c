#include "funk.c"
#include <stdio.h>

int main() {
  int l[6] = {1,2,5,4,87,6};
  int *pl = l;
  int s = sizeof(l) / sizeof(int);
  struct CData *cd = createCData(cd, pl, s);
  int i;
  printf("Before Bubble Sort:\n");
  for (i = 0; i < s; i++) {
    printf("%d ", cd->list[i]);
  }
  cBubbleSort(cd, s);
  printf("\nAfter Bubble Sort:\n");
  for (i = 0; i < s; i++) {
    printf("%d ", cd->list[i]);
  }
  return 0;
}
