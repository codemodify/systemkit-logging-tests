package tests

import (
	"testing"
	"time"

	logging "github.com/codemodify/systemkit-logging"
	loggingFormat "github.com/codemodify/systemkit-logging-formatters-timerfc3339nano"
	loggingMixAsync "github.com/codemodify/systemkit-logging-mixers-async"
	loggingMixBuffered "github.com/codemodify/systemkit-logging-mixers-buffered"
	loggingMixMulti "github.com/codemodify/systemkit-logging-mixers-multi"
	loggingPersistConsole "github.com/codemodify/systemkit-logging-persisters-consolewithcolors"
	loggingPersistFile "github.com/codemodify/systemkit-logging-persisters-file"
)

func Test_05(t *testing.T) {
	logging.SetLogger(
		logging.NewPipe([]logging.CoreLogger{
			loggingFormat.NewTimeRFC3339NanoFormatter(),
			loggingMixAsync.NewAsyncLogger(
				loggingMixBuffered.NewBufferedLogger(
					loggingMixMulti.NewMultiLogger([]logging.CoreLogger{
						loggingPersistConsole.NewConsoleLoggerDefaultColors(),
						loggingPersistFile.NewFileLoggerDefaultName(),
					}),
					loggingMixBuffered.BufferConfig{
						MaxLogEntries: 1,
					},
				),
			),
		}),
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

	time.Sleep(5 * time.Second)
}
