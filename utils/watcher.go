package utils

import (
	"encoding/json"
	"time"

	"github.com/nats-io/go-nats"
)

var _ Watcher = (*watcher)(nil)

type Watcher interface {
	// Next is retrieving the next message from the topic
	Next() (interface{}, error)
	// Stop is unsubscribing from the topic
	Stop()
}

type Opts struct {
	Timeout time.Duration
}

type watcher struct {
	sub  *nats.Subscription
	opts *Opts
}

func NewWatcher(opts *Opts, sub *nats.Subscription) Watcher {
	var w = new(watcher)
	w.opts = opts
	w.sub = sub

	return w
}

func (w *watcher) Next() (interface{}, error) {
	var msg interface{}

	for {
		m, err := w.sub.NextMsg(w.opts.Timeout)
		if err != nil && err == nats.ErrTimeout {
			continue
		} else if err != nil {
			return nil, err
		}

		if err = json.Unmarshal(m.Data, &msg); err != nil {
			return nil, err
		}
		break
	}

	return msg, nil
}

func (w *watcher) Stop() {
	w.sub.Unsubscribe()
}
