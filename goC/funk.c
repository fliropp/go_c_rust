//funk.c
#include <stdio.h>
#include <stdlib.h>

int cLoop() {
  int s = 0;
  int i;
  for ( i = 0; i < 10001; i++) {
    s = s + i;
  }
  return s;
}

void swap(int* xp, int* yp)
{
    int temp = *xp;
    *xp = *yp;
    *yp = temp;
}

struct CData {
  int size;
  int list[];
};

struct CData *createCData(struct CData *cd, int list[], int size) {
    cd = malloc( sizeof(*cd) + sizeof(int) * size);
    int i;
    cd->size = size;
    for (i = 0; i < size; i++) {
      cd->list[i] = list[i];
    }
    return cd;

}

void cBubbleSort(struct CData *data, int s) {
  int i,j;
  for (int i = 0; i < s - 1; i++) {
    for (int j = 0; j < s - i - 1; j++) {
      if (data->list[j] > data->list[j + 1]) {
        swap(&data->list[j], &data->list[j + 1]);
      }
    }
  }

}
