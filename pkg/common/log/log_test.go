package log

import (
	"testing"
)

func TestLog(t *testing.T) {
	logger := New()
	logger.Info("TEST")
	logger.Warn("TEST")
	logger.Error("TEST")
	logger.Debug("TEST")
}
