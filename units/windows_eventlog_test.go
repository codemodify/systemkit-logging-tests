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

	logging.Instance().KeepOnlyLogs(logging.TypeDebug)

	logging.Instance().Trace("Trace line")
	logging.Instance().Panic("Panic line")
	logging.Instance().Fatal("Fatal line")
	logging.Instance().Error("Error line")
	logging.Instance().Warning("Warning line")
	logging.Instance().Info("Info line")
	logging.Instance().Success("Success line")
	logging.Instance().Debug("Debug line")
}
