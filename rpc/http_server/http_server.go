package http_server

import (
	"context"
	"fmt"
	"github.com/OnenessLabs/genie/config"
	"github.com/OnenessLabs/genie/dkg"
	"github.com/OnenessLabs/genie/dkg/packages/cryptolib"
	"github.com/OnenessLabs/genie/dkg/packages/peering"
	"github.com/OnenessLabs/genie/dkg/packages/tcrypto"
	iotago "github.com/iotaledger/iota.go/v3"
	"github.com/iotaledger/iota.go/v3/tpkg"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/samber/lo"
	"go.uber.org/fx"
	"net/http"
	"sync"
	"time"
)

type App struct {
	undoneCtx context.Context
	ctx       context.Context
	ctxCancel func()

	wg       sync.WaitGroup
	stopping uint32

	echo    *echo.Echo
	dkgNode *dkg.Node
}

func NewEchoServer(lc fx.Lifecycle, cfg *config.Config, dkgNode *dkg.Node) *echo.Echo {
	app := &App{}
	app.echo = echo.New()

	app.dkgNode = dkgNode

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			fmt.Println("Starting DKG node server at ")
			// InitServer - run server
			// CORS default
			// Allows requests from any origin wth GET, HEAD, PUT, POST or DELETE method.
			app.echo.Use(middleware.Logger())
			app.echo.Use(middleware.CORS())
			app.echo.POST("/v1/node/dks", app.generateDKS)
			app.echo.GET("/v1/node/dks/{sharedAddress}", app.getDKSInfo)
			app.echo.GET("/metrics", echo.WrapHandler(promhttp.Handler()))

			go app.echo.Start(cfg.AppConfig.HttpAddr)

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
	return app.echo

}

// DKSharesPostRequest is a POST request for creating new DKShare.
type DKSharesPostRequest struct {
	PeerPubKeysOrNames []string `json:"peerIdentities" swagger:"desc(Names or hex encoded public keys of trusted peers to run DKG on.),required"`
	Threshold          uint16   `json:"threshold" swagger:"desc(Should be =< len(PeerPublicIdentities)),required,min(1)"`
	TimeoutMS          uint32   `json:"timeoutMS" swagger:"desc(Timeout in milliseconds.),required,min(1)"`
}

// DKSharesInfo stands for the DKShare representation, returned by the GET and POST methods.
type DKSharesInfo struct {
	Address         string   `json:"address" swagger:"desc(New generated shared address.),required"`
	PeerIdentities  []string `json:"peerIdentities" swagger:"desc(Identities of the nodes sharing the key. (Hex)),required"`
	PeerIndex       *uint16  `json:"peerIndex" swagger:"desc(Index of the node returning the share, if it is a member of the sharing group.),required,min(1)"`
	PublicKey       string   `json:"publicKey" swagger:"desc(Used public key. (Hex)),required"`
	PublicKeyShares []string `json:"publicKeyShares" swagger:"desc(Public key shares for all the peers. (Hex)),required"`
	Threshold       uint16   `json:"threshold" swagger:"required,min(1)"`
}

func (app *App) generateDKS(e echo.Context) error {
	generateDKSRequest := DKSharesPostRequest{}

	if err := e.Bind(&generateDKSRequest); err != nil {
		return e.String(http.StatusBadRequest, fmt.Sprintf("Invalid property: %v", "body"))
	}

	sharesInfo, err := app.GenerateDistributedKey(generateDKSRequest.PeerPubKeysOrNames, generateDKSRequest.Threshold, time.Duration(generateDKSRequest.TimeoutMS)*time.Millisecond)
	if err != nil {
		panic(err)
	}

	return e.JSON(http.StatusOK, sharesInfo)
}

func (app *App) getDKSInfo(e echo.Context) error {
	_, sharedAddress, err := iotago.ParseBech32(e.Param("sharedAddress"))
	if err != nil {
		return e.String(http.StatusBadRequest, fmt.Sprintf("Invalid property: %v", sharedAddress))
	}

	sharesInfo, err := app.GetShares(sharedAddress)
	if err != nil {
		panic(err)
	}

	return e.JSON(http.StatusOK, sharesInfo)
}

const (
	roundRetry = 1 * time.Second // Retry for Peer <-> Peer communication.
	stepRetry  = 3 * time.Second // Retry for Initiator -> Peer communication.
)

func (app *App) GenerateDistributedKey(peerPubKeysOrNames []string, threshold uint16, timeout time.Duration) (*DKSharesInfo, error) {
	trustedPeers, err := app.dkgNode.TrustedNetworkManager.TrustedPeersByPubKeyOrName(peerPubKeysOrNames)
	if err != nil {
		return nil, err
	}
	peerPubKeys := lo.Map(trustedPeers, func(tp *peering.TrustedPeer, _ int) *cryptolib.PublicKey {
		return tp.PubKey()
	})

	dkShare, err := app.dkgNode.GenerateDistributedKey(peerPubKeys, threshold, roundRetry, stepRetry, timeout)
	if err != nil {
		return nil, err
	}

	dkShareInfo, err := app.createDKModel(dkShare)
	if err != nil {
		return nil, err
	}

	return dkShareInfo, nil
}

func (app *App) createDKModel(dkShare tcrypto.DKShare) (*DKSharesInfo, error) {
	publicKey, err := dkShare.DSSSharedPublic().MarshalBinary()
	if err != nil {
		return nil, err
	}

	dssPublicShares := dkShare.DSSPublicShares()
	pubKeySharesHex := make([]string, len(dssPublicShares))
	for i := range dssPublicShares {
		publicKeyShare, err := dssPublicShares[i].MarshalBinary()
		if err != nil {
			return nil, err
		}

		pubKeySharesHex[i] = iotago.EncodeHex(publicKeyShare)
	}

	peerIdentities := dkShare.GetNodePubKeys()
	peerIdentitiesHex := make([]string, len(peerIdentities))
	for i := range peerIdentities {
		peerIdentitiesHex[i] = peerIdentities[i].String()
	}

	dkShareInfo := &DKSharesInfo{
		Address:         dkShare.GetAddress().Bech32(tpkg.TestProtoParas.Bech32HRP),
		PeerIdentities:  peerIdentitiesHex,
		PeerIndex:       dkShare.GetIndex(),
		PublicKey:       iotago.EncodeHex(publicKey),
		PublicKeyShares: pubKeySharesHex,
		Threshold:       dkShare.GetT(),
	}

	return dkShareInfo, nil
}

func (app *App) GetShares(sharedAddress iotago.Address) (*DKSharesInfo, error) {
	dkShare, err := app.dkgNode.DkShareRegistryProvider.LoadDKShare(sharedAddress)
	if err != nil {
		return nil, err
	}

	dkShareInfo, err := app.createDKModel(dkShare)
	if err != nil {
		return nil, err
	}

	return dkShareInfo, nil
}
