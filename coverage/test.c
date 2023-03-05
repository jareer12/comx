#include <stdio.h>
#include "./module.c"

int main()
{
    for (size_t i = 0; i < 10; i++)
    {
        printf("%lu", i);
    }

    return 0;
}