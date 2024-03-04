//go:build memdb

package beacon

import (
	"testing"

	"github.com/OnenessLabs/genie/newrand/chain"
	"github.com/OnenessLabs/genie/newrand/chain/memdb"
	"github.com/OnenessLabs/genie/newrand/log"
)

func createStore(_ *testing.T, _ log.Logger, _ *BeaconTest, _ int) (chain.Store, error) {
	return memdb.NewStore(10), nil
}
