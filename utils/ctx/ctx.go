package ctx

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/thhuang/go-server/utils/clock"
)

type logKey string

type CTX struct {
	context.Context
	logrus.FieldLogger
	clock.Clock
}

func Backbround() CTX {
	return CTX{
		Context:     context.Background(),
		FieldLogger: logrus.StandardLogger(),
		Clock:       clock.New(time.Now()),
	}
}

func (context CTX) Now() time.Time {
	return context.Clock.Now().UTC()
}

func WithValue(parent CTX, key string, val interface{}) CTX {
	return CTX{
		Context:     context.WithValue(parent, logKey(key), val),
		FieldLogger: parent.FieldLogger.WithField(key, val),
		Clock:       parent.Clock,
	}
}

func WithValues(parent CTX, kvs map[string]interface{}) CTX {
	c := parent
	for k, v := range kvs {
		c = WithValue(c, k, v)
	}
	return c
}

func WithCancel(parent CTX) (CTX, context.CancelFunc) {
	context, cancel := context.WithCancel(parent)
	return CTX{
		Context:     context,
		FieldLogger: parent.FieldLogger,
	}, cancel
}
