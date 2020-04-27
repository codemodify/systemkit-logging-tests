package tests

import (
	"testing"

	logging "github.com/codemodify/systemkit-logging"
	loggingPersistConsole "github.com/codemodify/systemkit-logging-persisters-console"
)

func Test_01(t *testing.T) {
	logging.SetLogger(loggingPersistConsole.NewConsoleLogger())

	logging.Trace("Trace line")
	logging.Panic("Panic line")
	logging.Fatal("Fatal line")
	logging.Error("Error line")
	logging.Warning("Warning line")
	logging.Info("Info line")
	logging.Success("Success line")
	logging.Debug("Debug line")
}
