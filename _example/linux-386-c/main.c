#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <signal.h>
#include <sys/ptrace.h>
#include <sys/wait.h>
#include <sys/user.h>

int main(int argc, char *argv[], char *envp[])
{
    int pid, status, cnt;
    struct user_regs_struct regs;

    pid = fork();
    if (!pid) { /* tracee process */
        close(1);
        ptrace(PTRACE_TRACEME, 0, NULL, NULL);
        execve(argv[1], argv + 1, envp); /* execute program */
    }

    while (1) { /* tracer loop */
        waitpid(pid, &status, 0);
        if (WIFEXITED(status))
            break;

        cnt++;
        if (cnt % 2) {
            ptrace(PTRACE_GETREGS, pid, NULL, &regs);
            fprintf(stderr, "System Call: %ld\n", (long int)regs.orig_eax);
        }
        ptrace(PTRACE_SYSCALL, pid, NULL, NULL);
    }
    exit(0);
}