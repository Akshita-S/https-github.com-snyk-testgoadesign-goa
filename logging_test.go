package goa_test

import (
	"bytes"
	"log"

	"github.com/goadesign/goa"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"golang.org/x/net/context"
)

var _ = Describe("Info", func() {
	Context("with a nil Log", func() {
		It("doesn't log and doesn't crash", func() {
			Ω(func() { goa.Info(context.Background(), "foo", "bar") }).ShouldNot(Panic())
		})
	})
})

var _ = Describe("Error", func() {
	Context("with a nil Log", func() {
		It("doesn't log and doesn't crash", func() {
			Ω(func() { goa.Error(context.Background(), "foo", "bar") }).ShouldNot(Panic())
		})
	})
})

var _ = Describe("StdLogger", func() {
	Context("with a valid Log", func() {
		var logger goa.Logger
		const msg = "message"
		data := []interface{}{"data", "foo"}

		var out bytes.Buffer

		BeforeEach(func() {
			stdlogger := log.New(&out, "", log.LstdFlags)
			logger = goa.NewStdLogger(stdlogger)
		})

		It("Info logs", func() {
			logger.Info(msg, data...)
			Ω(out.String()).Should(ContainSubstring(msg + " data=foo"))
		})

		It("Error logs", func() {
			logger.Error(msg, data...)
			Ω(out.String()).Should(ContainSubstring(msg + " data=foo"))
		})
	})
})
