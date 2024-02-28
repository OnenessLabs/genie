package common

import (
	"fmt"
)

// Must be manually updated!
// Before releasing: Verify the version number and set Meta to ""
// After releasing: Increase the Patch number and set Meta to "pre/rcN/..."
var version = Version{
	Major: 0,
	Minor: 0,
	Patch: 1,
	Meta:  "pre",
}

// Set via -ldflags. Example:
//
//	go install -ldflags "-X common.BUILDDATE=`date -u +%d/%m/%Y@%H:%M:%S` -X common.GITCOMMIT=`git rev-parse HEAD`
//
// See the Makefile and the Dockerfile in the root directory of the repo
var (
	COMMIT    = ""
	BUILDDATE = ""
)

func GetAppVersion() Version {
	return version
}

type Version struct {
	Major uint32
	Minor uint32
	Patch uint32
	Meta  string
}

func (v Version) String() string {
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
}

// WithMeta holds the textual version string including the metadata.
func (v Version) WithMeta() string {
	pre := ""
	if v.Meta != "" {
		pre = "-"
	}
	return fmt.Sprintf("%d.%d.%d%s%s", v.Major, v.Minor, v.Patch, pre, v.Meta)
}

func (v Version) WithCommit() string {
	vsn := v.WithMeta()
	if len(COMMIT) >= 8 {
		vsn += "-" + COMMIT[:8]
	}
	if (v.Meta != "") && (BUILDDATE != "") {
		vsn += "-" + BUILDDATE
	}
	return vsn
}
