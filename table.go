package systracer

import "errors"

var ErrUndefinedArchitecture = errors.New("undefined architecture")

func GetSyscallName(arch string, num int) (Syscall, error) {
	switch arch {
	case "x86":
		return GetX86SyscallName(num)
	default:
		return SyscallUndefined, ErrUndefinedArchitecture
	}
}
