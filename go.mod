module github.com/andersnormal/voskhod

require (
	github.com/docker/docker v1.13.1
	github.com/docker/go-connections v0.4.0
	github.com/fiorix/protoc-gen-cobra v0.0.0-20181029091941-dffa0bfa45cc
	github.com/golang/groupcache v0.0.0-20190129154638-5b532d6fd5ef // indirect
	github.com/golang/protobuf v1.2.0
	github.com/andersnormal/voskhod v0.0.0-20190116221807-a3eef93d9019
	github.com/kelseyhightower/envconfig v1.3.0
	github.com/nats-io/gnatsd v1.3.0
	github.com/nats-io/go-nats v1.7.0
	github.com/nats-io/go-nats-streaming v0.4.0
	github.com/nats-io/nats-streaming-server v0.11.2
	github.com/satori/go.uuid v1.2.0
	github.com/sirupsen/logrus v1.3.0
	github.com/spf13/cobra v0.0.3
	github.com/spf13/pflag v1.0.3
	github.com/spf13/viper v1.2.1
	go.etcd.io/bbolt v1.3.2 // indirect
	go.etcd.io/etcd v3.3.11+incompatible
	golang.org/x/net v0.0.0-20190110200230-915654e7eabc
	golang.org/x/oauth2 v0.0.0-20190111185915-36a7019397c4
	golang.org/x/sync v0.0.0-20181221193216-37e7f081c4d4
	google.golang.org/grpc v1.17.0
)

replace github.com/coreos/bbolt v1.3.0 => go.etcd.io/bbolt v1.3.1-etcd.8

exclude github.com/coreos/etcd v3.3.10+incompatible
