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

static void handler(int signum) {
	printf("crash signum:%d\n", signum);
	char cmd[50];
	sprintf(cmd, "gcore %u", getpid());
	system(cmd);
	char path[PATH_MAX];
	char processname[1024];
	get_executable_path(path, processname, sizeof(path));
	sprintf(cmd, "./gdb_print.sh ./%s ./core.%u", processname, getpid());
	system(cmd);
	exit(1);
}

static void __attribute__ ((constructor)) sigsetup(void) {
	struct sigaction act;
	memset(&act, 0, sizeof act);
	act.sa_handler = handler;
	act.sa_flags = SA_RESETHAND;
	sigaction(SIGSEGV, &act, NULL);
	sigaction(SIGABRT, &act, NULL);
}
*/
import "C"

func sigsetup() {
	C.sigsetup()
}
