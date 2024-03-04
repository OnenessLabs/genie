package lib

import (
	"bytes"
	"context"
	"encoding/hex"
	"errors"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/urfave/cli/v2"

	"github.com/OnenessLabs/genie/newrand/client"
	httpmock "github.com/OnenessLabs/genie/newrand/client/test/http/mock"
	commonutils "github.com/OnenessLabs/genie/newrand/common"
	"github.com/OnenessLabs/genie/newrand/crypto"
	"github.com/OnenessLabs/genie/newrand/test/mock"
)

var (
	opts         []client.Option
	latestClient client.Client
)

const (
	fakeGossipRelayAddr = "/ip4/8.8.8.8/tcp/9/p2p/QmSoLju6m7xTh3DuokvT3886QRYqxAzb1kShaanJgW36yx"
	fakeChainHash       = "6093f9e4320c285ac4aab50ba821cd5678ec7c5015d3d9d11ef89e2a99741e83"
)

func mockAction(c *cli.Context) error {
	res, err := Create(c, false, opts...)
	latestClient = res
	return err
}

func run(args []string) error {
	app := cli.NewApp()
	app.Name = "mock-client"
	app.Flags = ClientFlags
	app.Action = mockAction

	return app.Run(args)
}

func TestClientLib(t *testing.T) {
	opts = []client.Option{}
	err := run([]string{"mock-client"})
	if err == nil {
		t.Fatal("need to specify a connection method.", err)
	}

	sch, err := crypto.GetSchemeFromEnv()
	require.NoError(t, err)

	addr, info, cancel, _ := httpmock.NewMockHTTPPublicServer(t, false, sch)
	defer cancel()

	grpcLis, _ := mock.NewMockGRPCPublicServer(t, ":0", false, sch)
	go grpcLis.Start()
	defer grpcLis.Stop(context.Background())

	args := []string{"mock-client", "--url", "http://" + addr, "--grpc-connect", grpcLis.Addr(), "--insecure"}
	err = run(args)
	if err != nil {
		t.Fatal("GRPC should work", err)
	}

	args = []string{"mock-client", "--url", "https://" + addr}
	err = run(args)
	if err == nil {
		t.Fatal("http needs insecure or hash", err)
	}

	args = []string{"mock-client", "--url", "http://" + addr, "--hash", hex.EncodeToString(info.Hash())}
	err = run(args)
	if err != nil {
		t.Fatal("http should construct", err)
	}

	args = []string{"mock-client", "--relay", fakeGossipRelayAddr}
	err = run(args)
	if err == nil {
		t.Fatal("relays need URL to get chain info and hash", err)
	}

	args = []string{"mock-client", "--relay", fakeGossipRelayAddr, "--hash", hex.EncodeToString(info.Hash())}
	err = run(args)
	if err == nil {
		t.Fatal("relays need URL to get chain info and hash", err)
	}

	args = []string{"mock-client", "--url", "http://" + addr, "--relay", fakeGossipRelayAddr, "--hash", hex.EncodeToString(info.Hash())}
	err = run(args)
	if err != nil {
		t.Fatal("unable to get relay to work", err)
	}
}

func TestClientLibGroupConfTOML(t *testing.T) {
	err := run([]string{"mock-client", "--relay", fakeGossipRelayAddr, "--group-conf", groupTOMLPath()})
	if err != nil {
		t.Fatal(err)
	}
}

func TestClientLibGroupConfJSON(t *testing.T) {
	sch, err := crypto.GetSchemeFromEnv()
	require.NoError(t, err)

	addr, info, cancel, _ := httpmock.NewMockHTTPPublicServer(t, false, sch)
	defer cancel()

	var b bytes.Buffer
	info.ToJSON(&b, nil)

	infoPath := filepath.Join(t.TempDir(), "info.json")

	err = os.WriteFile(infoPath, b.Bytes(), 0644)
	if err != nil {
		t.Fatal(err)
	}

	err = run([]string{"mock-client", "--url", "http://" + addr, "--group-conf", infoPath})
	if err != nil {
		t.Fatal(err)
	}
}

func TestClientLibChainHashOverrideError(t *testing.T) {
	err := run([]string{
		"mock-client",
		"--relay",
		fakeGossipRelayAddr,
		"--group-conf",
		groupTOMLPath(),
		"--hash",
		fakeChainHash,
	})
	if !errors.Is(err, commonutils.ErrInvalidChainHash) {
		t.Fatal("expected error from mismatched chain hashes. Got: ", err)
	}
}

func TestClientLibListenPort(t *testing.T) {
	err := run([]string{"mock-client", "--relay", fakeGossipRelayAddr, "--port", "0.0.0.0:0", "--group-conf", groupTOMLPath()})
	if err != nil {
		t.Fatal(err)
	}
}

func groupTOMLPath() string {
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		return ""
	}
	return filepath.Join(filepath.Dir(file), "..", "..", "..", "test", "default.toml")
}
