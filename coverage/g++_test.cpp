#include <iostream>

int main()
{
    long unsigned int i = __UINT64_MAX__;
    int8_t ti = 0;

    printf("%lu", i);
    printf("%d", ti);

    return 0;
}