package pkg

import (
	"os"
	"syscall"
	"os/signal"
	c "context"

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

