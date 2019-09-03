module github.com/noah-blockchain/explorer-genesis-uploader

go 1.12

replace mellium.im/sasl v0.2.1 => github.com/mellium/sasl v0.2.1

require (
	github.com/go-pg/pg v8.0.5+incompatible
	github.com/noah-blockchain/noah-explorer-extender v0.1.0
	github.com/noah-blockchain/noah-explorer-tools v0.1.0
	github.com/sirupsen/logrus v1.4.2
	github.com/valyala/fasthttp v1.4.0
)
