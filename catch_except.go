// +build !plan9,!windows

package main

/*
#include <signal.h>
#include <stdlib.h>
#include <string.h>
#include <stdio.h>
#include <sys/types.h>
#include <unistd.h>
#include <limits.h>

static size_t get_executable_path( char* processdir,char* processname, size_t len)
{
	char* path_end;
	if(readlink("/proc/self/exe", processdir,len) <=0)
		return -1;
	path_end = strrchr(processdir,  '/');
	if(path_end == NULL)
		return -1;
	++path_end;
	strcpy(processname, path_end);
	*path_end = '\0';
	return (size_t)(path_end - processdir);
}

static void print_core(int signum, siginfo_t *info, void *secret, struct sigaction *oldact) {
	printf("crash signum:%d si_code:%d\n", signum, info->si_code);
	char cmd[50];
	sprintf(cmd, "gcore %u", getpid());
	system(cmd);
	char path[PATH_MAX];
	char processname[1024];
	get_executable_path(path, processname, sizeof(path));
	sprintf(cmd, "./gdb_print.sh ./%s ./core.%u", processname, getpid());
	system(cmd);
	oldact->sa_sigaction(signum, info, secret);
}

static struct sigaction oldabrtact;
static void abrtsigaction(int signum, siginfo_t *info, void *secret) {
	print_core(signum, info, secret, &oldabrtact);
}

static struct sigaction oldsegvact;
static void segvsigaction(int signum, siginfo_t *info, void *secret) {
	print_core(signum, info, secret, &oldsegvact);
}

static void sigsetup(void) {
	struct sigaction act;
	memset(&act, 0, sizeof act);
	act.sa_flags = SA_ONSTACK | SA_SIGINFO;
	act.sa_sigaction = segvsigaction;
	sigaction(SIGSEGV, &act, &oldsegvact);
	act.sa_sigaction = abrtsigaction;
	sigaction(SIGABRT, &act, &oldabrtact);
}
*/
import "C"

// Sigsetup Sigsetup
func Sigsetup() {
	C.sigsetup()
}
