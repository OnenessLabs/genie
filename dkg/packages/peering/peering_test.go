package peering_test

import (
	"github.com/OnenessLabs/genie/dkg/packages/peering"
	"github.com/OnenessLabs/genie/dkg/packages/util/rwutil"
	"testing"
)

func TestPeerMessageSerialization(t *testing.T) {
	msg := &peering.PeerMessageNet{
		PeerMessageData: peering.NewPeerMessageData(
			peering.RandomPeeringID(),
			byte(10),
			peering.FirstUserMsgCode+17,
			[]byte{1, 2, 3, 4, 5}),
	}
	rwutil.BytesTest(t, msg, peering.PeerMessageNetFromBytes)
}
