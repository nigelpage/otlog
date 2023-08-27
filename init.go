package otlog

import (
	"runtime/debug"
)

var moduleVersion string

// Actions in here are only ever performed once
func init() {
	// Retirieve module version number
	bi, ok := debug.ReadBuildInfo();

	if !ok {
		panic("unable to retrieve module version number!")
	}
	moduleVersion = bi.Main.Version
}