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

	It("returns an invalid -1 fd when in error", func() {
		fd, err := RetFd(0, 0)
		Expect(err).To(HaveOccurred())
		Expect(fd).To(Equal(int(-1)))
	})

	It("calculates _IO correctly", func() {
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
