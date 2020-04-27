package tests

import (
	"testing"

	logging "github.com/codemodify/systemkit-logging"
	loggingFormat "github.com/codemodify/systemkit-logging-formatters-timerfc3339nano"
	loggingMixMulti "github.com/codemodify/systemkit-logging-mixers-multi"
	loggingPersistConsole "github.com/codemodify/systemkit-logging-persisters-consolewithcolors"
)

func Test_04(t *testing.T) {
	logging.SetLogger(
		loggingMixMulti.NewMultiLogger([]logging.CoreLogger{
			// PIPE 1 - will format
			logging.NewPipe([]logging.CoreLogger{
				loggingFormat.NewTimeRFC3339NanoFormatter(),
				loggingPersistConsole.NewConsoleLoggerDefaultColors(),
			}),

			// PIPE 2 - will NOT format
			logging.NewPipe([]logging.CoreLogger{
				loggingPersistConsole.NewConsoleLoggerDefaultColors(),
			}),
		}),
	)

	logging.Trace("Trace line")
	logging.Panic("Panic line")
	logging.Fatal("Fatal line")
	logging.Error("Error line")
	logging.Warning("Warning line")
	logging.Info("Info line")
	logging.Success("Success line")
	logging.Debug("Debug line")
}
