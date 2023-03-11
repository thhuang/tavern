package ctx

import (
	"context"
	"time"

	"github.com/thhuang/go-server/utils/clock"
)

type CTX struct {
	context.Context
	// logrus.FieldLogger
	clock.Clock
}

func Backbround() CTX {
	return CTX{
		Context: context.Background(),
		// FieldLogger: logrus.StandardLogger(),
		Clock: clock.New(time.Now()),
	}
}
