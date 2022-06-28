#include "funk.c"
#include <stdio.h>

int main() {

  int l[6] = {1,2,5,4,87,6};
  struct CData cData;
  cData.list = l;
  int i;
  int s = sizeof(l) / sizeof(int);

  cBubbleSort(&cData, s);
  for (i = 0; i < sizeof(l) / sizeof(int); i++)
      printf("%d ", l[i]);

  return 0;
}
