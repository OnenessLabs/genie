package dkg

import (
	"context"
	"fmt"
	"github.com/OnenessLabs/genie/dkg/packages/cryptolib"
	"github.com/OnenessLabs/genie/dkg/packages/metrics"
	"github.com/OnenessLabs/genie/dkg/packages/peering/lpp"
	"github.com/OnenessLabs/genie/dkg/packages/registry"
	"github.com/iotaledger/hive.go/logger"
	iotago "github.com/iotaledger/iota.go/v3"
	"github.com/iotaledger/iota.go/v3/tpkg"
	"go.uber.org/fx"
)

type ParametersPeering struct {
	PeeringURL string `default:"0.0.0.0:4000" usage:"node host address as it is recognized by other peers"`
	Port       int    `default:"4000" usage:"port for Wasp committee connection/peering"`
}

var ParamsPeering = &ParametersPeering{}
var ParamsRegistries = &ParametersRegistries{}

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

func NewDKGServer(lc fx.Lifecycle) *Node {
	nodeIdentity := cryptolib.NewKeyPair()

	trustedPeersRegistryProvider, err := registry.NewTrustedPeersRegistryImpl(ParamsRegistries.TrustedPeers.FilePath)

	loggerParams := &logger.Config{}
	loggerRoot, err := logger.NewRootLogger(*loggerParams)
	if err != nil {
		panic(err)
	}

	var cLogger *logger.Logger
	cLogger = loggerRoot.Named("App")
	ParamsPeering.PeeringURL = "0.0.0.0:4000"
	ParamsPeering.Port = 4000
	netImpl, tnmImpl, err := lpp.NewNetworkProvider(
		ParamsPeering.PeeringURL,
		ParamsPeering.Port,
		nodeIdentity,
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

	dkSharesRegistry, err := registry.NewDKSharesRegistry(ParamsRegistries.DKShares.Path, nodeIdentity.GetPrivateKey(), protoParams.Bech32HRP)

	dkgNode, err := NewNode(
		nodeIdentity, netImpl, dkSharesRegistry,
		cLogger,
	)
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
