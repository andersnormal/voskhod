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

package dockerclient

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/docker/docker/api/types"
	dc "github.com/docker/docker/client"
	"github.com/docker/go-connections/sockets"
	"github.com/katallaxie/voskhod/docker/dockeriface"
)

var _ dockeriface.Client = (*client)(nil)
var _ Client = (*client)(nil)

// ErrRedirect is the error returned by checkRedirect when the request is non-GET.
var ErrRedirect = errors.New("unexpected redirect in response")

// Client is the interface
type Client interface {
	dockeriface.Client
}

type client struct {
	_timeOnce sync.Once
	lock      sync.Mutex
	dc        *dc.Client
}

// New creates a new docker client
func New() (Client, error) {
	var err error
	var client = new(client)

	httpClient, err := defaultHTTPClient(DefaultDockerHost)
	if err != nil {
		return nil, err
	}

	client.dc, err = dc.NewClient(DefaultDockerHost, "1.25", httpClient, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return client, err
}

func defaultHTTPClient(host string) (*http.Client, error) {
	proto, addr, _, err := dc.ParseHost(host)
	if err != nil {
		return nil, err
	}

	transport := new(http.Transport)
	sockets.ConfigureTransport(transport, proto, addr)
	return &http.Client{
		Transport:     transport,
		CheckRedirect: CheckRedirect,
	}, nil
}

// CheckRedirect specifies the policy for dealing with redirect responses:
// If the request is non-GET return `ErrRedirect`. Otherwise use the last response.
//
// Go 1.8 changes behavior for HTTP redirects (specifically 301, 307, and 308) in the client .
// The Docker client (and by extension docker API client) can be made to to send a request
// like POST /containers//start where what would normally be in the name section of the URL is empty.
// This triggers an HTTP 301 from the daemon.
// In go 1.8 this 301 will be converted to a GET request, and ends up getting a 404 from the daemon.
// This behavior change manifests in the client in that before the 301 was not followed and
// the client did not generate an error, but now results in a message like Error response from daemon: page not found.
func CheckRedirect(req *http.Request, via []*http.Request) error {
	if via[0].Method == http.MethodGet {
		return http.ErrUseLastResponse
	}
	return ErrRedirect
}

func (c *client) Version(ctx context.Context, time time.Duration) (types.Version, error) {
	ctx, cancel := context.WithTimeout(ctx, time)
	defer cancel()

	return c.dc.ServerVersion(ctx)
}
