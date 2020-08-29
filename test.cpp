#include <stdio.h>
#include <assert.h>
#include <string.h>
#include <stdlib.h>
#include "test.h"

void fn2(char *arg)
{
    int stackvar2 = 256;
    printf("Argument %s\n", arg);
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
    printf("test_crash from C and here the str is from Go: %s\n", str);
    fn1(1092);
    printf("xxxx");
}

class A
{
public:
    void alloc(int n) { ptr = malloc(n); }
    void release() { free(ptr); }

private:
    void *ptr;
};

// 空指针调用
void test_crash2()
{
    A *a = (A *)0;
    a->release();
}