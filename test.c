#include <stdio.h>
#include <assert.h>
#include <string.h>

void fn2(char *arg)
{
    int stackvar2 = 256;
    printf("Argument %s", arg);
    assert(1 == 2);
}

void fn1(int arg)
{
    int stackvar3 = 512;
    char var[256];
    strcpy(var, "deadbeef");
    fn2(var);
}

void test_crash(char *str)
{
    printf("test_crash from C and here the str is from Go: %s", str);
    fn1(1092);
    printf("xxxx");
}