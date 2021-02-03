#include "catch_except.h"
#include <stdio.h>
#include <signal.h>
#include <stdlib.h>
#include <string.h>

static void segvhandler(int signum)
{
    printf("crash !!!!\n");
    exit(1);
}

void sigsetup(void)
{
    printf("sigsetup ...\n");
    struct sigaction act;

    memset(&act, 0, sizeof act);
    act.sa_handler = segvhandler;
    sigaction(SIGSEGV, &act, NULL);
    sigaction(SIGABRT, &act, NULL);
}
