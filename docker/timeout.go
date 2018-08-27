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

package docker

import "time"

const (
	// DefaultListContainersTimeout is the timeout for
	// listing container from the Docker API.
	DefaultListContainersTimeout = 10 * time.Minute
	// DefaultLoadImageTimeout is the timeout for loading an image
	// by the Docker API.
	DefaultLoadImageTimeout = 10 * time.Minute
	// DefaultCreateContainerTimeout is the timeout for creating an
	// container by the Docker API.
	DefaultCreateContainerTimeout = 4 * time.Minute
	// DefaultStopContainerTimeout is the tiemout for stopping
	// an container by the Docker API.
	DefaultStopContainerTimeout = 30 * time.Second
	// DefaultRemoveContainerTimeout is the timeout for removing an
	// container by the Docker API.
	DefaultRemoveContainerTimeout = 5 * time.Minute
	// DefaultInspectContainerTimeout is the timeout for inspecing an
	// container by the Docker API.
	DefaultInspectContainerTimeout = 30 * time.Second
	// DefaultRemoveImageTimeout is the timeout for removing an image
	// by the Docker API.
	DefaultRemoveImageTimeout = 3 * time.Minute
	// DefaultVersionTimeout is the timeout for getting the version from
	// the Docker API.
	DefaultVersionTimeout = 10 * time.Second
)
