package logger

import (
	"sync"

	l "github.com/nats-io/nats-streaming-server/logger"
	log "github.com/sirupsen/logrus"
)

var _ l.Logger = (*natsLogger)(nil)

type natsLogger struct {
	log *log.Entry
	sync.RWMutex
}

func New() *natsLogger {
	return &natsLogger{}
}

func (n *natsLogger) SetLogger(l *log.Entry) {
	n.Lock()
	defer n.Unlock()

	n.log = l
}

func (n *natsLogger) Errorf(format string, v ...interface{}) {
	n.logFunc(func(log *log.Entry, format string, v ...interface{}) {
		log.Errorf(format, v...)
	}, format, v...)
}

func (n *natsLogger) Debugf(format string, v ...interface{}) {
	n.logFunc(func(log *log.Entry, format string, v ...interface{}) {
		log.Debugf(format, v...)
	}, format, v...)
}

func (n *natsLogger) Fatalf(format string, v ...interface{}) {
	n.logFunc(func(log *log.Entry, format string, v ...interface{}) {
		log.Fatalf(format, v...)
	}, format, v...)
}

func (n *natsLogger) Noticef(format string, v ...interface{}) {
	n.logFunc(func(log *log.Entry, format string, v ...interface{}) {
		log.Infof(format, v...)
	}, format, v...)
}

func (n *natsLogger) Tracef(format string, v ...interface{}) {
	n.logFunc(func(log *log.Entry, format string, v ...interface{}) {
		return
	}, format, v...)
}

func (n *natsLogger) logFunc(f func(log *log.Entry, format string, v ...interface{}), format string, args ...interface{}) {
	n.Lock()
	defer n.Unlock()

	if n.log == nil {
		return
	}

	f(n.log, format, args...)
}
