module github.com/OnenessLabs/genie

go 1.21.6

replace go.dedis.ch/kyber/v3 => github.com/kape1395/kyber/v3 v3.0.14-0.20230124095845-ec682ff08c93 // branch: dkg-2suites

require (
	github.com/ethereum/go-ethereum v1.13.14
	github.com/fjl/memsize v0.0.2
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.19.1
	github.com/hashicorp/go-bexpr v0.1.10
	github.com/iotaledger/hive.go/constraints v0.0.0-20231106113411-94ac829adbb2
	github.com/iotaledger/hive.go/crypto v0.0.0-20231106113411-94ac829adbb2
	github.com/iotaledger/hive.go/ds v0.0.0-20231106113411-94ac829adbb2
	github.com/iotaledger/hive.go/ierrors v0.0.0-20231106113411-94ac829adbb2
	github.com/iotaledger/hive.go/kvstore v0.0.0-20231106113411-94ac829adbb2
	github.com/iotaledger/hive.go/logger v0.0.0-20231106113411-94ac829adbb2
	github.com/iotaledger/hive.go/runtime v0.0.0-20231106113411-94ac829adbb2
	github.com/iotaledger/hive.go/serializer/v2 v2.0.0-rc.1.0.20231106113411-94ac829adbb2
	github.com/iotaledger/iota.go/v3 v3.0.0-rc.3
	github.com/libp2p/go-libp2p v0.30.0
	github.com/mattn/go-colorable v0.1.13
	github.com/mattn/go-isatty v0.0.19
	github.com/minio/blake2b-simd v0.0.0-20160723061019-3f5f724cb5b1
	github.com/mr-tron/base58 v1.2.0
	github.com/multiformats/go-multiaddr v0.12.2
	github.com/samber/lo v1.38.1
	github.com/stretchr/testify v1.8.4
	github.com/urfave/cli/v2 v2.27.1
	github.com/wollac/iota-crypto-demo v0.0.0-20221117162917-b10619eccb98
	go.dedis.ch/kyber/v3 v3.1.0
	go.uber.org/automaxprocs v1.5.2
	go.uber.org/fx v1.20.1
	golang.org/x/crypto v0.18.0
	golang.org/x/exp v0.0.0-20231110203233-9a3e6036ecaa
	google.golang.org/genproto/googleapis/api v0.0.0-20240125205218-1f4bbc51befe
	google.golang.org/grpc v1.62.0
	google.golang.org/protobuf v1.32.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)

