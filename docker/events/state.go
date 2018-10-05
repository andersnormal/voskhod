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

package events

import (
	"errors"
	"sync"
	"time"

	"github.com/docker/docker/api/types/events"
)

const (
	defaultEventWatchTimeout = 100 * time.Millisecond
)

type Events interface {
	AddEventListener(listener chan<- events.Message) error
	RemoveEventListener(listener chan<- events.Message) error
}

type eventsState struct {
	sync.RWMutex
	sync.WaitGroup
	MessagesC <-chan events.Message // has to be changed to type
	ErrC      <-chan error
	listeners []chan<- events.Message
	watch     bool
}

// APIActor represents an actor that accomplishes something for an event
type APIActor struct {
	ID         string            `json:"id,omitempty"`
	Attributes map[string]string `json:"attributes,omitempty"`
}

type APIEvent struct {
	// New API Fields in 1.22
	Action string   `json:"action,omitempty"`
	Type   string   `json:"type,omitempty"`
	Actor  APIActor `json:"actor,omitempty"`

	// Old API fields for < 1.22
	Status string `json:"status,omitempty"`
	ID     string `json:"id,omitempty"`
	From   string `json:"from,omitempty"`

	// Fields in both
	Time     int64 `json:"time,omitempty"`
	TimeNano int64 `json:"timeNano,omitempty"`
}

var (
	ErrListenerAlreadyExists = errors.New("Listener already exists")

	ErrListenerDoesNotExists = errors.New("Listener does not exists")

	ErrNoListeners = errors.New("No listeners")

	// EOFEvent is sent when the event listener receives an EOF error.
	EOFEvent = &APIEvent{
		Type:   "EOF",
		Status: "EOF",
	}
)

// AddEventListener is adding an event listener to the event state
func (e *eventsState) AddEventListener(listener chan<- events.Message) error {
	return e.addListener(listener)
}

// RemoveEventListener removes a listener from the monitor.
func (e *eventsState) RemoveEventListener(listener chan<- events.Message) error {
	return e.removeListener(listener)
}

// New connects container events from the Docker daemon with an event state
func New(messageC <-chan events.Message, errC <-chan error) Events {
	return &eventsState{
		MessagesC: messageC,
		ErrC:      errC,
	}
}

func (e *eventsState) addListener(listener chan<- events.Message) error {
	var err error

	e.Lock()
	defer e.Unlock()

	if e.listenerExists(listener) {
		return ErrListenerAlreadyExists
	}

	e.listeners = append(e.listeners, listener)
	e.Add(1) // keep track of listeners

	// watch events
	e.watchEvents()

	return err
}

func (e *eventsState) removeListener(listener chan<- events.Message) error {
	var err error

	e.Lock()
	defer e.Unlock()

	if !e.listenerExists(listener) {
		return ErrListenerDoesNotExists
	}

	var newListeners []chan<- events.Message
	for _, l := range e.listeners {
		if l != listener {
			newListeners = append(newListeners, l)
		}
	}
	e.listeners = newListeners

	return err
}

func (e *eventsState) listenerExists(a chan<- events.Message) bool {
	for _, b := range e.listeners {
		if b == a {
			return true
		}
	}
	return false
}

func (e *eventsState) listernersCount() int {
	e.RLock()
	defer e.RUnlock()
	return len(e.listeners)
}

func (e *eventsState) noListeners() bool {
	e.RLock()
	defer e.RUnlock()

	return len(e.listeners) == 0
}

func (e *eventsState) closeListeners() {
	for _, l := range e.listeners {
		close(l)
		e.Done()
	}
	e.listeners = nil
}

func (e *eventsState) watchEvents() { // pass in a valid client
	if e.watch {
		return
	}

	go func() {
		timeout := time.After(defaultEventWatchTimeout)

		for {
			select {
			case ev, ok := <-e.MessagesC:
				if !ok {
					return
				}

				e.broadcastEvent(ev)
			case <-timeout:
				continue
			}
		}
	}()

	e.watch = true
}

func (e *eventsState) broadcastEvent(event events.Message) {
	e.RLock()
	defer e.RUnlock()

	// block to wait for event to be send
	e.Add(1)
	defer e.Done()

	if len(e.listeners) == 0 {
		// e.err <- ErrNoListeners
		return
	}

	for _, listener := range e.listeners {
		// should then be buffer
		select {
		case listener <- event:
		default:
		}
	}
}
