/*

Test, the C bindings.

Warning, this file is autogenerated by `cbindgen`.
Do not modify this manually.

*/

#include <stdarg.h>
#include <stdbool.h>
#include <stdint.h>
#include <stdlib.h>

typedef struct RData {
    uintptr_t size;
    int32_t *list;
} RData;

int32_t rLoop(void);

int32_t rfib(int32_t n);

struct RData new_rd_data(int32_t *ptr_lzt, uintptr_t size);

void rBubbleSort(struct RData *ptr_data, uintptr_t n, uintptr_t c);
