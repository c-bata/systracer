# systracer

Yet another system call tracer written in Go.
This is a sample repository for my talk about "How to write a system call tracer for Linux/x86." at Aizu University.

## Usage

Currently, this tool supports Linux/x86 only. Usage is like this:

```console
$ ./systracer-linux-386 --summary ./hello
Wait returned: stop signal: trace/breakpoint trap
Hello World! 1 ./hello
Summary:
        1|write
        1|execve
        4|brk
        1|uname
        1|mmap2
        1|fstat64
        1|set_thread_area
```

## How to work

[![how-to-work](https://github.com/c-bata/assets/raw/master/systrace/how-to-trace-system-calls.png)](#)

C implementation is [_example/linux-386-c/main.c](https://github.com/c-bata/systrace/blob/master/_example/linux-386-c/main.c).

