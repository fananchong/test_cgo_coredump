package main

/*
#include <stdio.h>
#include <signal.h>
#include <stdlib.h>
#include <string.h>

extern void OnExcept();

static void segvhandler(int signum)
{
    OnExcept();
    printf("crash !!!\n");
    exit(1);
}

void sigsetup1(void)
{
    printf("sigsetup ...\n");
    struct sigaction act;

    memset(&act, 0, sizeof act);
    act.sa_handler = segvhandler;
    sigaction(SIGSEGV, &act, NULL);
    sigaction(SIGABRT, &act, NULL);
}
*/
import "C"

func sigsetup() {
	C.sigsetup1()
}
