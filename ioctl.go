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

//go:build linux && !(wasm || mips || mipsle || mips64 || mips64le || mips64p32 || mips64p32le || ppc || ppc64 || ppc64le || sparc || sparc64)

package ioctl

import (
	"errors"

	"golang.org/x/sys/unix"
)

// IOC returns an ioctl(2) request value, calculated from the specific ioctl
// call properties: parameter in/out direction, type of ioctl, command number,
// and finally parameter size.
func IOC(dir, ioctype, nr, size uint) uint {
	return (dir << IOC_DIRSHIFT) | (ioctype << IOC_TYPESHIFT) | (nr << IOC_NRSHIFT) | (size << IOC_SIZESHIFT)
}

// IO returns an ioctl(2) request value for a request that doesn't have any
// additional request parameter.
func IO(ioctype, nr uint) uint {
	return IOC(IOC_NONE, ioctype, nr, 0)
}

// IOR returns an ioctl(2) request value for a request that has an additional
// request parameter that the userland wants to read and the kernel is writing.
func IOR(ioctype, nr, size uint) uint {
	return IOC(IOC_READ, ioctype, nr, size)
}

// IOW returns an ioctl(2) request value for a request that has an additional
// request parameter that the userland wants to read and the kernel is writing.
func IOW(ioctype, nr, size uint) uint {
	return IOC(IOC_WRITE, ioctype, nr, size)
}

// IORW returns an ioctl(2) request value for a request that has an additional
// request parameter that the userland first wants to write, the kernel then
// reads and updates it, and the userland finally wants to read it afterwards.
func IORW(ioctype, nr, size uint) uint {
	return IOC(IOC_READ|IOC_WRITE, ioctype, nr, size)
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
