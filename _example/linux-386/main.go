// +build linux,386
package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"syscall"

	"github.com/c-bata/systracer"
)

func main() {
	var outputSummary bool
	flag.BoolVar(&outputSummary, "summary", false, "If true, just output the summary")
	flag.Parse()

	var regs syscall.PtraceRegs
	var ss = systracer.NewCounter()

	args := flag.Args()
	cmd := exec.Command(args[0], args[1:]...)
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

			if !outputSummary {
				name, _ := systracer.GetSyscall(int(regs.Orig_eax))
				fmt.Printf("\t%s\n", name)
			}

			ss.Inc(int(regs.Orig_eax))
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
	if outputSummary {
		fmt.Println("Summary:")
		ss.Print()
	}
}
