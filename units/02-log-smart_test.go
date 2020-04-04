package tests

import (
	"testing"

	logging "github.com/codemodify/systemkit-logging"
	loggingFormat "github.com/codemodify/systemkit-logging-formatters-timerfc3339nano"
	loggingMixMulti "github.com/codemodify/systemkit-logging-mixers-multi"
	loggingPersistConsole "github.com/codemodify/systemkit-logging-persisters-console"
	loggingPersistFile "github.com/codemodify/systemkit-logging-persisters-file"
)

func Test_02(t *testing.T) {
	logging.SetLogger(
		logging.NewPipe([]logging.Logger{
			loggingFormat.NewTimeRFC3339NanoFormatter(),
			loggingMixMulti.NewMultiLogger([]logging.Logger{
				loggingPersistConsole.NewConsoleLoggerDefaultColors(),
				loggingPersistFile.NewFileLoggerDefaultName(),
			}),
		}),
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
