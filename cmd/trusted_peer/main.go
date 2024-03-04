package main

import (
	"github.com/OnenessLabs/genie/dkg"
	"github.com/OnenessLabs/genie/dkg/packages/cryptolib"
	"github.com/OnenessLabs/genie/dkg/packages/registry"
)

func main() {
	dkg.ParamsRegistries.TrustedPeers.FilePath = "waspdb/trusted_peers.json"
	trustedPeersRegistryProvider, err := registry.NewTrustedPeersRegistryImpl(dkg.ParamsRegistries.TrustedPeers.FilePath)

	if err != nil {

	}
	if trustedPeersRegistryProvider != nil {
		nodePubKey, err := cryptolib.PublicKeyFromString("0xc8b1d3a85cee9949bde989144da8391269d2632866794a6eab11d0b7777c07cc")
		if err != nil {
			panic(err)
		}
		trustedPeersRegistryProvider.TrustPeer("node4001", nodePubKey, "0.0.0.0:4001")
		nodePubKey, err = cryptolib.PublicKeyFromString("0x3b5059ef01bb33b471e7ba24e8b7a4880eb545c274d65ac3e5019dd0e835e067")
		if err != nil {
			panic(err)
		}
		trustedPeersRegistryProvider.TrustPeer("node4000", nodePubKey, "0.0.0.0:4000")
		nodePubKey, err = cryptolib.PublicKeyFromString("0xe40f69d5a97e26b38296ec55485767498b2d2565eb814720d09e1eb31a765f7e")
		if err != nil {
			panic(err)
		}
		trustedPeersRegistryProvider.TrustPeer("node4002", nodePubKey, "0.0.0.0:4002")
	}
}
