//go:build !postgres && !memdb

package beacon

import (
	"context"
	"testing"

	"github.com/OnenessLabs/genie/newrand/chain"
	"github.com/OnenessLabs/genie/newrand/chain/boltdb"
	"github.com/OnenessLabs/genie/newrand/log"
	context2 "github.com/OnenessLabs/genie/newrand/test/context"
)

func createStore(t *testing.T, l log.Logger, b *BeaconTest, idx int) (chain.Store, error) {
	ctx, _, _ := context2.PrevSignatureMattersOnContext(t, context.Background())
	return boltdb.NewBoltStore(ctx, l, b.paths[idx], nil)
}
