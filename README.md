# `ioctl`

[![PkgGoDev](https://pkg.go.dev/badge/github.com/thediveo/ioctl)](https://pkg.go.dev/github.com/thediveo/ioctl)
[![GitHub](https://img.shields.io/github/license/thediveo/ioctl)](https://img.shields.io/github/license/thediveo/ioctl)
![build and test](https://github.com/thediveo/ioctl/workflows/build%20and%20test/badge.svg?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/thediveo/ioctl)](https://goreportcard.com/report/github.com/thediveo/ioctl)
![Coverage](https://img.shields.io/badge/Coverage-100.0%25-brightgreen)

A tiny package to help with dealing with constructing Linux
[ioctl(2)](https://man7.org/linux/man-pages/man2/ioctl.2.html) request values
that are not already included in the
[sys/unix](https://pkg.go.dev/golang.org/x/sys/unix) standard package.

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
var NS_GET_USERNS = _IO(NSIO, 0x1)

func main() {
  fd, err := ioctl.RetFd(nsfd, NS_GET_USERNS)
}
```

## Make Targets

- `make`: lists all targets.
- `make coverage`: runs all tests with coverage and then **updates the coverage
  badge in `README.md`**.
- `make pkgsite`: installs [`x/pkgsite`](https://golang.org/x/pkgsite/cmd/pkgsite), as
  well as the [`browser-sync`](https://www.npmjs.com/package/browser-sync) and
  [`nodemon`](https://www.npmjs.com/package/nodemon) npm packages first, if not
  already done so. Then runs the `pkgsite` and hot reloads it whenever the
  documentation changes.
- `make report`: installs
  [`@gojp/goreportcard`](https://github.com/gojp/goreportcard) if not yet done
  so and then runs it on the code base.
- `make test`: runs **all** tests, always.
- `make vuln`: installs
  [`x/vuln/cmd/govulncheck`](https://golang.org/x/vuln/cmd/govulncheck) and then
  runs it.

## Copyright and License

Copyright 2023 Harald Albrecht, licensed under the Apache License, Version 2.0.
