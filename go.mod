module github.com/coinbase/rosetta-sdk-go

go 1.22.5

require (
	github.com/DataDog/zstd v1.5.2
	github.com/NemProject/nem/gocrypto v0.0.1
	github.com/Zilliqa/gozilliqa-sdk v1.2.1-0.20201201074141-dd0ecada1be6
	github.com/btcsuite/btcd v0.22.1
	github.com/cenkalti/backoff v2.2.1+incompatible
	github.com/coinbase/kryptology v1.8.0
	github.com/coinbase/rosetta-sdk-go/types v1.0.0
	github.com/dgraph-io/badger/v2 v2.2007.4
	github.com/ethereum/go-ethereum v1.10.21
	github.com/fatih/color v1.13.0
	github.com/gorilla/mux v1.8.0
	github.com/lucasjones/reggen v0.0.0-20180717132126-cdb49ff09d77
	github.com/neilotoole/errgroup v0.1.6
	github.com/segmentio/fasthash v1.0.3
	github.com/stretchr/testify v1.7.2
	github.com/tidwall/gjson v1.14.2
	github.com/tidwall/sjson v1.2.5
	github.com/vmihailenco/msgpack/v5 v5.3.5
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/btcsuite/btcutil v1.0.3-0.20201208143702-a53e38424cce // indirect
	github.com/bwesterb/go-ristretto v1.2.0 // indirect
	github.com/cespare/xxhash v1.1.0 // indirect
	github.com/consensys/gnark-crypto v0.5.3 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dgraph-io/ristretto v0.0.3 // indirect
	github.com/dgryski/go-farm v0.0.0-20200201041132-a6ae2369ad13 // indirect
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/klauspost/compress v1.12.3 // indirect
	github.com/mattn/go-colorable v0.1.9 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/objx v0.1.1 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.0 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	golang.org/x/crypto v0.25.0 // indirect
	golang.org/x/net v0.21.0 // indirect
	golang.org/x/sys v0.22.0 // indirect
	google.golang.org/protobuf v1.33.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/coinbase/rosetta-sdk-go/types v1.0.0 => ./types

// temporary until gocrypto is published
replace github.com/NemProject/nem/gocrypto v0.0.1 => ../../NemProject/nem/gocrypto
