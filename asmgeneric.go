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

/*
This is ugly ioctl command (request) enconding stuff. Requests are 32 bits in
total, with the command in the lower 16 bits and the size of the parameter
structure taking up the lower 14 bits of the request's upper 16 bits.

ATTENTION: the following definitions follow the kernel's [ioctl.h]. They hold
“only” for the “asm-generic” platforms, such as x86, arm, and others. Currently
the only platforms having a different ioctl request field mapping are: alpha,
mips, powerpc, and sparc. And alpha isn't even an architecture supported by Go,
go figure.

We keep the original C identifiers on purpose (but have to remove the leading
“_”) and otherwise don't care about patronizing linters.

[ioctl.h]: https://elixir.bootlin.com/linux/v6.2.11/source/include/uapi/asm-generic/ioctl.h
*/
const (
	IOC_NRBITS   = 8
	IOC_TYPEBITS = 8
	IOC_SIZEBITS = 14

	IOC_NRSHIFT   = 0
	IOC_TYPESHIFT = IOC_NRSHIFT + IOC_NRBITS
	IOC_SIZESHIFT = IOC_TYPESHIFT + IOC_TYPEBITS
	IOC_DIRSHIFT  = IOC_SIZESHIFT + IOC_SIZEBITS

	IOC_NONE  = uint(0)
	IOC_WRITE = uint(1)
	IOC_READ  = uint(2)
)
