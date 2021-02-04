// +build !plan9,!windows

package main

/*
#include <signal.h>
#include <stdlib.h>
#include <string.h>

extern void onExcept(int signum);

static void abrthandler(int signum) {
	onExcept(signum);
}

static void segvhandler(int signum) {
	onExcept(signum);
	exit(0);
}

static void __attribute__ ((constructor)) sigsetup(void) {
	struct sigaction act;
	memset(&act, 0, sizeof act);
	act.sa_handler = segvhandler;
	sigaction(SIGSEGV, &act, NULL);
	act.sa_handler = abrthandler;
	sigaction(SIGABRT, &act, NULL);
}
*/
import "C"

func sigsetup() {
	C.sigsetup()
}
