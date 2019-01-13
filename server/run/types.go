package server

import (
	"context"
	"sync"

	"github.com/katallaxie/voskhod/server/config"
	"github.com/katallaxie/voskhod/server/nats"

	"github.com/andersnormal/pkg"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

// Signal is the channel to control the Voskhod Agent
type Signal int

// Server describes the interface to a Voskhod Agent
type Server interface {
	// ServeAPI is starting the API listener
	ServeAPI()
	// ServeNATS is starting the NATs
	ServeNats(s nats.Server)
	// Ready is waiting for the server to be ready
	Ready() error
	// Wait is waiting for everything to end :-)
	Wait() error
}

type server struct {
	// config to be passed along
	cfg *config.Config

	// error Group
	errG *errgroup.Group

	// error Context
	errCtx context.Context

	// track the grpc server
	api *grpc.Server

	// NATs
	nats nats.Server

	// logger instance
	logger *log.Entry

	// mutex
	sync.RWMutex

	// ready
	ready pkg.ReadyEvents
}

type API struct {
	// config to be used with the api
	cfg *config.Config

	// logger
	logger *log.Entry
}

type APIReady struct {
	*pkg.ReadyEvent
}

type NatsReady struct {
	*pkg.ReadyEvent
}
