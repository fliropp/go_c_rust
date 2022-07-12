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

void swap(int* xp, int* yp) {
    int temp = *xp;
    *xp = *yp;
    *yp = temp;
}

typedef struct CData {
  int size;
  int *list;
} CData;

struct CData *createCData(CData *cd, int *list, int size_list) {
    cd = malloc(sizeof(CData));
    cd->list = (int*) malloc(sizeof(int) * size_list);
    cd->size = size_list;

    int i;
    for (i = 0; i < size_list; i++) {
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
