package tests

import (
	"testing"

	logging "github.com/codemodify/systemkit-logging"
	loggingFormat "github.com/codemodify/systemkit-logging-formatters-timerfc3339nano"
	loggingMixMulti "github.com/codemodify/systemkit-logging-mixers-multi"
	loggingPersistConsole "github.com/codemodify/systemkit-logging-persisters-consolewithcolors"
	loggingPersistFile "github.com/codemodify/systemkit-logging-persisters-file"
)

func Test_02(t *testing.T) {
	logging.SetLogger(
		logging.NewPipe([]logging.CoreLogger{
			loggingFormat.NewTimeRFC3339NanoFormatter(),
			loggingMixMulti.NewMultiLogger([]logging.CoreLogger{
				loggingPersistConsole.NewConsoleLoggerDefaultColors(),
				loggingPersistFile.NewFileLoggerDefaultName(),
			}),
		}),
	).KeepOnlyLogs(logging.TypeDebug) // keep logs up to

	logging.Trace("Trace line")
	logging.Panic("Panic line")
	logging.Fatal("Fatal line")
	logging.Error("Error line")
	logging.Warning("Warning line")
	logging.Info("Info line")
	logging.Success("Success line")
	logging.Debug("Debug line")
}
