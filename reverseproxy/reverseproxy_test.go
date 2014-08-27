package reverseproxy_test

import (
	. "github.com/datianshi/goproxy-cf/reverseproxy"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Reverseproxy", func() {
	It("join string correctly", func() {

		Ω(SingleJoiningSlash("/app1", "/app2")).To(Equal("/app1/app2"))
	})
})
