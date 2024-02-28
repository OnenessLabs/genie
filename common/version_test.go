package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	version123 = Version{
		Major: 1,
		Minor: 2,
		Patch: 3,
		Meta:  "",
	}

	version123pre = Version{
		Major: 1,
		Minor: 2,
		Patch: 3,
		Meta:  "pre",
	}
)

func TestVersionStringNoPre(t *testing.T) {
	assert.Equal(t, "1.2.3", version123.String(), "Incorrect version string.")
}

func TestVersionStringPre(t *testing.T) {
	actual := version123pre.String()
	expected := "1.2.3-pre"

	if actual != expected {
		t.Fatalf("Incorrect version string. Actual: %s, expected: %s", actual, expected)
	}
}
