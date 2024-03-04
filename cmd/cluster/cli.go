package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/OnenessLabs/genie/internal/debug"
	"github.com/OnenessLabs/genie/internal/flags"
	"github.com/ethereum/go-ethereum/console/prompt"
	"github.com/urfave/cli/v2"
	"go.uber.org/automaxprocs/maxprocs"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
)

var app = flags.NewApp("the Oneness network node cli")

// DKSharesPostRequest struct for DKSharesPostRequest
type DKSharesPostRequest struct {
	// Names or hex encoded public keys of trusted peers to run DKG on.
	PeerIdentities []string `json:"peerIdentities"`
	// Should be =< len(PeerPublicIdentities)
	Threshold uint32 `json:"threshold"`
	// Timeout in milliseconds.
	TimeoutMS uint32 `json:"timeoutMS"`
}

// NewDKSharesPostRequest instantiates a new DKSharesPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDKSharesPostRequest(peerIdentities []string, threshold uint32, timeoutMS uint32) *DKSharesPostRequest {
	this := DKSharesPostRequest{}
	this.PeerIdentities = peerIdentities
	this.Threshold = threshold
	this.TimeoutMS = timeoutMS
	return &this
}

const ENV_VARS_PREFIX = "GENIE"

func init() {
	// Initialize the CLI app and start
	app.Action = func(ctx *cli.Context) error {
		if args := ctx.Args().Slice(); len(args) > 0 {
			return fmt.Errorf("invalid command: %q", args[0])
		}

		to := uint32(60 * 1000)
		threshold := 3
		peerPubKeys := []string{"0xc8b1d3a85cee9949bde989144da8391269d2632866794a6eab11d0b7777c07cc",
			"0x3b5059ef01bb33b471e7ba24e8b7a4880eb545c274d65ac3e5019dd0e835e067",
			"0xe40f69d5a97e26b38296ec55485767498b2d2565eb814720d09e1eb31a765f7e"}

		dKSharesPostRequest := DKSharesPostRequest{
			Threshold:      uint32(threshold),
			TimeoutMS:      to,
			PeerIdentities: peerPubKeys,
		}

		jsonData, err := json.Marshal(dKSharesPostRequest)
		if err != nil {
			// Handle the error
		}

		req, err := http.NewRequest("POST", "http://localhost:8081/v1/node/dks", bytes.NewBuffer(jsonData))
		if err != nil {
			// Handle the error
		}

		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			// Handle the error
		}
		defer resp.Body.Close()

		fmt.Println("response Status:", resp.Status)
		fmt.Println("response Headers:", resp.Header)
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("response Body:", string(body))

		return nil
	}
	sort.Sort(cli.CommandsByName(app.Commands))

	app.Flags = flags.Merge(
		debug.Flags,
	)
	flags.AutoEnvVars(app.Flags, ENV_VARS_PREFIX)

	app.Before = func(ctx *cli.Context) error {
		maxprocs.Set() // Automatically set GOMAXPROCS to match Linux container CPU quota.
		flags.MigrateGlobalFlags(ctx)
		if err := debug.Setup(ctx); err != nil {
			return err
		}
		flags.CheckEnvVars(ctx, app.Flags, ENV_VARS_PREFIX)
		return nil
	}
	app.After = func(ctx *cli.Context) error {
		debug.Exit()
		prompt.Stdin.Close() // Resets terminal mode.
		return nil
	}
}

func main() {
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
