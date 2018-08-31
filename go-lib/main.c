#include "add.h"
#include <stdio.h>

int main() {
    int a = 10;
    int b = 5;

    int x = add(a, b);

    printf("%d + %d = %d\n", a, b, x);

    return 0;
}
