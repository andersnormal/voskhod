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

package cmd

import (
	"os"
	"os/signal"
	"syscall"
)

func (r *root) configureSignals() {
	r.sys = make(chan os.Signal, 1)

	signal.Notify(r.sys, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR1)
}

func (r *root) exitSignal() {
	r.exit <- 1
}

// watchSignals is watching configured signals
func (r *root) watchSignals() {
	// defer
	defer r.exitSignal()

	// config singals
	r.configureSignals()

	// loop blocking
	for {
		sig := <-r.sys
		switch sig {
		case syscall.SIGUSR1:
		default:
			r.cancel()
			return
		}
	}
}
