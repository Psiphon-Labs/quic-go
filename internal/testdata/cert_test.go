package testdata

import (
	"io"

	tls "github.com/Psiphon-Labs/psiphon-tls"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("certificates", func() {
	It("returns certificates", func() {
		ln, err := tls.Listen("tcp", "localhost:4433", GetTLSConfig())
		Expect(err).ToNot(HaveOccurred())

		go func() {
			defer GinkgoRecover()
			conn, err := ln.Accept()
			Expect(err).ToNot(HaveOccurred())
			defer conn.Close()
			_, err = conn.Write([]byte("foobar"))
			Expect(err).ToNot(HaveOccurred())
		}()

		conn, err := tls.Dial("tcp", "localhost:4433", &tls.Config{RootCAs: GetRootCA()})
		Expect(err).ToNot(HaveOccurred())
		data, err := io.ReadAll(conn)
		Expect(err).ToNot(HaveOccurred())
		Expect(string(data)).To(Equal("foobar"))
	})
})
