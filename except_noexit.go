//go:build !plan9 && !windows
// +build !plan9,!windows

package main

/*
#include <signal.h>
#include <stdlib.h>
#include <string.h>
#include <stdio.h>
#include <sys/types.h>
#include <unistd.h>
#include <setjmp.h>

static jmp_buf env;

static void handle(int signum, siginfo_t *info, void *secret) {
	printf("crash signum:%d si_code:%d\n", signum, info->si_code);
	longjmp(env, 1);
}

static void sigsetup2(void) {
	struct sigaction act;
	memset(&act, 0, sizeof act);
	act.sa_flags = SA_ONSTACK | SA_SIGINFO;
	act.sa_sigaction = handle;
	sigaction(SIGSEGV, &act, 0);
	sigaction(SIGABRT, &act, 0);
}

typedef void(*cb)(void);

static int mysetjmp(cb f) {
	int r = setjmp(env);
	if (r == 0) {
		f();
	} else {
		printf("异常后恢复\n");
	}
}

*/
import "C"

// Sigsetup Sigsetup
func Sigsetup2() {
	C.sigsetup2()
}

func SafeCall(f C.cb) {
	C.mysetjmp(f)
}