require (
	filippo.io/edwards25519 v1.0.0 // indirect
	github.com/StackExchange/wmi v1.2.1 // indirect
	github.com/benbjohnson/clock v1.3.5 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/btcsuite/btcd/btcec/v2 v2.2.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/containerd/cgroups v1.1.0 // indirect
	github.com/coreos/go-systemd/v22 v22.5.0 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.3 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/davidlazar/go-crypto v0.0.0-20200604182044-b73af7476f6c // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.2.0 // indirect
	github.com/docker/go-units v0.5.0 // indirect
	github.com/elastic/gosigar v0.14.2 // indirect
	github.com/flynn/noise v1.0.0 // indirect
	github.com/francoispqt/gojay v1.2.13 // indirect
	github.com/go-ole/go-ole v1.3.0 // indirect
	github.com/go-task/slim-sprig v0.0.0-20230315185526-52ccab3ef572 // indirect
	github.com/godbus/dbus/v5 v5.1.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/mock v1.6.0 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/gopacket v1.1.19 // indirect
	github.com/google/pprof v0.0.0-20230821062121-407c9e7a662f // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/holiman/uint256 v1.2.4 // indirect
	github.com/huin/goupnp v1.3.0 // indirect
	github.com/iancoleman/orderedmap v0.3.0 // indirect
	github.com/iotaledger/hive.go/lo v0.0.0-20231106113411-94ac829adbb2 // indirect
	github.com/iotaledger/hive.go/stringify v0.0.0-20231019081410-32f61b05bebe // indirect
	github.com/iotaledger/iota.go v1.0.0 // indirect
	github.com/ipfs/go-cid v0.4.1 // indirect
	github.com/ipfs/go-log/v2 v2.5.1 // indirect
	github.com/jackpal/go-nat-pmp v1.0.2 // indirect
	github.com/jbenet/go-temp-err-catcher v0.1.0 // indirect
	github.com/klauspost/compress v1.16.7 // indirect
	github.com/klauspost/cpuid/v2 v2.2.6 // indirect
	github.com/koron/go-ssdp v0.0.4 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/libp2p/go-buffer-pool v0.1.0 // indirect
	github.com/libp2p/go-cidranger v1.1.0 // indirect
	github.com/libp2p/go-flow-metrics v0.1.0 // indirect
	github.com/libp2p/go-libp2p-asn-util v0.3.0 // indirect
	github.com/libp2p/go-msgio v0.3.0 // indirect
	github.com/libp2p/go-nat v0.2.0 // indirect
	github.com/libp2p/go-netroute v0.2.1 // indirect
	github.com/libp2p/go-reuseport v0.4.0 // indirect
	github.com/libp2p/go-yamux/v4 v4.0.1 // indirect
	github.com/marten-seemann/tcp v0.0.0-20210406111302-dfbc87cc63fd // indirect
	github.com/mattn/go-runewidth v0.0.13 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.4 // indirect
	github.com/miekg/dns v1.1.55 // indirect
	github.com/mikioh/tcpinfo v0.0.0-20190314235526-30a79bb1804b // indirect
	github.com/mikioh/tcpopt v0.0.0-20190314235656-172688c1accc // indirect
	github.com/minio/sha256-simd v1.0.1 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/mitchellh/pointerstructure v1.2.0 // indirect
	github.com/multiformats/go-base32 v0.1.0 // indirect
	github.com/multiformats/go-base36 v0.2.0 // indirect
	github.com/multiformats/go-multiaddr-dns v0.3.1 // indirect
	github.com/multiformats/go-multiaddr-fmt v0.1.0 // indirect
	github.com/multiformats/go-multibase v0.2.0 // indirect
	github.com/multiformats/go-multicodec v0.9.0 // indirect
	github.com/multiformats/go-multihash v0.2.3 // indirect
	github.com/multiformats/go-multistream v0.4.1 // indirect
	github.com/multiformats/go-varint v0.0.7 // indirect
	github.com/onsi/ginkgo/v2 v2.11.0 // indirect
	github.com/opencontainers/runtime-spec v1.1.0 // indirect
	github.com/pbnjay/memory v0.0.0-20210728143218-7b4eea64cf58 // indirect
	github.com/pelletier/go-toml/v2 v2.1.0 // indirect
	github.com/peterh/liner v1.1.1-0.20190123174540-a2c9a5303de7 // indirect
	github.com/petermattis/goid v0.0.0-20230904192822-1876fd5063bc // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/prometheus/client_golang v1.14.0 // indirect
	github.com/prometheus/client_model v0.4.0 // indirect
	github.com/prometheus/common v0.37.0 // indirect
	github.com/prometheus/procfs v0.8.0 // indirect
	github.com/quic-go/qpack v0.4.0 // indirect
	github.com/quic-go/qtls-go1-20 v0.3.2 // indirect
	github.com/quic-go/quic-go v0.37.6 // indirect
	github.com/quic-go/webtransport-go v0.5.3 // indirect
	github.com/raulk/go-watchdog v1.3.0 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/sasha-s/go-deadlock v0.3.1 // indirect
	github.com/shirou/gopsutil v3.21.4-0.20210419000835-c7a38de76ee5+incompatible // indirect
	github.com/spaolacci/murmur3 v1.1.0 // indirect
	github.com/tklauser/go-sysconf v0.3.12 // indirect
	github.com/tklauser/numcpus v0.6.1 // indirect
	github.com/xrash/smetrics v0.0.0-20201216005158-039620a65673 // indirect
	go.dedis.ch/fixbuf v1.0.3 // indirect
	go.dedis.ch/protobuf v1.0.11 // indirect
	go.uber.org/dig v1.17.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.26.0 // indirect
	golang.org/x/mod v0.14.0 // indirect
	golang.org/x/net v0.20.0 // indirect
	golang.org/x/sync v0.6.0 // indirect
	golang.org/x/sys v0.16.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	golang.org/x/tools v0.15.0 // indirect
	google.golang.org/genproto v0.0.0-20240123012728-ef4313101c80 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240125205218-1f4bbc51befe // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	lukechampine.com/blake3 v1.2.1 // indirect
)
