package signal

import (
	"context"
	"os"
	"os/signal"

	"github.com/namo-io/go-kit/log"
)

// WaitSignal this is blocking function until signal comes or context cancel
// ex) signals: syscall.SIGKILL, syscall.SIGTERM, syscall.SIGINT, os.Interrupt
func WaitSignal(ctx context.Context, signals ...os.Signal) {
	logger := log.New().WithContext(ctx)
	logger.Info("start wait signal")

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, signals...)

	sig := <-sigs
	logger.Infof("signal: %s", sig.String())

	signal.Stop(sigs)
}
