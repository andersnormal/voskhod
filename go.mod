module github.com/andersnormal/voskhod

go 1.13

replace github.com/coreos/bbolt v1.3.0 => go.etcd.io/bbolt v1.3.1-etcd.8

exclude github.com/coreos/etcd v3.3.13+incompatible

require (
	github.com/Microsoft/go-winio v0.4.14 // indirect
	github.com/andersnormal/pkg v0.0.0-20190904210201-9dfdf11cc13f
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/docker/docker v1.13.1
	github.com/docker/go-connections v0.4.0
	github.com/docker/go-units v0.4.0 // indirect
	github.com/fiorix/protoc-gen-cobra v0.0.0-20181029091941-dffa0bfa45cc
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/golang/protobuf v1.3.2
	github.com/hashicorp/go-hclog v0.9.2 // indirect
	github.com/hashicorp/go-immutable-radix v1.1.0 // indirect
	github.com/hashicorp/golang-lru v0.5.3 // indirect
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/konsorten/go-windows-terminal-sequences v1.0.2 // indirect
	github.com/magiconair/properties v1.8.1 // indirect
	github.com/nats-io/gnatsd v1.4.1
	github.com/nats-io/go-nats v1.7.2
	github.com/nats-io/nats-server/v2 v2.1.0 // indirect
	github.com/nats-io/nats-streaming-server v0.16.2
	github.com/nats-io/nats.go v1.8.1
	github.com/nats-io/stan.go v0.5.0
	github.com/opencontainers/go-digest v1.0.0-rc1 // indirect
	github.com/pelletier/go-toml v1.5.0 // indirect
	github.com/pkg/errors v0.8.1
	github.com/prometheus/procfs v0.0.5 // indirect
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/afero v1.2.2 // indirect
	github.com/spf13/cobra v0.0.5
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.4.0
	github.com/stretchr/testify v1.4.0
	golang.org/x/crypto v0.0.0-20191011191535-87dc89f01550 // indirect
	golang.org/x/net v0.0.0-20190603091049-60506f45cf65
	golang.org/x/oauth2 v0.0.0-20180821212333-d2e6202438be
	golang.org/x/sys v0.0.0-20191020212454-3e7259c5e7c2 // indirect
	google.golang.org/grpc v1.21.0
)
