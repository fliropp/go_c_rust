//funk.c
#include <stdio.h>

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
  int list[9];
};

void cBubbleSort(struct CData *data, int s) {
  int i,j;
  for (int i = 0; i < s - 1; i++) {
    for (int j = 0; j < s - i - 1; j++) {
      if (data->list[j] > data->list[j + 1]) {
        swap(&data->list[j], &data->list[j + 1]);
      }
    }
  }

/*void cBubbleSort(int *list, int s) {
  int i,j;
  for (int i = 0; i < s - 1; i++) {
    for (int j = 0; j < s - i - 1; j++) {
      if (list[j] > list[j + 1]) {
        swap(&list[j], &list[j + 1]);
      }
    }
  }*/
}
