package pkg

import (
	"os"
	"os/signal"
	c "context"
	"syscall"

	log "github.com/sirupsen/logrus"
)

// func /////////////////////////////////////////

func InterruptContext(ctx c.Context) (c.Context, c.CancelFunc) {
	return signal.NotifyContext(
		ctx,
		os.Interrupt,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		syscall.SIGHUP,
	)
}

