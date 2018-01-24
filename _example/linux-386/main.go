// +build linux,386
package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"

	"github.com/c-bata/systracer"
)

func main() {
	var regs syscall.PtraceRegs
	var ss = systracer.NewCounter()

	cmd := exec.Command(os.Args[1], os.Args[2:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Ptrace: true,
	}

	cmd.Start()
	err := cmd.Wait()
	if err != nil {
		fmt.Printf("Wait returned: %v\n", err)
	}

	pid := cmd.Process.Pid
	exit := true

	for {
		if exit {
			err = syscall.PtraceGetRegs(pid, &regs)
			if err != nil {
				break
			}

			name, _ := systracer.GetSyscallName("x86", int(regs.Orig_eax))
			fmt.Printf("\t%s\n", name)

			ss.Inc(regs.Orig_eax)
		}

		err = syscall.PtraceSyscall(pid, 0)
		if err != nil {
			panic(err)
		}

		_, err = syscall.Wait4(pid, nil, 0, nil)
		if err != nil {
			panic(err)
		}

		exit = !exit
	}
	fmt.Println("Summary:")
	ss.Print()
}
