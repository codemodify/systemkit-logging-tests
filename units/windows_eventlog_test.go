// +build windows

package tests

import (
	"testing"

	logging "github.com/codemodify/systemkit-logging"
	loggingPersistWinEvent "github.com/codemodify/systemkit-logging-persisters-windowseventlog"
)

func Test_windows_eventlog(t *testing.T) {
	logging.SetLogger(
		loggingPersistWinEvent.NewWindowsEventLogger(),
	)

	logging.KeepOnlyLogs(logging.TypeDebug)

	logging.Trace("Trace line")
	logging.Panic("Panic line")
	logging.Fatal("Fatal line")
	logging.Error("Error line")
	logging.Warning("Warning line")
	logging.Info("Info line")
	logging.Success("Success line")
	logging.Debug("Debug line")
}
