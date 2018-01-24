package systracer

import (
	"fmt"
	"os"
	"text/tabwriter"
)

type SyscallCounter []int

const maxSyscalls = 303

func NewCounter() SyscallCounter {
	return make(SyscallCounter, maxSyscalls)
}

func (s SyscallCounter) Inc(syscallID int32) error {
	if syscallID > maxSyscalls {
		return fmt.Errorf("invalid syscall ID (%x)", syscallID)
	}

	s[syscallID]++
	return nil
}

func (s SyscallCounter) Print() {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 8, ' ', tabwriter.AlignRight|tabwriter.Debug)
	for k, v := range s {
		if v > 0 {
			name, _ := GetSyscallName("x86", k)
			fmt.Fprintf(w, "%d\t%s\n", v, name)
		}
	}
	w.Flush()
}

func (s SyscallCounter) GetName(syscallID int) Syscall {
	name, _ := GetSyscallName("x86", syscallID)
	return name
}
