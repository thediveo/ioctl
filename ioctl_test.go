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
	"os"

	"golang.org/x/sys/unix"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/thediveo/success"
)

func nsIno[R ~int | ~string](ref R) uint64 {
	GinkgoHelper()

	var nsStat unix.Stat_t
	switch ref := any(ref).(type) {
	case int:
		Expect(unix.Fstat(ref, &nsStat)).To(Succeed())
	case string:
		Expect(unix.Stat(ref, &nsStat)).To(Succeed())
	}
	return nsStat.Ino
}

var _ = Describe("ioctl requests", func() {

	Context("request value calculation", func() {

		const (
			ioctype = 0x12
			iocnr   = 0x42
			iosize  = 0x234 // must be <= 2**13-1 (some archs)
		)

		// ah, the beauty of full coverage...
		ioc := func(dir, ioctype, nr, size uint) uint {
			return (dir << IOC_DIRSHIFT) | (ioctype << IOC_TYPESHIFT) | (nr << IOC_NRSHIFT) | (size << IOC_SIZESHIFT)
		}

		DescribeTable("IO*",
			func(actual, expected uint) {
				Expect(actual).To(Equal(expected), "0x%08x 0x%08x", actual, expected)
			},
			Entry("IO no R, no W",
				IO(ioctype, iocnr),
				ioc(IOC_NONE, ioctype, iocnr, 0)),
			Entry("IOR",
				IOR(ioctype, iocnr, iosize),
				ioc(IOC_READ, ioctype, iocnr, iosize)),
			Entry("IOW",
				IOW(ioctype, iocnr, iosize),
				ioc(IOC_WRITE, ioctype, iocnr, iosize)),
			Entry("IORW",
				IORW(ioctype, iocnr, iosize),
				ioc(IOC_WRITE|IOC_READ, ioctype, iocnr, iosize)),
		)

	})

	Context("RetFd", func() {

		It("returns an invalid -1 fd when in error", func() {
			fd, err := RetFd(0, 0)
			Expect(err).To(HaveOccurred())
			Expect(fd).To(Equal(int(-1)))
		})

		It("calculates _IO and executes RetFd correctly", func() {
			const NSIO = 0xb7
			var NS_GET_USERNS = IO(NSIO, 0x1)

			netnsf := Successful(os.Open("/proc/self/ns/net"))
			defer netnsf.Close()
			usernsfd, err := RetFd(int(netnsf.Fd()), NS_GET_USERNS)
			Expect(err).NotTo(HaveOccurred())
			defer unix.Close(usernsfd)
			Expect(nsIno(usernsfd)).To(Equal(nsIno("/proc/self/ns/user")))
		})

	})

})
