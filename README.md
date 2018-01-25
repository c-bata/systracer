# systracer

Yet another system call tracer written in Go.
This is a sample repository for my talk about "How to develop system call tracer." at Aizu University.

## Usage

Currently, this tool supported Linux/x86 only. Usage is like this:

```console
$ ./systracer-linux-386 ./hello
Wait returned: stop signal: trace/breakpoint trap
	execve
	uname
	brk
	brk
	set_thread_area
	brk
	brk
	fstat64
	mmap2
Hello World! 1 ./hello
	write
```

`--summary` option is available like following:

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
