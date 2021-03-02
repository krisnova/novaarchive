package cron

import (
	"os"
	"testing"

	"github.com/kris-nova/logger"
)

var (

	// testService is the service we can call
	// to run our time.Duration crons
	testService *Service = NewService("testservice")
)

// TestMain is used to bootstrap the testService
func TestMain(m *testing.M) {
	exitCode := m.Run()
	if exitCode != 0 {
		logger.Always("Exiting with code [%d]", exitCode)
		os.Exit(exitCode)
	}
	logger.Always("Win!")
}
