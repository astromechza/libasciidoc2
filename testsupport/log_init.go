package testsupport

import (
	"github.com/davecgh/go-spew/spew"
)

func init() {
	// also, configuration for spew (when dumping structures to compare results)
	spew.Config.DisableCapacities = true
	spew.Config.DisablePointerAddresses = true
	spew.Config.DisablePointerMethods = true
	spew.Config.DisableUnexported = true
}
