//go:build postgres

package beacon

import (
	"context"
	"fmt"
	"testing"

	"github.com/OnenessLabs/genie/newrand/chain"
	"github.com/OnenessLabs/genie/newrand/chain/postgresdb/pgdb"
	"github.com/OnenessLabs/genie/newrand/log"
	"github.com/OnenessLabs/genie/newrand/test"
	context2 "github.com/OnenessLabs/genie/newrand/test/context"
)

var c *test.Container

func TestMain(m *testing.M) {
	var err error
	c, err = test.StartPGDB()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer test.StopPGDB(c)

	m.Run()
}

func createStore(t *testing.T, l log.Logger, _ *BeaconTest, _ int) (chain.Store, error) {
	dbName := test.ComputeDBName()
	_, dbConn := test.NewUnit(t, c, dbName)

	ctx, _, _ := context2.PrevSignatureMattersOnContext(t, context.Background())

	return pgdb.NewStore(ctx, l, dbConn, dbName)
}
