package testsupport

import (
	"github.com/davecgh/go-spew/spew"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.TextFormatter{
		DisableQuote: true, // see https://github.com/sirupsen/logrus/issues/608#issuecomment-745137306
	})

	// also, configuration for spew (when dumping structures to compare results)
	spew.Config.DisableCapacities = true
	spew.Config.DisablePointerAddresses = true
	spew.Config.DisablePointerMethods = true
	spew.Config.DisableUnexported = true
}
