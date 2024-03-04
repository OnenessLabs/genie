//go:build !postgres && !memdb

package core

import (
	"testing"

	"github.com/OnenessLabs/genie/newrand/chain"
)

func WithTestDB(_ *testing.T, _ string) []ConfigOption {
	return []ConfigOption{
		WithDBStorageEngine(chain.BoltDB),
	}
}
