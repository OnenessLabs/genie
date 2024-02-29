// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package peering_test

import (
	"github.com/OnenessLabs/genie/dkg/packages/peering"
	"github.com/OnenessLabs/genie/dkg/packages/util/rwutil"
	"testing"
)

func TestPeeringIDSerialization(t *testing.T) {
	peeringID := peering.RandomPeeringID()

	rwutil.ReadWriteTest(t, &peeringID, new(peering.PeeringID))
}
