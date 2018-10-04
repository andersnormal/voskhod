// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package agent

import (
	"context"
	"sync"

	"github.com/katallaxie/voskhod/config"
	"github.com/katallaxie/voskhod/docker/dockerapi"

	log "github.com/sirupsen/logrus"
)

// Signal is the channel to control the Voskhod Agent
type Signal int

// Agent describes the interface to a Voskhod Agent
type Agent interface {
	// Start does all things necessary to start an agent
	Start() func() error
	// Stop is doing all things necessary to nicely stop an agent
	Stop() error
}

type agent struct {
	cfg *config.Config
	ctx context.Context

	dc dockerclient.Client

	logger *log.Entry

	// lock is used to safely access the client
	lock sync.RWMutex
}
