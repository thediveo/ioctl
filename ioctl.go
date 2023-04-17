// Copyright 2023 Harald Albrecht.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ioctl

import (
	"errors"

	"golang.org/x/sys/unix"
)

/*
Ugly IOCTL stuff.

ATTENTION: the [following definitions] hold only for the "asm-generic"
platforms, such as x86, arm, and others. Currently the only platforms having a
different ioctl request field mapping are: alpha, mips, powerpc, and sparc.

We keep the original C identifiers on purpose and don't care about linters
trying to patronizing us.

[following definitions]: https://elixir.bootlin.com/linux/v6.2.11/source/include/uapi/asm-generic/ioctl.h
*/
const (
	_IOC_NRBITS   = 8
	_IOC_TYPEBITS = 8
	_IOC_SIZEBITS = 14

	_IOC_NRSHIFT   = 0
	_IOC_TYPESHIFT = _IOC_NRSHIFT + _IOC_NRBITS
	_IOC_SIZESHIFT = _IOC_TYPESHIFT + _IOC_TYPEBITS
	_IOC_DIRSHIFT  = _IOC_SIZESHIFT + _IOC_SIZEBITS

	_IOC_NONE = uint(0)
)

// Returns an ioctl(2) request value, calculated from the specific ioctl call
// properties: parameter in/out direction, type of ioctl, command number, and
// finally parameter size.
func _IOC(dir, ioctype, nr, size uint) uint {
	return (dir << _IOC_DIRSHIFT) | (ioctype << _IOC_TYPESHIFT) | (nr << _IOC_NRSHIFT) | (size << _IOC_SIZESHIFT)
}

// Returns an ioctl(2) request value for a request that doesn't have any
// additional request parameter.
func _IO(ioctype, nr uint) uint {
	return _IOC(_IOC_NONE, ioctype, nr, 0)
}

// RetFd issues the specified ioctĺ request and returns the successful result as
// a file descriptor, or an error. In contrast to [unix.IoctlRetInt], RetFd
// returns an invalid file descriptor -1
func RetFd(fd int, request uint) (int, error) {
	nsfd, _, errno := unix.Syscall(unix.SYS_IOCTL,
		uintptr(fd), uintptr(request), uintptr(0))
	if errno != 0 {
		return -1, errors.New(errno.Error())
	}
	return int(nsfd), nil
}
