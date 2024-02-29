// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package registry

import (
	"context"
	"github.com/OnenessLabs/genie/dkg/packages/cryptolib"
	"github.com/OnenessLabs/genie/dkg/packages/peering"
	"github.com/OnenessLabs/genie/dkg/packages/tcrypto"

	iotago "github.com/iotaledger/iota.go/v3"
)

type NodeIdentityProvider interface {
	NodeIdentity() *cryptolib.KeyPair
	NodePublicKey() *cryptolib.PublicKey
}

type DKShareRegistryProvider interface {
	SaveDKShare(dkShare tcrypto.DKShare) error
	LoadDKShare(sharedAddress iotago.Address) (tcrypto.DKShare, error)
}

type TrustedPeersRegistryProvider interface {
	IsTrustedPeer(pubKey *cryptolib.PublicKey) error
	TrustPeer(name string, pubKey *cryptolib.PublicKey, accountURL string) (*peering.TrustedPeer, error)
	DistrustPeer(pubKey *cryptolib.PublicKey) (*peering.TrustedPeer, error)
	TrustedPeers() ([]*peering.TrustedPeer, error)
	TrustedPeersByPubKeyOrName(pubKeysOrNames []string) ([]*peering.TrustedPeer, error)
	TrustedPeersListener(callback func([]*peering.TrustedPeer)) context.CancelFunc
}
