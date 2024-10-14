package testsupport

import (
	"fmt"
	"github.com/bytesparadise/libasciidoc/pkg/log"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/davecgh/go-spew/spew"
	"github.com/google/go-cmp/cmp"
	gomegatypes "github.com/onsi/gomega/types"
	"github.com/pkg/errors"
)

// MatchDocumentFragments a custom matcher to verify that a document matches the given expectation
// Similar to the standard `Equal` matcher, but display a diff when the values don't match
func MatchDocumentFragments(expected []types.DocumentFragment) gomegatypes.GomegaMatcher {
	return &documentFragmentsMatcher{
		expected: expected,
	}
}

type documentFragmentsMatcher struct {
	expected []types.DocumentFragment
	diffs    string
}

func (m *documentFragmentsMatcher) Match(actual interface{}) (success bool, err error) {
	if _, ok := actual.([]types.DocumentFragment); !ok {
		return false, errors.Errorf("MatchDocumentFragments matcher expects a '[]types.DocumentFragment' (actual: %T)", actual)
	}
	if diff := cmp.Diff(m.expected, actual, opts...); diff != "" {
		if log.DebugEnabled() {
			log.Debugf("actual document fragments:\n%s", spew.Sdump(actual))
			log.Debugf("expected document fragments:\n%s", spew.Sdump(m.expected))
		}
		m.diffs = diff
		return false, nil
	}
	return true, nil
}

func (m *documentFragmentsMatcher) FailureMessage(_ interface{}) (message string) {
	return fmt.Sprintf("expected document fragments to match:\n%s", m.diffs)
}

func (m *documentFragmentsMatcher) NegatedFailureMessage(_ interface{}) (message string) {
	return fmt.Sprintf("expected document fragments not to match:\n%s", m.diffs)
}
