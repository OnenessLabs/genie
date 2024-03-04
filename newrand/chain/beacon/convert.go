package beacon

import (
	"github.com/OnenessLabs/genie/newrand/chain"
	"github.com/OnenessLabs/genie/newrand/protobuf/common"
	proto "github.com/OnenessLabs/genie/newrand/protobuf/drand"
)

func beaconToProto(b *chain.Beacon, beaconID string) *proto.BeaconPacket {
	return &proto.BeaconPacket{
		PreviousSignature: b.PreviousSig,
		Round:             b.Round,
		Signature:         b.Signature,
		Metadata:          &common.Metadata{BeaconID: beaconID},
	}
}

func protoToBeacon(p *proto.BeaconPacket) *chain.Beacon {
	return &chain.Beacon{
		Round:       p.GetRound(),
		Signature:   p.GetSignature(),
		PreviousSig: p.GetPreviousSignature(),
	}
}
