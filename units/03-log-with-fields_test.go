package tests

import (
	"fmt"
	"testing"

	logging "github.com/codemodify/systemkit-logging"
	loggingWithFields "github.com/codemodify/systemkit-logging-extenders-withfields"
	loggingFormat "github.com/codemodify/systemkit-logging-formatters-timerfc3339nano"
	loggingMixMulti "github.com/codemodify/systemkit-logging-mixers-multi"
	loggingPersistConsole "github.com/codemodify/systemkit-logging-persisters-console"
	loggingPersistFile "github.com/codemodify/systemkit-logging-persisters-file"
)

func Test_03(t *testing.T) {
	myLogger := loggingWithFields.NewWithFieldsLogger(
		logging.SetLogger(
			logging.NewPipe([]logging.Logger{
				loggingFormat.NewTimeRFC3339NanoFormatter(),
				loggingMixMulti.NewMultiLogger([]logging.Logger{
					loggingPersistConsole.NewConsoleLoggerDefaultColors(),
					loggingPersistFile.NewFileLoggerDefaultName(),
				}),
			}),
		),
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

	fmt.Println()

	myLogger.TraceWithFields(loggingWithFields.Fields{
		"TAG":     "tag-1",
		"message": "Trace line",
	})
	myLogger.PanicWithFields(loggingWithFields.Fields{
		"TAG":     "tag-1",
		"message": "Panic line",
	})
	myLogger.FatalWithFields(loggingWithFields.Fields{
		"TAG":     "tag-1",
		"message": "Fatal line",
	})
	myLogger.ErrorWithFields(loggingWithFields.Fields{
		"TAG":     "tag-1",
		"message": "Error line",
	})
	myLogger.WarningWithFields(loggingWithFields.Fields{
		"TAG":     "tag-1",
		"message": "Warning line",
	})
	myLogger.InfoWithFields(loggingWithFields.Fields{
		"TAG":     "tag-1",
		"message": "Info line",
	})
	myLogger.SuccessWithFields(loggingWithFields.Fields{
		"TAG":     "tag-1",
		"message": "Success line",
	})
	myLogger.DebugWithFields(loggingWithFields.Fields{
		"TAG":     "tag-1",
		"message": "Debug line",
	})
}
