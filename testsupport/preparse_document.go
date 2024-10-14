package testsupport

import (
	"github.com/bytesparadise/libasciidoc/pkg/log"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/bytesparadise/libasciidoc/pkg/parser"

	"github.com/pkg/errors"
)

func PreparseDocument(source string, options ...interface{}) (string, error) {
	settings := []configuration.Setting{
		configuration.WithFilename("test.adoc"),
		configuration.WithAttribute("basebackend-html", true), // TODO: still needed?
	}
	opts := []parser.Option{}
	for _, o := range options {
		switch o := o.(type) {
		case configuration.Setting:
			settings = append(settings, o)
		case parser.Option:
			opts = append(opts, o)
		default:
			return "", errors.Errorf("unexpected type of option: '%T'", o)
		}
	}
	result, err := parser.Preprocess(strings.NewReader(source), configuration.NewConfiguration(settings...), opts...)
	if log.DebugEnabled() && err == nil {
		log.Debugf("preparsed document:\n%s", result)
	}
	return result, err

}
