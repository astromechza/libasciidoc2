package testsupport_test

import (
	"fmt"
	"log/slog"

	"github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("console assertions", func() {

	content := `{"level":"DEBUG","msg":"processing foo.adoc: ----\ninclude::../../test/includes/unknown.adoc[leveloffset=+1]\n----"}
{"level":"DEBUG","line":1,"msg":"initializing a new DraftDocument with 1 block element(s)"}
{"level":"DEBUG","line":1,"msg":"parsing '../../test/includes/unknown.adoc'..."}
{"error":"open unknown.adoc: no such file or directory","level":"ERROR","start_offset":0,"end_offset":60,"msg":"failed to include '../../test/includes/unknown.adoc'"}
{"level":"DEBUG","line":1,"msg":"restoring current working dir to: github.com/bytesparadise/libasciidoc/pkg/parser"}`

	var out testsupport.ConsoleOutput
	BeforeEach(func() {
		out = *testsupport.NewConsoleOutput()
		_, err := out.Write([]byte(content))
		Expect(err).NotTo(HaveOccurred())
	})

	Context("with message, offset and level", func() {

		It("should find expected level/message", func() {
			// given
			matcher := testsupport.ContainJSONLogWithOffset(slog.LevelError, 0, 60, "failed to include '../../test/includes/unknown.adoc'")
			// when
			result, err := matcher.Match(out)
			// then
			Expect(err).ToNot(HaveOccurred())
			Expect(result).To(BeTrue())
		})

		It("should not find expected level/message with wrong level", func() {
			// given an incorrect log level
			matcher := testsupport.ContainJSONLogWithOffset(slog.LevelWarn, 0, 60, "failed to include '../../test/includes/unknown.adoc'")
			// when
			result, err := matcher.Match(out)
			// then
			Expect(err).ToNot(HaveOccurred())
			Expect(result).To(BeFalse())
		})

		It("should not find expected level/message with wrong start_offset", func() {
			// given an incorrect start_offset value
			matcher := testsupport.ContainJSONLogWithOffset(slog.LevelWarn, 10, 60, "failed to include '../../test/includes/unknown.adoc'")
			// when
			result, err := matcher.Match(out)
			// then
			Expect(err).ToNot(HaveOccurred())
			Expect(result).To(BeFalse())
		})

		It("should not find expected level/message with wrong end_offset", func() {
			// given an incorrect end_offset value
			matcher := testsupport.ContainJSONLogWithOffset(slog.LevelWarn, 0, 10, "failed to include '../../test/includes/unknown.adoc'")
			// when
			result, err := matcher.Match(out)
			// then
			Expect(err).ToNot(HaveOccurred())
			Expect(result).To(BeFalse())
		})

		It("should not find expected level/message with wrong msg", func() {
			// given an incorrect msg value
			matcher := testsupport.ContainJSONLogWithOffset(slog.LevelWarn, 0, 10, "something else")
			// when
			result, err := matcher.Match(out)
			// then
			Expect(err).ToNot(HaveOccurred())
			Expect(result).To(BeFalse())
			// also verify the messages
			Expect(matcher.FailureMessage(out)).To(Equal(fmt.Sprintf(`expected console to contain log {"level": "%s", "start_offset":%d, "end_offset":%d, "msg":"%s"}`, slog.LevelWarn, 0, 10, "something else")))
			Expect(matcher.NegatedFailureMessage(out)).To(Equal(fmt.Sprintf(`expected console not to contain log {"level": "%s", "start_offset":%d, "end_offset":%d, "msg":"%s"}`, slog.LevelWarn, 0, 10, "something else")))
		})
	})

	Context("with message and level", func() {

		It("should find expected level/message", func() {
			// given
			matcher := testsupport.ContainJSONLog(slog.LevelError, "failed to include '../../test/includes/unknown.adoc'")
			// when
			result, err := matcher.Match(out)
			// then
			Expect(err).ToNot(HaveOccurred())
			Expect(result).To(BeTrue())
		})

		It("should not find expected level/message with wrong level", func() {
			// given an incorrect log level
			matcher := testsupport.ContainJSONLog(slog.LevelWarn, "failed to include '../../test/includes/unknown.adoc'")
			// when
			result, err := matcher.Match(out)
			// then
			Expect(err).ToNot(HaveOccurred())
			Expect(result).To(BeFalse())
		})

		It("should not find expected level/message with wrong msg", func() {
			// given an incorrect msg value
			matcher := testsupport.ContainJSONLog(slog.LevelWarn, "something else")
			// when
			result, err := matcher.Match(out)
			// then
			Expect(err).ToNot(HaveOccurred())
			Expect(result).To(BeFalse())
			// also verify the messages
			Expect(matcher.FailureMessage(out)).To(Equal(fmt.Sprintf(`expected console to contain log {"level": "%s", "msg":"%s"}`, slog.LevelWarn, "something else")))
			Expect(matcher.NegatedFailureMessage(out)).To(Equal(fmt.Sprintf(`expected console not to contain log {"level": "%s", "msg":"%s"}`, slog.LevelWarn, "something else")))
		})
	})

	Context("with level only", func() {

		It("should find with single given level", func() {
			// given
			matcher := testsupport.ContainAnyMessageWithLevels(slog.LevelError)
			// when
			result, err := matcher.Match(out)
			// then
			Expect(err).ToNot(HaveOccurred())
			Expect(result).To(BeTrue())
		})

		It("should find with multiple given levels", func() {
			// given
			matcher := testsupport.ContainAnyMessageWithLevels(slog.LevelError, slog.LevelWarn)
			// when
			result, err := matcher.Match(out)
			// then
			Expect(err).ToNot(HaveOccurred())
			Expect(result).To(BeTrue())
		})

		It("should not find with single given level", func() {
			// given
			matcher := testsupport.ContainAnyMessageWithLevels(slog.LevelWarn)
			// when
			result, err := matcher.Match(out)
			// then
			Expect(err).ToNot(HaveOccurred())
			Expect(result).To(BeFalse())
			// also verify the messages
		})

		It("should not find with multiple given levels", func() {
			// given
			matcher := testsupport.ContainAnyMessageWithLevels(slog.LevelWarn, slog.LevelInfo)
			// when
			result, err := matcher.Match(out)
			// then
			Expect(err).ToNot(HaveOccurred())
			Expect(result).To(BeFalse())
			// also verify the messages

		})

		It("should return error when invalid type is input", func() {
			// given
			matcher := testsupport.ContainAnyMessageWithLevels(slog.LevelError) // unknown message
			// when
			result, err := matcher.Match(1) // not a reader
			// then
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("ContainAnyMessageWithLevels matcher expects an io.Reader (actual: int)"))
			Expect(result).To(BeFalse())
		})
	})

})
