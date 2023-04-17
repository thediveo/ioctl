/*
Package ioctl is a thin convenience wrapper for those ioctl request numbers that are not defined
as plain numbers but instead as macros in the Linux kernel header files.

# Example

The request values for [ioctl_ns(2)] operations are defined using C macros in
[include/uapi/linux/nsfs.h] in the Linux kernel header sources as follows:

	#define NSIO	0xb7

	// Returns a file descriptor that refers to an owning user namespace
	#define NS_GET_USERNS		_IO(NSIO, 0x1)

These definitions can now be applied to Go code as follows:

	import "github.com/thediveo/ioctl"

	var NS_GET_USERNS = _IO(NSIO, 0x1)

	fd, err := ioctl.RetFd(nsfd, NS_GET_USERNS)

[ioctl_ns(2)]: https://man7.org/linux/man-pages/man2/ioctl_ns.2.html
[include/uapi/linux/nsfs.h]: https://elixir.bootlin.com/linux/v6.2.11/source/include/uapi/linux/nsfs.h#L10
*/
package ioctl
