module github.com/bytesparadise/libasciidoc

go 1.23

require (
	// Chroma is used for code block highlighting
	github.com/alecthomas/chroma/v2 v2.3.0
	// Spew is a pretty printer.. we don't really need this tbh
	github.com/davecgh/go-spew v1.1.1
	// go-cmp is used for generating pretty diffs in test results
	github.com/google/go-cmp v0.5.9

	// ==== BUILD DEPENENCIES

	// pigeon is a parser generator - its only used at build time but is needed here for versioning
	github.com/mna/pigeon v1.1.0
	// ginkgo is a testing framework
	github.com/onsi/ginkgo/v2 v2.7.0
	// gomega is another testing framework
	github.com/onsi/gomega v1.24.2
	// pkg errors is used for wrapping error support - this could be replaced with standard library
	github.com/pkg/errors v0.9.1

	// ==== TEST DEPENDENCIES

	// logrus is for structured/leveled logging - this can be replaced now with modern GO
	github.com/sirupsen/logrus v1.7.0
	// Testify is a very common and trusted testing utility
	github.com/stretchr/testify v1.8.0
)

require (
	github.com/dlclark/regexp2 v1.4.0 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-task/slim-sprig v0.0.0-20210107165309-348f09dbbbc0 // indirect
	github.com/google/pprof v0.0.0-20210407192527-94a9f03dee38 // indirect
	github.com/kr/pretty v0.1.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/mod v0.8.0 // indirect
	golang.org/x/net v0.17.0 // indirect
	golang.org/x/sys v0.13.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	golang.org/x/tools v0.6.0 // indirect
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

// include support for disabling unexported fields
// TODO: still needed?
replace github.com/davecgh/go-spew => github.com/flw-cn/go-spew v1.1.2-0.20200624141737-10fccbfd0b23
