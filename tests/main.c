#include <stdio.h>

extern long convert(void);

int main() 
{
    long a = convert();
    printf("%ld\n", a);

    return 0;
}