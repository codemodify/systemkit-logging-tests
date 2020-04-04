package tests

import (
	"testing"
	"time"

	logging "github.com/codemodify/systemkit-logging"
	loggingFormat "github.com/codemodify/systemkit-logging-formatters-timerfc3339nano"
	loggingMixAsync "github.com/codemodify/systemkit-logging-mixers-async"
	loggingMixBuffered "github.com/codemodify/systemkit-logging-mixers-buffered"
	loggingMixMulti "github.com/codemodify/systemkit-logging-mixers-multi"
	loggingPersistConsole "github.com/codemodify/systemkit-logging-persisters-console"
	loggingPersistFile "github.com/codemodify/systemkit-logging-persisters-file"
)

func Test_04(t *testing.T) {
	logging.SetLogger(
		logging.NewPipe([]logging.Logger{
			loggingFormat.NewTimeRFC3339NanoFormatter(),
			loggingMixAsync.NewAsyncLogger(
				loggingMixBuffered.NewBufferedLogger(
					loggingMixMulti.NewMultiLogger([]logging.Logger{
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

	logging.Instance().KeepOnlyLogs(logging.TypeDebug)

	logging.Instance().Trace("Trace line")
	logging.Instance().Panic("Panic line")
	logging.Instance().Fatal("Fatal line")
	logging.Instance().Error("Error line")
	logging.Instance().Warning("Warning line")
	logging.Instance().Info("Info line")
	logging.Instance().Success("Success line")
	logging.Instance().Debug("Debug line")

	time.Sleep(5 * time.Second)
}
