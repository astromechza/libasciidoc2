package testsupport_test

import (
	"time"

	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" // nolint:golint
	. "github.com/onsi/gomega" // nolint:golintt
)

var _ = Describe("document metadata", func() {

	lastUpdated := time.Now()
	expected := types.Metadata{
		LastUpdated: lastUpdated.Format(configuration.LastUpdatedFormat),
		Title:       "Title",
		TableOfContents: &types.TableOfContents{
			MaxDepth: 2,
			Sections: []*types.ToCSection{
				{
					ID:    "_section_1",
					Level: 1,
					Title: "Section 1",
				},
			},
		},
	}

	It("should match", func() {
		// given
		actual := `= Title
:toc:
		
== Section 1`
		// when
		result, err := testsupport.DocumentMetadata(actual, lastUpdated)
		// then
		Expect(err).ToNot(HaveOccurred())
		Expect(result).To(Equal(expected))
	})

	It("should not match", func() {
		// given
		actual := `foo`
		// when
		result, err := testsupport.DocumentMetadata(actual, lastUpdated)
		// then
		Expect(err).ToNot(HaveOccurred())
		Expect(result).NotTo(Equal(expected))
	})

})
