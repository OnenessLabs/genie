package dkg

import (
	"context"
	"fmt"
	"github.com/OnenessLabs/genie/config"
	"github.com/OnenessLabs/genie/dkg/packages/cryptolib"
	"github.com/OnenessLabs/genie/dkg/packages/metrics"
	"github.com/OnenessLabs/genie/dkg/packages/peering/lpp"
	"github.com/OnenessLabs/genie/dkg/packages/registry"
	hivep2p "github.com/iotaledger/hive.go/crypto/p2p"
	"github.com/iotaledger/hive.go/logger"
	"github.com/iotaledger/hive.go/runtime/ioutils"
	iotago "github.com/iotaledger/iota.go/v3"
	"github.com/iotaledger/iota.go/v3/tpkg"
	"go.uber.org/fx"
	"path"
)

type ParametersPeering struct {
	PeeringURL string `default:"0.0.0.0:4000" usage:"node host address as it is recognized by other peers"`
	Port       int    `default:"4000" usage:"port for Wasp committee connection/peering"`
}

var ParamsPeering = &ParametersPeering{}

type ParametersRegistries struct {
	Chains struct {
		FilePath string `default:"waspdb/chains/chain_registry.json" usage:"the path to the chain registry file"`
	}
	DKShares struct {
		Path string `default:"waspdb/dkshares" usage:"the path to the distributed key shares registries folder"`
	}
	TrustedPeers struct {
		FilePath string `default:"waspdb/trusted_peers.json" usage:"the path to the trusted peers registry file"`
	}
	ConsensusState struct {
		Path string `default:"waspdb/chains/consensus" usage:"the path to the consensus state registries folder"`
	}
}

// ParametersP2P contains the definition of the parameters used by p2p.
type ParametersP2P struct {
	Identity struct {
		PrivateKey string `default:"" usage:"private key used to derive the node identity (optional)"`
		FilePath   string `default:"waspdb/identity/identity.key" usage:"the path to the node identity PEM file"`
	}

	Database struct {
		// Defines the path to the p2p database.
		Path string `default:"waspdb/p2pstore" usage:"the path to the p2p database"`
	} `name:"db"`
}

var (
	ParamsRegistries = &ParametersRegistries{}
	ParamsP2P        = &ParametersP2P{}
)

func NodeIdentityRegistry(lc fx.Lifecycle, cfg *config.Config, cLogger *logger.Logger) *registry.NodeIdentity {
	ParamsP2P.Database.Path = "waspdb/p2pstore"
	ParamsP2P.Identity.FilePath = "waspdb/identity/identity.key"
	if err := ioutils.CreateDirectory(ParamsP2P.Database.Path, 0o700); err != nil {
		cLogger.Panicf("could not create peer store database dir '%s': %w", ParamsP2P.Database.Path, err)
	}

	// make sure nobody copies around the peer store since it contains the private key of the node
	cLogger.Infof(`WARNING: never share your "%s" or "%s" folder as both contain your node's private key!`, ParamsP2P.Database.Path, path.Dir(ParamsP2P.Identity.FilePath))

	// load up the previously generated identity or create a new one
	privKey, newlyCreated, err := hivep2p.LoadOrCreateIdentityPrivateKey(ParamsP2P.Identity.FilePath, ParamsP2P.Identity.PrivateKey)
	if err != nil {
		cLogger.Panic(err)
	}

	if newlyCreated {
		cLogger.Infof(`stored new private key for peer identity under "%s"`, ParamsP2P.Identity.FilePath)
	} else {
		cLogger.Infof(`loaded existing private key for peer identity from "%s"`, ParamsP2P.Identity.FilePath)
	}

	privKeyBytes, err := privKey.Raw()
	if err != nil {
		cLogger.Panicf("unable to convert private key for peer identity: %s", err)
	}

	waspPrivKey, err := cryptolib.PrivateKeyFromBytes(privKeyBytes)
	if err != nil {
		cLogger.Panicf("unable to convert private key for peer identity: %s", err)
	}

	waspKeyPair := cryptolib.KeyPairFromPrivateKey(waspPrivKey)
	cLogger.Infof("this node identity: %v", waspKeyPair.GetPublicKey())
	return registry.NewNodeIdentity(waspKeyPair)
}

func NewDKGServer(lc fx.Lifecycle, cfg *config.Config, cLogger *logger.Logger, nodeIdentity *registry.NodeIdentity) *Node {
	ParamsRegistries.TrustedPeers.FilePath = "waspdb/trusted_peers.json"
	ParamsRegistries.DKShares.Path = "waspdb/dkshares"
	trustedPeersRegistryProvider, err := registry.NewTrustedPeersRegistryImpl(ParamsRegistries.TrustedPeers.FilePath)

	if err != nil {
		panic(err)
	}

	cLogger.Infof("pub:%s", nodeIdentity.NodeIdentity().GetPublicKey())
	ParamsPeering.PeeringURL = cfg.AppConfig.PeeringURL
	ParamsPeering.Port = cfg.AppConfig.Port
	netImpl, tnmImpl, err := lpp.NewNetworkProvider(
		ParamsPeering.PeeringURL,
		ParamsPeering.Port,
		nodeIdentity.NodeIdentity(),
		trustedPeersRegistryProvider,
		metrics.NewPeeringMetricsProvider(),
		cLogger,
	)

	if tnmImpl == nil {

	}
	protoParams := &iotago.ProtocolParameters{
		Version:     tpkg.TestProtoParas.Version,
		NetworkName: tpkg.TestProtoParas.NetworkName,
		Bech32HRP:   tpkg.TestProtoParas.Bech32HRP,
		MinPoWScore: tpkg.TestProtoParas.MinPoWScore,
		RentStructure: iotago.RentStructure{
			VByteCost:    10,
			VBFactorData: 1,
			VBFactorKey:  1,
		},
		TokenSupply: tpkg.TestProtoParas.TokenSupply,
	}

	dkSharesRegistry, err := registry.NewDKSharesRegistry(ParamsRegistries.DKShares.Path, nodeIdentity.NodeIdentity().GetPrivateKey(), protoParams.Bech32HRP)

	dkgNode, err := NewNode(
		nodeIdentity.NodeIdentity(), netImpl, dkSharesRegistry,
		cLogger,
	)
	dkgNode.TrustedNetworkManager = tnmImpl

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			fmt.Println("Starting DKG node server at ")

			return nil
		},
		OnStop: func(ctx context.Context) error {
			dkgNode.Close()
			return nil
		},
	})
	return dkgNode
}
