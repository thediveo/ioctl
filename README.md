# `ioctl`

[![PkgGoDev](https://pkg.go.dev/badge/github.com/thediveo/ioctl)](https://pkg.go.dev/github.com/thediveo/ioctl)
[![GitHub](https://img.shields.io/github/license/thediveo/ioctl)](https://img.shields.io/github/license/thediveo/ioctl)
![build and test](https://github.com/thediveo/ioctl/actions/workflows/buildandtest.yaml/badge.svg?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/thediveo/ioctl)](https://goreportcard.com/report/github.com/thediveo/ioctl)
![Coverage](https://img.shields.io/badge/Coverage-100.0%25-brightgreen)

A tiny package to help with dealing with constructing Linux
[ioctl(2)](https://man7.org/linux/man-pages/man2/ioctl.2.html) request values
that are not already included in the
[sys/unix](https://pkg.go.dev/golang.org/x/sys/unix) standard package.

For devcontainer instructions, please see the [section "DevContainer"
below](#devcontainer).

## Usage

A good example are
[ioctl_ns(2)](https://man7.org/linux/man-pages/man2/ioctl_ns.2.html) operations
on Linux namespaces. For instance, this ioctl operation request value is defined
in
[include/uapi/linux/nsfs.h](https://elixir.bootlin.com/linux/v6.2.11/source/include/uapi/linux/nsfs.h#L10)
in the Linux kernel C headers as follows:

```c
#define NSIO	0xb7

/* Returns a file descriptor that refers to an owning user namespace */
#define NS_GET_USERNS		_IO(NSIO, 0x1)
```

These definitions can now be applied to Go code as follows:

```go
import "github.com/thediveo/ioctl"

const NSIO = 0xb7
var NS_GET_USERNS = ioctl.IO(NSIO, 0x1)

func main() {
  fd, err := ioctl.RetFd(nsfd, NS_GET_USERNS)
}
```

## DevContainer

> [!CAUTION]
>
> Do **not** use VSCode's "~~Dev Containers: Clone Repository in Container
> Volume~~" command, as it is utterly broken by design, ignoring
> `.devcontainer/devcontainer.json`.

1. `git clone https://github.com/thediveo/enumflag`
2. in VSCode: Ctrl+Shift+P, "Dev Containers: Open Workspace in Container..."
3. select `enumflag.code-workspace` and off you go...

## Supported Go Versions

`native` supports versions of Go that are noted by the [Go release
policy](https://golang.org/doc/devel/release.html#policy), that is, major
versions _N_ and _N_-1 (where _N_ is the current major version).

## Copyright and License

Copyright 2023, 2025 Harald Albrecht, licensed under the Apache License, Version
2.0.
